package astextract

import "testing"

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				input: `fmt.Println("init")`,
			},
			want:    "&ast.CallExpr {\n  Fun: &ast.SelectorExpr {\n    X: &ast.Ident {\n      Name: \"fmt\",\n    },\n    Sel: &ast.Ident {\n      Name: \"Println\",\n    },\n  },\n  Args: []ast.Expr {\n    &ast.BasicLit {\n      Kind: token.STRING,\n      Value: \"\\\"init\\\"\",\n    },\n  },\n}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
