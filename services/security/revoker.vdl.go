// This file was auto-generated by the veyron vdl tool.
// Source: revoker.vdl

package security

import (
	"veyron.io/veyron/veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron.io/veyron/veyron2"
	_gen_context "veyron.io/veyron/veyron2/context"
	_gen_ipc "veyron.io/veyron/veyron2/ipc"
	_gen_naming "veyron.io/veyron/veyron2/naming"
	_gen_vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	_gen_wiretype "veyron.io/veyron/veyron2/wiretype"
)

// RevocationToken can be presented to a revocation service to revoke a caveat
type RevocationToken [16]byte

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Revoker is the interface for preventing discharges from being issued. The
// dicharger ensures that no discharges will be issued for caveats that
// have been explicitly revoked using this interface. To prevent discharge
// stealing caveats just have to be unique; the exact structure is not relevant
// to the client or the verifier. To make Revoker's job easy, each caveat
// contains a SHA256 hash of its revocation token. To revoke a caveat C and
// have it added to the discharger's blacklist, one simply needs to call
// Revoke(x) with an x s.t.  SHA256(x) = C. All caveats for which this has not
// been revoked will get discharges, irrespective of who created them. This
// means that the existence of a valid discharge does not imply that a
// corresponding caveat exists, and even if it does, it may not be meant for
// use with this revocation service. Just looking at discharges is meaningless,
// a valid (Caveat, Discharge) pair is what can be relied on for
// authentication. Not keeping track of non-revoked caveats enables
// performance improvements on the Discharger side.
// Revoker is the interface the client binds and uses.
// Revoker_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Revoker_ExcludingUniversal interface {
	// Revoke ensures that iff a nil is returned, all discharge requests to the
	// caveat with nonce sha256(caveatPreimage) are going to be denied.
	Revoke(ctx _gen_context.T, caveatPreimage RevocationToken, opts ..._gen_ipc.CallOpt) (err error)
}
type Revoker interface {
	_gen_ipc.UniversalServiceMethods
	Revoker_ExcludingUniversal
}

// RevokerService is the interface the server implements.
type RevokerService interface {

	// Revoke ensures that iff a nil is returned, all discharge requests to the
	// caveat with nonce sha256(caveatPreimage) are going to be denied.
	Revoke(context _gen_ipc.ServerContext, caveatPreimage RevocationToken) (err error)
}

// BindRevoker returns the client stub implementing the Revoker
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindRevoker(name string, opts ..._gen_ipc.BindOpt) (Revoker, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubRevoker{defaultClient: client, name: name}

	return stub, nil
}

// NewServerRevoker creates a new server stub.
//
// It takes a regular server implementing the RevokerService
// interface, and returns a new server stub.
func NewServerRevoker(server RevokerService) interface{} {
	return &ServerStubRevoker{
		service: server,
	}
}

// clientStubRevoker implements Revoker.
type clientStubRevoker struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubRevoker) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubRevoker) Revoke(ctx _gen_context.T, caveatPreimage RevocationToken, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Revoke", []interface{}{caveatPreimage}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRevoker) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRevoker) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRevoker) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubRevoker wraps a server that implements
// RevokerService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubRevoker struct {
	service RevokerService
}

func (__gen_s *ServerStubRevoker) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Revoke":
		return []interface{}{security.Label(4)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubRevoker) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Revoke"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "caveatPreimage", Type: 66},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.ArrayType{Elem: 0x41, Len: 0x10, Name: "veyron.io/veyron/veyron/services/security.RevocationToken", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubRevoker) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubRevoker) Revoke(call _gen_ipc.ServerCall, caveatPreimage RevocationToken) (err error) {
	err = __gen_s.service.Revoke(call, caveatPreimage)
	return
}
