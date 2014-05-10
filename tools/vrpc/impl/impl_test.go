package impl_test

import (
	"bytes"
	"sort"
	"strings"
	"testing"

	"veyron/lib/cmdline"
	"veyron/tools/vrpc/impl"
	"veyron/tools/vrpc/test_base"

	"veyron2"
	"veyron2/ipc"
	"veyron2/naming"
	"veyron2/rt"
	"veyron2/vlog"
)

type server struct{}

// TypeTester interface implementation

func (*server) Bool(call ipc.Context, i1 bool) (bool, error) {
	vlog.VI(2).Info("Bool(%v) was called.", i1)
	return i1, nil
}

func (*server) Float32(call ipc.Context, i1 float32) (float32, error) {
	vlog.VI(2).Info("Float32(%u) was called.", i1)
	return i1, nil
}

func (*server) Float64(call ipc.Context, i1 float64) (float64, error) {
	vlog.VI(2).Info("Float64(%u) was called.", i1)
	return i1, nil
}

func (*server) Int32(call ipc.Context, i1 int32) (int32, error) {
	vlog.VI(2).Info("Int32(%v) was called.", i1)
	return i1, nil
}

func (*server) Int64(call ipc.Context, i1 int64) (int64, error) {
	vlog.VI(2).Info("Int64(%v) was called.", i1)
	return i1, nil
}

func (*server) String(call ipc.Context, i1 string) (string, error) {
	vlog.VI(2).Info("String(%v) was called.", i1)
	return i1, nil
}

func (*server) Byte(call ipc.Context, i1 byte) (byte, error) {
	vlog.VI(2).Info("Byte(%v) was called.", i1)
	return i1, nil
}

func (*server) UInt32(call ipc.Context, i1 uint32) (uint32, error) {
	vlog.VI(2).Info("UInt32(%u) was called.", i1)
	return i1, nil
}

func (*server) UInt64(call ipc.Context, i1 uint64) (uint64, error) {
	vlog.VI(2).Info("UInt64(%u) was called.", i1)
	return i1, nil
}

func (*server) InputArray(call ipc.Context, i1 [2]uint8) error {
	vlog.VI(2).Info("CInputArray(%v) was called.", i1)
	return nil
}

func (*server) OutputArray(call ipc.Context) ([2]uint8, error) {
	vlog.VI(2).Info("COutputArray() was called.")
	return [2]uint8{1, 2}, nil
}

func (*server) InputMap(call ipc.Context, i1 map[uint8]uint8) error {
	vlog.VI(2).Info("CInputMap(%v) was called.", i1)
	return nil
}

func (*server) OutputMap(call ipc.Context) (map[uint8]uint8, error) {
	vlog.VI(2).Info("COutputMap() was called.")
	return map[uint8]uint8{1: 2}, nil
}

func (*server) InputSlice(call ipc.Context, i1 []uint8) error {
	vlog.VI(2).Info("CInputSlice(%v) was called.", i1)
	return nil
}

func (*server) OutputSlice(call ipc.Context) ([]uint8, error) {
	vlog.VI(2).Info("COutputSlice() was called.")
	return []uint8{1, 2}, nil
}

func (*server) InputStruct(call ipc.Context, i1 test_base.Struct) error {
	vlog.VI(2).Info("CInputStruct(%v) was called.", i1)
	return nil
}

func (*server) OutputStruct(call ipc.Context) (test_base.Struct, error) {
	vlog.VI(2).Info("COutputStruct() was called.")
	return test_base.Struct{X: 1, Y: 2}, nil
}

func (*server) NoArguments(call ipc.Context) error {
	vlog.VI(2).Info("NoArguments() was called.")
	return nil
}

func (*server) MultipleArguments(call ipc.Context, i1, i2 int32) (int32, int32, error) {
	vlog.VI(2).Info("MultipleArguments(%v,%v) was called.", i1, i2)
	return i1, i2, nil
}

func (*server) StreamingOutput(call ipc.Context, nStream int32, item bool, reply test_base.TypeTesterServiceStreamingOutputStream) error {
	vlog.VI(2).Info("StreamingOutput(%v,%v) was called.", nStream, item)
	for i := int32(0); i < nStream; i++ {
		reply.Send(item)
	}
	return nil
}

func startServer(t *testing.T, r veyron2.Runtime) (ipc.Server, naming.Endpoint, error) {
	dispatcher := ipc.SoloDispatcher(test_base.NewServerTypeTester(&server{}), nil)
	server, err := r.NewServer()
	if err != nil {
		t.Errorf("NewServer failed: %v", err)
		return nil, nil, err
	}
	if err := server.Register("", dispatcher); err != nil {
		t.Errorf("Register failed: %v", err)
		return nil, nil, err
	}
	endpoint, err := server.Listen("tcp", "localhost:0")
	if err != nil {
		t.Errorf("Listen failed: %v", err)
		return nil, nil, err
	}
	if err := server.Publish(""); err != nil {
		t.Errorf("Publish failed: %v", err)
		return nil, nil, err
	}
	return server, endpoint, nil
}

func stopServer(t *testing.T, server ipc.Server) {
	if err := server.Stop(); err != nil {
		t.Errorf("server.Stop failed: %v", err)
	}
}

func testInvocation(t *testing.T, buffer *bytes.Buffer, cmd *cmdline.Command, args []string, expected string) {
	buffer.Reset()
	if err := cmd.Execute(args); err != nil {
		t.Errorf("%v", err)
		return
	}
	if output := strings.Trim(buffer.String(), "\n"); output != expected {
		t.Errorf("Incorrect invoke output: expected %s, got %s", expected, output)
		return
	}
}

func testError(t *testing.T, cmd *cmdline.Command, args []string, expected string) {
	if err := cmd.Execute(args); err == nil || !strings.Contains(err.Error(), expected) {
		t.Errorf("Expected error: ...%v..., got: %v", expected, err)
	}
}

func TestVRPC(t *testing.T) {
	runtime := rt.Init()
	// Skip defer runtime.Shutdown() to avoid messing up other tests in the
	// same process.
	server, endpoint, err := startServer(t, runtime)
	if err != nil {
		return
	}
	defer stopServer(t, server)

	// Setup the command-line.
	cmd := impl.Root()
	var stdout, stderr bytes.Buffer
	cmd.Init(nil, &stdout, &stderr)

	name := naming.JoinAddressName(endpoint.String(), "//")
	// Test the 'describe' command.
	if err := cmd.Execute([]string{"describe", name}); err != nil {
		t.Errorf("%v", err)
		return
	}

	expectedSignature := []string{
		"func Bool(I1 bool) (O1 bool, E error)",
		"func Float32(I1 float32) (O1 float32, E error)",
		"func Float64(I1 float64) (O1 float64, E error)",
		"func Int32(I1 int32) (O1 int32, E error)",
		"func Int64(I1 int64) (O1 int64, E error)",
		"func String(I1 string) (O1 string, E error)",
		"func Byte(I1 byte) (O1 byte, E error)",
		"func UInt32(I1 uint32) (O1 uint32, E error)",
		"func UInt64(I1 uint64) (O1 uint64, E error)",
		"func InputArray(I1 [2]byte) (E error)",
		"func InputMap(I1 map[byte]byte) (E error)",
		"func InputSlice(I1 []byte) (E error)",
		"func InputStruct(I1 struct{X int32, Y int32}) (E error)",
		"func OutputArray() (O1 [2]byte, E error)",
		"func OutputMap() (O1 map[byte]byte, E error)",
		"func OutputSlice() (O1 []byte, E error)",
		"func OutputStruct() (O1 struct{X int32, Y int32}, E error)",
		"func NoArguments() (error)",
		"func MultipleArguments(I1 int32, I2 int32) (O1 int32, O2 int32, E error)",
		"func StreamingOutput(NumStreamItems int32, StreamItem bool) stream<_, bool> (error)",
	}

	signature := make([]string, 0, len(expectedSignature))
	line, err := stdout.ReadBytes('\n')
	for err == nil {
		signature = append(signature, strings.Trim(string(line), "\n"))
		line, err = stdout.ReadBytes('\n')
	}

	sort.Strings(signature)
	sort.Strings(expectedSignature)

	if len(signature) != len(expectedSignature) {
		t.Fatalf("signature lengths don't match %v and %v.", len(signature), len(expectedSignature))
	}

	for i, expectedSig := range expectedSignature {
		if expectedSig != signature[i] {
			t.Errorf("signature line doesn't match: %v and %v\n", expectedSig, signature[i])
		}
	}

	// Test the 'invoke' command.

	tests := [][]string{
		[]string{"Bool", "Bool(true) = [true, <nil>]", "[\"bool\",true]"},
		[]string{"Float32", "Float32(3.2) = [3.2, <nil>]", "[\"float32\",3.2]"},
		[]string{"Float64", "Float64(6.4) = [6.4, <nil>]", "[\"float64\",6.4]"},
		[]string{"Int32", "Int32(-32) = [-32, <nil>]", "[\"int32\",-32]"},
		[]string{"Int64", "Int64(-64) = [-64, <nil>]", "[\"int64\",-64]"},
		[]string{"String", "String(Hello World!) = [Hello World!, <nil>]", "[\"string\",\"Hello World!\"]"},
		[]string{"Byte", "Byte(8) = [8, <nil>]", "[\"byte\",8]"},
		[]string{"UInt32", "UInt32(32) = [32, <nil>]", "[\"uint32\",32]"},
		[]string{"UInt64", "UInt64(64) = [64, <nil>]", "[\"uint64\",64]"},
		// TODO(jsimsa): The InputArray currently triggers an error in the
		// vom decoder. Benj is looking into this.
		//
		// []string{"InputArray", "InputArray([1 2]) = []", "[\"[2]uint\",[1,2]]"},
		[]string{"InputMap", "InputMap(map[1:2]) = [<nil>]", "[\"map[uint]uint\",{\"1\":\"2\"}]"},
		// TODO(jsimsa): The InputSlice currently triggers an error in the
		// vom decoder. Benj is looking into this.
		//
		// []string{"InputSlice", "InputSlice([1 2]) = []", "[\"[]uint\",[1,2]]"},
		[]string{"InputStruct", "InputStruct({1 2}) = [<nil>]",
			"[\"type\",\"veyron2/vrpc/test_base.Struct struct{X int32;Y int32}\"] [\"Struct\",{\"X\":1,\"Y\":2}]"},
		// TODO(jsimsa): The OutputArray currently triggers an error in the
		// vom decoder. Benj is looking into this.
		//
		// []string{"OutputArray", "OutputArray() = [1 2]"}
		[]string{"OutputMap", "OutputMap() = [map[1:2], <nil>]"},
		[]string{"OutputSlice", "OutputSlice() = [[1 2], <nil>]"},
		[]string{"OutputStruct", "OutputStruct() = [{1 2}, <nil>]"},
		[]string{"NoArguments", "NoArguments() = [<nil>]"},
		[]string{"MultipleArguments", "MultipleArguments(1, 2) = [1, 2, <nil>]", "[\"uint32\",1]", "[\"uint32\",2]"},
		[]string{"StreamingOutput", "StreamingOutput(3, true) = <<\n0: true\n1: true\n2: true\n>> [<nil>]", "[\"int8\",3]", "[\"bool\",true ]"},
		[]string{"StreamingOutput", "StreamingOutput(0, true) = [<nil>]", "[\"int8\",0]", "[\"bool\",true ]"},
	}

	for _, test := range tests {
		testInvocation(t, &stdout, cmd, append([]string{"invoke", name, test[0]}, test[2:]...), test[1])
	}

	testErrors := [][]string{
		[]string{"Bool", "usage error"},
		[]string{"DoesNotExit", "invoke: method DoesNotExit not found"},
	}
	for _, test := range testErrors {
		testError(t, cmd, append([]string{"invoke", name, test[0]}, test[2:]...), test[1])
	}
}
