package ast_compare

type Merge interface {
	merge() (any, error)
}
type FileAst struct {
	Comment   CommentCompare
	Package   PackageCompare
	Import    ImportCompare
	Var       VarCompare
	Func      FuncCompare
	Struct    map[string]StructCompare
	Interface map[string]InterfaceCompare
}

type FuncCompare struct {
	new      []any
	old      []any
	conflict map[string][2]any
}

type StructCompare struct {
	new      []any
	old      []any
	conflict map[string][2]any
}

type Import struct {
	alias string
	pkg   string
}
type ImportCompare struct {
	new      []Import
	old      []Import
	conflict map[string][2]Import
}
type VarCompare struct {
	new      []any
	old      []any
	conflict map[string][2]any
}

type PackageCompare struct {
	old string
	new string
}

func (c *PackageCompare) merge() (string, error) {
	return c.new, nil
}

type InterfaceCompare struct {
	new      []any
	old      []any
	conflict map[string][2]any
}

func (c *InterfaceCompare) merge() ([]any, error) {
	rest := c.old
	rest = append(rest, c.new...)
	for s, anies := range c.conflict {
		rest = append(rest, anies[0])
		rest = append(rest, interfaceFuncRename(s, anies[1]))
	}
	return c.new, nil
}

func interfaceFuncRename(name string, n any) any {

	return n
}

type CommentCompare struct {
}
