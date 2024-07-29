package astextract

import (
	"bytes"
	"go/parser"
	"go/token"
)

// Parse converts input to ast
func Parse(input string) (string, error) {
	var b bytes.Buffer

	fset := token.NewFileSet()
	filters := AppendFilters(PosFilter, KeywordFilter, ZeroFilter)

	expr, err := parser.ParseExpr(input)
	if err == nil {
		err = Fprint(&b, fset, expr, filters)
		return b.String(), err
	}

	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		return "", err
	}

	for i, comment := range f.Comments {
		if comment.Pos() > f.Package {
			f.Comments = f.Comments[:i]
		}
	}

	err = Fprint(&b, fset, f, filters)
	if err != nil {
		return "", err
	}

	return b.String(), err
}
