package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"

	"github.com/masseelch/elk/internal/parser"
	"github.com/stoewer/go-strcase"
)

// FormatNode formats an AST node and returns the formatted code as a string.
func FormatNode(node *ast.File) (string, error) {
	fset := token.NewFileSet()
	var buf bytes.Buffer
	// 使用 go/format 包来格式化代码
	if err := format.Node(&buf, fset, node); err != nil {
		panic(err)
	}
	return buf.String(), nil
}

func Gen(cn MsgDefine, pg Package) (string, error) {
	//fset := token.NewFileSet()
	f := ast.File{
		Name: ast.NewIdent(pg.Name),
	}

	err := handleGen(cn, pg, &f)
	if err != nil {
		return "", err
	}
	err = outerGen(cn, pg, &f)
	if err != nil {
		return "", err
	}

	return FormatNode(&f)
}

func handleGen(cn MsgDefine, pg Package, f *ast.File) error {

	var handleList []*ast.Field
	handleFs := func(msg *parser.Msg) {
		handleDef := ast.Field{
			Names: []*ast.Ident{ast.NewIdent(msgIDToHandleName(msg.ID))},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						{

							Type: &ast.SelectorExpr{
								X:   ast.NewIdent("pb"),
								Sel: ast.NewIdent("Context"),
							},
						},
						{
							Type: &ast.StarExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent("proto"),
									Sel: ast.NewIdent(msg.Name),
								},
							},
						},
					},
				},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{Type: ast.NewIdent("error")},
					},
				},
			},
		}
		handleList = append(handleList, &handleDef)
	}
	// 公共消息
	for _, msg := range pg.Coms {
		if msg.Handle {
			handleFs(msg)
		}
	}

	// 特有消息
	for _, msg := range pg.Cms {
		if msg.Handle {
			handleFs(msg)
		}
	}

	spxv := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: &ast.Ident{
					Name: "Handle",
				},
				Type: &ast.InterfaceType{
					Methods: &ast.FieldList{
						List: handleList,
					},
				},
			},
		},
	}

	f.Decls = append(f.Decls, spxv)

	return nil
}

func msgIDToHandleName(id string) string {
	hn := id
	if strings.HasSuffix(id, "_REQ") {
		hn = strings.TrimSuffix(id, "_REQ")
	} else if strings.HasSuffix(id, "_Req") {
		hn = strings.TrimSuffix(id, "_Req")
	} else if strings.HasSuffix(id, "_req") {
		hn = strings.TrimSuffix(id, "_req")
	}
	return strcase.UpperCamelCase(hn)
}

func outerGen(cn MsgDefine, pg Package, f *ast.File) error {
	var msgtp []ast.Stmt
	sgp := func(msg *parser.Msg) {
		swcInfo := &ast.CaseClause{
			List: []ast.Expr{
				&ast.StarExpr{
					X: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "proto",
						},
						Sel: ast.NewIdent(msg.Name),
					},
				},
			},
			Body: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.SelectorExpr{
							X: &ast.Ident{
								Name: "gr",
							},
							Sel: &ast.Ident{
								Name: "MsgID",
							},
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{
								Name: "int",
							},
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X: &ast.Ident{
										Name: "proto",
									},
									Sel: ast.NewIdent(msg.ID),
								},
							},
						},
					},
				},
			},
		}
		msgtp = append(msgtp, swcInfo)
	}
	// 公共消息
	for _, msg := range pg.Coms {
		sgp(msg)
	}

	// 特有消息
	for _, msg := range pg.Cms {
		sgp(msg)
	}
	msgtp = append(msgtp, &ast.CaseClause{
		Body: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.Ident{
						Name: "nil",
					},
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "fmt",
							},
							Sel: &ast.Ident{
								Name: "Errorf",
							},
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"cat not support data to out\"",
							},
						},
					},
				},
			},
		},
	})

	spxv := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: "Out",
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "msg",
							},
						},
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "pg",
							},
							Sel: &ast.Ident{
								Name: "Message",
							},
						},
					},
					&ast.Field{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "opts",
							},
						},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									&ast.Field{
										Type: &ast.ArrayType{
											Elt: &ast.Ident{
												Name: "byte",
											},
										},
									},
								},
							},
							Results: &ast.FieldList{
								List: []*ast.Field{
									&ast.Field{
										Type: &ast.ArrayType{
											Elt: &ast.Ident{
												Name: "byte",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Type: &ast.StarExpr{
							X: &ast.Ident{
								Name: "GameReply",
							},
						},
					},
					&ast.Field{
						Type: &ast.Ident{
							Name: "error",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.DeclStmt{
					Decl: &ast.GenDecl{
						Tok: token.VAR,
						Specs: []ast.Spec{
							&ast.ValueSpec{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "gr",
									},
								},
								Type: &ast.Ident{
									Name: "GameReply",
								},
							},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: "bs",
						},
						&ast.Ident{
							Name: "err",
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: "pg",
								},
								Sel: &ast.Ident{
									Name: "Marshal",
								},
							},
							Args: []ast.Expr{
								&ast.Ident{
									Name: "msg",
								},
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "err",
						},
						Op: token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.Ident{
										Name: "nil",
									},
									&ast.Ident{
										Name: "err",
									},
								},
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X: &ast.Ident{
							Name: "opts",
						},
						Op: token.NEQ,
						Y: &ast.Ident{
							Name: "nil",
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.AssignStmt{
								Lhs: []ast.Expr{
									&ast.Ident{
										Name: "bs",
									},
								},
								Tok: token.ASSIGN,
								Rhs: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.Ident{
											Name: "opts",
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "bs",
											},
										},
									},
								},
							},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.SelectorExpr{
							X: &ast.Ident{
								Name: "gr",
							},
							Sel: &ast.Ident{
								Name: "Data",
							},
						},
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.Ident{
							Name: "bs",
						},
					},
				},
				&ast.TypeSwitchStmt{
					Assign: &ast.ExprStmt{
						X: &ast.TypeAssertExpr{
							X: &ast.Ident{
								Name: "msg",
							},
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.CaseClause{
								List: []ast.Expr{
									&ast.StarExpr{
										X: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "proto",
											},
											Sel: &ast.Ident{
												Name: "Echox",
											},
										},
									},
								},
								Body: []ast.Stmt{
									&ast.AssignStmt{
										Lhs: []ast.Expr{
											&ast.SelectorExpr{
												X: &ast.Ident{
													Name: "gr",
												},
												Sel: &ast.Ident{
													Name: "MsgID",
												},
											},
										},
										Tok: token.ASSIGN,
										Rhs: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.Ident{
													Name: "int",
												},
												Args: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.Ident{
															Name: "proto",
														},
														Sel: &ast.Ident{
															Name: "MSG_CMD_HELLO_REQ",
														},
													},
												},
											},
										},
									},
								},
							},
							&ast.CaseClause{
								List: []ast.Expr{
									&ast.StarExpr{
										X: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "pk",
											},
											Sel: &ast.Ident{
												Name: "Hne",
											},
										},
									},
								},
								Body: []ast.Stmt{
									&ast.AssignStmt{
										Lhs: []ast.Expr{
											&ast.SelectorExpr{
												X: &ast.Ident{
													Name: "gr",
												},
												Sel: &ast.Ident{
													Name: "MsgID",
												},
											},
										},
										Tok: token.ASSIGN,
										Rhs: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.Ident{
													Name: "int",
												},
												Args: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.Ident{
															Name: "proto",
														},
														Sel: &ast.Ident{
															Name: "MSG_CMD_HELLO_REQ",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.UnaryExpr{
							Op: token.AND,
							X: &ast.Ident{
								Name: "gr",
							},
						},
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}

	f.Decls = append(f.Decls, spxv)
	return nil
}
