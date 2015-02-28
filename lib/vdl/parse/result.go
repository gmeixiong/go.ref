package parse

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"v.io/x/ref/lib/vdl/vdlutil"
)

// Pos captures positional information during parsing.
type Pos struct {
	Line int // Line number, starting at 1
	Col  int // Column number (character count), starting at 1
}

// StringPos holds a string and a Pos.
type StringPos struct {
	String string
	Pos    Pos
}

// Returns true iff this Pos has been initialized.  The zero Pos is invalid.
func (p Pos) IsValid() bool {
	return p.Line > 0 && p.Col > 0
}

func (p Pos) String() string {
	if !p.IsValid() {
		return "[no pos]"
	}
	return fmt.Sprintf("%v:%v", p.Line, p.Col)
}

// InferPackageName returns the package name from a group of files.  Every file
// must specify the same package name, otherwise an error is reported in errs.
func InferPackageName(files []*File, errs *vdlutil.Errors) (pkgName string) {
	var firstFile string
	for _, f := range files {
		switch {
		case pkgName == "":
			firstFile = f.BaseName
			pkgName = f.PackageDef.Name
		case pkgName != f.PackageDef.Name:
			errs.Errorf("Files in the same directory must be in the same package; %v has package %v, but %v has package %v", firstFile, pkgName, f.BaseName, f.PackageDef.Name)
		}
	}
	return
}

// Representation of the components of an vdl file.  These data types represent
// the parse tree generated by the parse.

// File represents a parsed vdl file.
type File struct {
	BaseName   string       // Base name of the vdl file, e.g. "foo.vdl"
	PackageDef NamePos      // Name, position and docs of the "package" clause
	Imports    []*Import    // Imports listed in this file.
	ErrorDefs  []*ErrorDef  // Errors defined in this file
	TypeDefs   []*TypeDef   // Types defined in this file
	ConstDefs  []*ConstDef  // Consts defined in this file
	Interfaces []*Interface // Interfaces defined in this file
}

// Config represents a parsed config file.  Config files use a similar syntax as
// vdl files, with similar concepts.
type Config struct {
	FileName  string      // Config file name, e.g. "a/b/foo.config"
	ConfigDef NamePos     // Name, position and docs of the "config" clause
	Imports   []*Import   // Imports listed in this file.
	Config    ConstExpr   // Const expression exported from this config.
	ConstDefs []*ConstDef // Consts defined in this file.
}

// AddImports adds the path imports that don't already exist to c.
func (c *Config) AddImports(path ...string) {
	for _, p := range path {
		if !c.HasImport(p) {
			c.Imports = append(c.Imports, &Import{Path: p})
		}
	}
}

// HasImport returns true iff path exists in c.Imports.
func (c *Config) HasImport(path string) bool {
	for _, imp := range c.Imports {
		if imp.Path == path {
			return true
		}
	}
	return false
}

// Import represents an import definition, which is used to import other
// packages into an vdl file.  An example of the syntax in the vdl file:
//   import foo "some/package/path"
type Import struct {
	NamePos        // e.g. foo (from above), or typically empty
	Path    string // e.g. "some/package/path" (from above)
}

// LocalName returns the name used locally within the File to refer to the
// imported package.
func (i *Import) LocalName() string {
	if i.Name != "" {
		return i.Name
	}
	return path.Base(i.Path)
}

// ErrorDef represents an error definition.
type ErrorDef struct {
	NamePos             // error name, pos and doc
	Params  []*Field    // list of positional parameters
	Actions []StringPos // list of action code identifiers
	Formats []LangFmt   // list of language / format pairs
}

// LangFmt represents a language / format string pair.
type LangFmt struct {
	Lang StringPos // IETF language tag
	Fmt  StringPos // i18n format string in the given language
}

// Pos returns the position of the LangFmt.
func (x LangFmt) Pos() Pos {
	if x.Lang.Pos.IsValid() {
		return x.Lang.Pos
	}
	return x.Fmt.Pos
}

// Interface represents a set of embedded interfaces and methods.
type Interface struct {
	NamePos            // interface name, pos and doc
	Embeds  []*NamePos // names of embedded interfaces
	Methods []*Method  // list of methods
}

// Method represents a method in an interface.
type Method struct {
	NamePos               // method name, pos and doc
	InArgs    []*Field    // list of positional in-args
	OutArgs   []*Field    // list of positional out-args
	InStream  Type        // in-stream type, may be nil
	OutStream Type        // out-stream type, may be nil
	Tags      []ConstExpr // list of method tags
}

// Field represents fields in structs as well as method arguments.
type Field struct {
	NamePos      // field name, pos and doc
	Type    Type // field type, never nil
}

// NamePos represents a name, its associated position and documentation.
type NamePos struct {
	Name      string
	Pos       Pos    // position of first character in name
	Doc       string // docs that occur before the item
	DocSuffix string // docs that occur on the same line after the item
}

func (x *File) String() string      { return fmt.Sprintf("%+v", *x) }
func (x *Import) String() string    { return fmt.Sprintf("%+v", *x) }
func (x *ErrorDef) String() string  { return fmt.Sprintf("%+v", *x) }
func (x *Interface) String() string { return fmt.Sprintf("%+v", *x) }
func (x *Method) String() string    { return fmt.Sprintf("%+v", *x) }
func (x *Field) String() string     { return fmt.Sprintf("%+v", *x) }
func (x *NamePos) String() string   { return fmt.Sprintf("%+v", *x) }

// QuoteStripDoc takes a Doc string, which includes comment markers /**/ and
// double-slash, and returns a raw-quoted string.
//
// TODO(toddw): This should remove comment markers.  This is non-trivial, since
// we should handle removing leading whitespace "rectangles", and might want to
// retain inline /**/ or adjacent /**/ on the same line.  For now we just leave
// them in the output.
func QuoteStripDoc(doc string) string {
	trimmed := strings.Trim(doc, "\n")
	if strconv.CanBackquote(doc) {
		return "`" + trimmed + "`"
	}
	return strconv.Quote(trimmed)
}
