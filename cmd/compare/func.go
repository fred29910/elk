package main

import (
	"go/ast"
	"go/token"
)

func simpleData() *ast.File {
	return &ast.File{
		Name: &ast.Ident{
			Name: "ent",
		},
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"strings\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"tempvs/ent/compartment\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"tempvs/ent/item\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"entgo.io/ent\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"entgo.io/ent/dialect/sql\"",
						},
					},
				},
			},
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{
							Name: "Item",
						},
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Type: &ast.Ident{
											Name: "config",
										},
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`json:\"-\"`",
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "ID",
											},
										},
										Type: &ast.Ident{
											Name: "int",
										},
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`json:\"id,omitempty\"`",
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "Name",
											},
										},
										Type: &ast.Ident{
											Name: "string",
										},
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`json:\"name,omitempty\"`",
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "Edges",
											},
										},
										Type: &ast.Ident{
											Name: "ItemEdges",
										},
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`json:\"edges\"`",
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "compartment_contents",
											},
										},
										Type: &ast.StarExpr{
											X: &ast.Ident{
												Name: "int",
											},
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "selectValues",
											},
										},
										Type: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "sql",
											},
											Sel: &ast.Ident{
												Name: "SelectValues",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{
							Name: "ItemEdges",
						},
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											{
												Name: "Compartment",
											},
										},
										Type: &ast.StarExpr{
											X: &ast.Ident{
												Name: "Compartment",
											},
										},
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`json:\"compartment,omitempty\"`",
										},
									},
									{
										Names: []*ast.Ident{
											{
												Name: "loadedTypes",
											},
										},
										Type: &ast.ArrayType{
											Len: &ast.BasicLit{
												Kind:  token.INT,
												Value: "1",
											},
											Elt: &ast.Ident{
												Name: "bool",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "e",
								},
							},
							Type: &ast.Ident{
								Name: "ItemEdges",
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "CompartmentOrErr",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: &ast.Ident{
										Name: "Compartment",
									},
								},
							},
							{
								Type: &ast.Ident{
									Name: "error",
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "e",
									},
									Sel: &ast.Ident{
										Name: "Compartment",
									},
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
											&ast.SelectorExpr{
												X: &ast.Ident{
													Name: "e",
												},
												Sel: &ast.Ident{
													Name: "Compartment",
												},
											},
											&ast.Ident{
												Name: "nil",
											},
										},
									},
								},
							},
							Else: &ast.IfStmt{
								Cond: &ast.IndexExpr{
									X: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "e",
										},
										Sel: &ast.Ident{
											Name: "loadedTypes",
										},
									},
									Index: &ast.BasicLit{
										Kind:  token.INT,
										Value: "0",
									},
								},
								Body: &ast.BlockStmt{
									List: []ast.Stmt{
										&ast.ReturnStmt{
											Results: []ast.Expr{
												&ast.Ident{
													Name: "nil",
												},
												&ast.UnaryExpr{
													Op: token.AND,
													X: &ast.CompositeLit{
														Type: &ast.Ident{
															Name: "NotFoundError",
														},
														Elts: []ast.Expr{
															&ast.KeyValueExpr{
																Key: &ast.Ident{
																	Name: "label",
																},
																Value: &ast.SelectorExpr{
																	X: &ast.Ident{
																		Name: "compartment",
																	},
																	Sel: &ast.Ident{
																		Name: "Label",
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
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{
									Name: "nil",
								},
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: &ast.Ident{
											Name: "NotLoadedError",
										},
										Elts: []ast.Expr{
											&ast.KeyValueExpr{
												Key: &ast.Ident{
													Name: "edge",
												},
												Value: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"compartment\"",
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
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "scanValues",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									{
										Name: "columns",
									},
								},
								Type: &ast.ArrayType{
									Elt: &ast.Ident{
										Name: "string",
									},
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.ArrayType{
									Elt: &ast.Ident{
										Name: "any",
									},
								},
							},
							{
								Type: &ast.Ident{
									Name: "error",
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.Ident{
									Name: "values",
								},
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.Ident{
										Name: "make",
									},
									Args: []ast.Expr{
										&ast.ArrayType{
											Elt: &ast.Ident{
												Name: "any",
											},
										},
										&ast.CallExpr{
											Fun: &ast.Ident{
												Name: "len",
											},
											Args: []ast.Expr{
												&ast.Ident{
													Name: "columns",
												},
											},
										},
									},
								},
							},
						},
						&ast.RangeStmt{
							Key: &ast.Ident{
								Name: "i",
							},
							Tok: token.DEFINE,
							X: &ast.Ident{
								Name: "columns",
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.SwitchStmt{
										Tag: &ast.IndexExpr{
											X: &ast.Ident{
												Name: "columns",
											},
											Index: &ast.Ident{
												Name: "i",
											},
										},
										Body: &ast.BlockStmt{
											List: []ast.Stmt{
												&ast.CaseClause{
													List: []ast.Expr{
														&ast.SelectorExpr{
															X: &ast.Ident{
																Name: "item",
															},
															Sel: &ast.Ident{
																Name: "FieldID",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.AssignStmt{
															Lhs: []ast.Expr{
																&ast.IndexExpr{
																	X: &ast.Ident{
																		Name: "values",
																	},
																	Index: &ast.Ident{
																		Name: "i",
																	},
																},
															},
															Tok: token.ASSIGN,
															Rhs: []ast.Expr{
																&ast.CallExpr{
																	Fun: &ast.Ident{
																		Name: "new",
																	},
																	Args: []ast.Expr{
																		&ast.SelectorExpr{
																			X: &ast.Ident{
																				Name: "sql",
																			},
																			Sel: &ast.Ident{
																				Name: "NullInt64",
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
														&ast.SelectorExpr{
															X: &ast.Ident{
																Name: "item",
															},
															Sel: &ast.Ident{
																Name: "FieldName",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.AssignStmt{
															Lhs: []ast.Expr{
																&ast.IndexExpr{
																	X: &ast.Ident{
																		Name: "values",
																	},
																	Index: &ast.Ident{
																		Name: "i",
																	},
																},
															},
															Tok: token.ASSIGN,
															Rhs: []ast.Expr{
																&ast.CallExpr{
																	Fun: &ast.Ident{
																		Name: "new",
																	},
																	Args: []ast.Expr{
																		&ast.SelectorExpr{
																			X: &ast.Ident{
																				Name: "sql",
																			},
																			Sel: &ast.Ident{
																				Name: "NullString",
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
														&ast.IndexExpr{
															X: &ast.SelectorExpr{
																X: &ast.Ident{
																	Name: "item",
																},
																Sel: &ast.Ident{
																	Name: "ForeignKeys",
																},
															},
															Index: &ast.BasicLit{
																Kind:  token.INT,
																Value: "0",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.AssignStmt{
															Lhs: []ast.Expr{
																&ast.IndexExpr{
																	X: &ast.Ident{
																		Name: "values",
																	},
																	Index: &ast.Ident{
																		Name: "i",
																	},
																},
															},
															Tok: token.ASSIGN,
															Rhs: []ast.Expr{
																&ast.CallExpr{
																	Fun: &ast.Ident{
																		Name: "new",
																	},
																	Args: []ast.Expr{
																		&ast.SelectorExpr{
																			X: &ast.Ident{
																				Name: "sql",
																			},
																			Sel: &ast.Ident{
																				Name: "NullInt64",
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&ast.CaseClause{
													Body: []ast.Stmt{
														&ast.AssignStmt{
															Lhs: []ast.Expr{
																&ast.IndexExpr{
																	X: &ast.Ident{
																		Name: "values",
																	},
																	Index: &ast.Ident{
																		Name: "i",
																	},
																},
															},
															Tok: token.ASSIGN,
															Rhs: []ast.Expr{
																&ast.CallExpr{
																	Fun: &ast.Ident{
																		Name: "new",
																	},
																	Args: []ast.Expr{
																		&ast.SelectorExpr{
																			X: &ast.Ident{
																				Name: "sql",
																			},
																			Sel: &ast.Ident{
																				Name: "UnknownType",
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
								},
							},
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{
									Name: "values",
								},
								&ast.Ident{
									Name: "nil",
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "assignValues",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									{
										Name: "columns",
									},
								},
								Type: &ast.ArrayType{
									Elt: &ast.Ident{
										Name: "string",
									},
								},
							},
							{
								Names: []*ast.Ident{
									{
										Name: "values",
									},
								},
								Type: &ast.ArrayType{
									Elt: &ast.Ident{
										Name: "any",
									},
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.Ident{
									Name: "error",
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.IfStmt{
							Init: &ast.AssignStmt{
								Lhs: []ast.Expr{
									&ast.Ident{
										Name: "m",
									},
									&ast.Ident{
										Name: "n",
									},
								},
								Tok: token.DEFINE,
								Rhs: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.Ident{
											Name: "len",
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "values",
											},
										},
									},
									&ast.CallExpr{
										Fun: &ast.Ident{
											Name: "len",
										},
										Args: []ast.Expr{
											&ast.Ident{
												Name: "columns",
											},
										},
									},
								},
							},
							Cond: &ast.BinaryExpr{
								X: &ast.Ident{
									Name: "m",
								},
								Op: token.LSS,
								Y: &ast.Ident{
									Name: "n",
								},
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ReturnStmt{
										Results: []ast.Expr{
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
														Value: "\"mismatch number of scan values: %d != %d\"",
													},
													&ast.Ident{
														Name: "m",
													},
													&ast.Ident{
														Name: "n",
													},
												},
											},
										},
									},
								},
							},
						},
						&ast.RangeStmt{
							Key: &ast.Ident{
								Name: "j",
							},
							Tok: token.DEFINE,
							X: &ast.Ident{
								Name: "columns",
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.SwitchStmt{
										Tag: &ast.IndexExpr{
											X: &ast.Ident{
												Name: "columns",
											},
											Index: &ast.Ident{
												Name: "j",
											},
										},
										Body: &ast.BlockStmt{
											List: []ast.Stmt{
												&ast.CaseClause{
													List: []ast.Expr{
														&ast.SelectorExpr{
															X: &ast.Ident{
																Name: "item",
															},
															Sel: &ast.Ident{
																Name: "FieldID",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.AssignStmt{
															Lhs: []ast.Expr{
																&ast.Ident{
																	Name: "value",
																},
																&ast.Ident{
																	Name: "ok",
																},
															},
															Tok: token.DEFINE,
															Rhs: []ast.Expr{
																&ast.TypeAssertExpr{
																	X: &ast.IndexExpr{
																		X: &ast.Ident{
																			Name: "values",
																		},
																		Index: &ast.Ident{
																			Name: "j",
																		},
																	},
																	Type: &ast.StarExpr{
																		X: &ast.SelectorExpr{
																			X: &ast.Ident{
																				Name: "sql",
																			},
																			Sel: &ast.Ident{
																				Name: "NullInt64",
																			},
																		},
																	},
																},
															},
														},
														&ast.IfStmt{
															Cond: &ast.UnaryExpr{
																Op: token.NOT,
																X: &ast.Ident{
																	Name: "ok",
																},
															},
															Body: &ast.BlockStmt{
																List: []ast.Stmt{
																	&ast.ReturnStmt{
																		Results: []ast.Expr{
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
																						Value: "\"unexpected type %T for field id\"",
																					},
																					&ast.Ident{
																						Name: "value",
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
																		Name: "i",
																	},
																	Sel: &ast.Ident{
																		Name: "ID",
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
																				Name: "value",
																			},
																			Sel: &ast.Ident{
																				Name: "Int64",
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
														&ast.SelectorExpr{
															X: &ast.Ident{
																Name: "item",
															},
															Sel: &ast.Ident{
																Name: "FieldName",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.IfStmt{
															Init: &ast.AssignStmt{
																Lhs: []ast.Expr{
																	&ast.Ident{
																		Name: "value",
																	},
																	&ast.Ident{
																		Name: "ok",
																	},
																},
																Tok: token.DEFINE,
																Rhs: []ast.Expr{
																	&ast.TypeAssertExpr{
																		X: &ast.IndexExpr{
																			X: &ast.Ident{
																				Name: "values",
																			},
																			Index: &ast.Ident{
																				Name: "j",
																			},
																		},
																		Type: &ast.StarExpr{
																			X: &ast.SelectorExpr{
																				X: &ast.Ident{
																					Name: "sql",
																				},
																				Sel: &ast.Ident{
																					Name: "NullString",
																				},
																			},
																		},
																	},
																},
															},
															Cond: &ast.UnaryExpr{
																Op: token.NOT,
																X: &ast.Ident{
																	Name: "ok",
																},
															},
															Body: &ast.BlockStmt{
																List: []ast.Stmt{
																	&ast.ReturnStmt{
																		Results: []ast.Expr{
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
																						Value: "\"unexpected type %T for field name\"",
																					},
																					&ast.IndexExpr{
																						X: &ast.Ident{
																							Name: "values",
																						},
																						Index: &ast.Ident{
																							Name: "j",
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															Else: &ast.IfStmt{
																Cond: &ast.SelectorExpr{
																	X: &ast.Ident{
																		Name: "value",
																	},
																	Sel: &ast.Ident{
																		Name: "Valid",
																	},
																},
																Body: &ast.BlockStmt{
																	List: []ast.Stmt{
																		&ast.AssignStmt{
																			Lhs: []ast.Expr{
																				&ast.SelectorExpr{
																					X: &ast.Ident{
																						Name: "i",
																					},
																					Sel: &ast.Ident{
																						Name: "Name",
																					},
																				},
																			},
																			Tok: token.ASSIGN,
																			Rhs: []ast.Expr{
																				&ast.SelectorExpr{
																					X: &ast.Ident{
																						Name: "value",
																					},
																					Sel: &ast.Ident{
																						Name: "String",
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
												&ast.CaseClause{
													List: []ast.Expr{
														&ast.IndexExpr{
															X: &ast.SelectorExpr{
																X: &ast.Ident{
																	Name: "item",
																},
																Sel: &ast.Ident{
																	Name: "ForeignKeys",
																},
															},
															Index: &ast.BasicLit{
																Kind:  token.INT,
																Value: "0",
															},
														},
													},
													Body: []ast.Stmt{
														&ast.IfStmt{
															Init: &ast.AssignStmt{
																Lhs: []ast.Expr{
																	&ast.Ident{
																		Name: "value",
																	},
																	&ast.Ident{
																		Name: "ok",
																	},
																},
																Tok: token.DEFINE,
																Rhs: []ast.Expr{
																	&ast.TypeAssertExpr{
																		X: &ast.IndexExpr{
																			X: &ast.Ident{
																				Name: "values",
																			},
																			Index: &ast.Ident{
																				Name: "j",
																			},
																		},
																		Type: &ast.StarExpr{
																			X: &ast.SelectorExpr{
																				X: &ast.Ident{
																					Name: "sql",
																				},
																				Sel: &ast.Ident{
																					Name: "NullInt64",
																				},
																			},
																		},
																	},
																},
															},
															Cond: &ast.UnaryExpr{
																Op: token.NOT,
																X: &ast.Ident{
																	Name: "ok",
																},
															},
															Body: &ast.BlockStmt{
																List: []ast.Stmt{
																	&ast.ReturnStmt{
																		Results: []ast.Expr{
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
																						Value: "\"unexpected type %T for edge-field compartment_contents\"",
																					},
																					&ast.Ident{
																						Name: "value",
																					},
																				},
																			},
																		},
																	},
																},
															},
															Else: &ast.IfStmt{
																Cond: &ast.SelectorExpr{
																	X: &ast.Ident{
																		Name: "value",
																	},
																	Sel: &ast.Ident{
																		Name: "Valid",
																	},
																},
																Body: &ast.BlockStmt{
																	List: []ast.Stmt{
																		&ast.AssignStmt{
																			Lhs: []ast.Expr{
																				&ast.SelectorExpr{
																					X: &ast.Ident{
																						Name: "i",
																					},
																					Sel: &ast.Ident{
																						Name: "compartment_contents",
																					},
																				},
																			},
																			Tok: token.ASSIGN,
																			Rhs: []ast.Expr{
																				&ast.CallExpr{
																					Fun: &ast.Ident{
																						Name: "new",
																					},
																					Args: []ast.Expr{
																						&ast.Ident{
																							Name: "int",
																						},
																					},
																				},
																			},
																		},
																		&ast.AssignStmt{
																			Lhs: []ast.Expr{
																				&ast.StarExpr{
																					X: &ast.SelectorExpr{
																						X: &ast.Ident{
																							Name: "i",
																						},
																						Sel: &ast.Ident{
																							Name: "compartment_contents",
																						},
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
																								Name: "value",
																							},
																							Sel: &ast.Ident{
																								Name: "Int64",
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
												},
												&ast.CaseClause{
													Body: []ast.Stmt{
														&ast.ExprStmt{
															X: &ast.CallExpr{
																Fun: &ast.SelectorExpr{
																	X: &ast.SelectorExpr{
																		X: &ast.Ident{
																			Name: "i",
																		},
																		Sel: &ast.Ident{
																			Name: "selectValues",
																		},
																	},
																	Sel: &ast.Ident{
																		Name: "Set",
																	},
																},
																Args: []ast.Expr{
																	&ast.IndexExpr{
																		X: &ast.Ident{
																			Name: "columns",
																		},
																		Index: &ast.Ident{
																			Name: "j",
																		},
																	},
																	&ast.IndexExpr{
																		X: &ast.Ident{
																			Name: "values",
																		},
																		Index: &ast.Ident{
																			Name: "j",
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
							},
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{
									Name: "nil",
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "Value",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									{
										Name: "name",
									},
								},
								Type: &ast.Ident{
									Name: "string",
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "ent",
									},
									Sel: &ast.Ident{
										Name: "Value",
									},
								},
							},
							{
								Type: &ast.Ident{
									Name: "error",
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "i",
											},
											Sel: &ast.Ident{
												Name: "selectValues",
											},
										},
										Sel: &ast.Ident{
											Name: "Get",
										},
									},
									Args: []ast.Expr{
										&ast.Ident{
											Name: "name",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "QueryCompartment",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: &ast.Ident{
										Name: "CompartmentQuery",
									},
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.CallExpr{
											Fun: &ast.Ident{
												Name: "NewItemClient",
											},
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X: &ast.Ident{
														Name: "i",
													},
													Sel: &ast.Ident{
														Name: "config",
													},
												},
											},
										},
										Sel: &ast.Ident{
											Name: "QueryCompartment",
										},
									},
									Args: []ast.Expr{
										&ast.Ident{
											Name: "i",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "Update",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: &ast.Ident{
										Name: "ItemUpdateOne",
									},
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.CallExpr{
											Fun: &ast.Ident{
												Name: "NewItemClient",
											},
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X: &ast.Ident{
														Name: "i",
													},
													Sel: &ast.Ident{
														Name: "config",
													},
												},
											},
										},
										Sel: &ast.Ident{
											Name: "UpdateOne",
										},
									},
									Args: []ast.Expr{
										&ast.Ident{
											Name: "i",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "Unwrap",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: &ast.Ident{
										Name: "Item",
									},
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.Ident{
									Name: "_tx",
								},
								&ast.Ident{
									Name: "ok",
								},
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.TypeAssertExpr{
									X: &ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "i",
											},
											Sel: &ast.Ident{
												Name: "config",
											},
										},
										Sel: &ast.Ident{
											Name: "driver",
										},
									},
									Type: &ast.StarExpr{
										X: &ast.Ident{
											Name: "txDriver",
										},
									},
								},
							},
						},
						&ast.IfStmt{
							Cond: &ast.UnaryExpr{
								Op: token.NOT,
								X: &ast.Ident{
									Name: "ok",
								},
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.Ident{
												Name: "panic",
											},
											Args: []ast.Expr{
												&ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"ent: Item is not a transactional entity\"",
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
									X: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "i",
										},
										Sel: &ast.Ident{
											Name: "config",
										},
									},
									Sel: &ast.Ident{
										Name: "driver",
									},
								},
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.SelectorExpr{
									X: &ast.Ident{
										Name: "_tx",
									},
									Sel: &ast.Ident{
										Name: "drv",
									},
								},
							},
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{
									Name: "i",
								},
							},
						},
					},
				},
			},
			&ast.FuncDecl{
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								{
									Name: "i",
								},
							},
							Type: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
				Name: &ast.Ident{
					Name: "String",
				},
				Type: &ast.FuncType{
					Params: &ast.FieldList{},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.Ident{
									Name: "string",
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
											{
												Name: "builder",
											},
										},
										Type: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "strings",
											},
											Sel: &ast.Ident{
												Name: "Builder",
											},
										},
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "builder",
									},
									Sel: &ast.Ident{
										Name: "WriteString",
									},
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: "\"Item(\"",
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "builder",
									},
									Sel: &ast.Ident{
										Name: "WriteString",
									},
								},
								Args: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "fmt",
											},
											Sel: &ast.Ident{
												Name: "Sprintf",
											},
										},
										Args: []ast.Expr{
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"id=%v, \"",
											},
											&ast.SelectorExpr{
												X: &ast.Ident{
													Name: "i",
												},
												Sel: &ast.Ident{
													Name: "ID",
												},
											},
										},
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "builder",
									},
									Sel: &ast.Ident{
										Name: "WriteString",
									},
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: "\"name=\"",
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "builder",
									},
									Sel: &ast.Ident{
										Name: "WriteString",
									},
								},
								Args: []ast.Expr{
									&ast.SelectorExpr{
										X: &ast.Ident{
											Name: "i",
										},
										Sel: &ast.Ident{
											Name: "Name",
										},
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "builder",
									},
									Sel: &ast.Ident{
										Name: "WriteByte",
									},
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.CHAR,
										Value: "')'",
									},
								},
							},
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "builder",
										},
										Sel: &ast.Ident{
											Name: "String",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{
							Name: "Items",
						},
						Type: &ast.ArrayType{
							Elt: &ast.StarExpr{
								X: &ast.Ident{
									Name: "Item",
								},
							},
						},
					},
				},
			},
		},
		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"fmt\"",
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"strings\"",
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"tempvs/ent/compartment\"",
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"tempvs/ent/item\"",
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"entgo.io/ent\"",
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"entgo.io/ent/dialect/sql\"",
				},
			},
		},
		Comments: []*ast.CommentGroup{
			{
				List: []*ast.Comment{
					{
						Text: "// Code generated by ent, DO NOT EDIT.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Item is the model entity for the Item schema.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// ID of the ent.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Name holds the value of the \"name\" field.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Edges holds the relations/edges for other nodes in the graph.",
					},
					{
						Text: "// The values are being populated by the ItemQuery when eager-loading is set.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// ItemEdges holds the relations/edges for other nodes in the graph.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Compartment holds the value of the compartment edge.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// loadedTypes holds the information for reporting if a",
					},
					{
						Text: "// type was loaded (or requested) in eager-loading or not.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{

						Text: "// CompartmentOrErr returns the Compartment value or an error if the edge",
					},
					{

						Text: "// was not loaded in eager-loading, or loaded but was not found.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// scanValues returns the types for scanning values from sql.Rows.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// compartment_contents",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// assignValues assigns the values that were returned from sql.Rows (after scanning)",
					},
					{
						Text: "// to the Item fields.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Value returns the ent.Value that was dynamically selected and assigned to the Item.",
					},
					{
						Text: "// This includes values selected through modifiers, order, etc.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// QueryCompartment queries the \"compartment\" edge of the Item entity.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Update returns a builder for updating this Item.",
					},
					{
						Text: "// Note that you need to call Item.Unwrap() before calling this method if this Item",
					},
					{
						Text: "// was returned from a transaction, and the transaction was committed or rolled back.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,",
					},
					{
						Text: "// so that all future queries will be executed through the driver which created the transaction.",
					},
				},
			},
			{
				List: []*ast.Comment{
					{
						Text: "// String implements the fmt.Stringer.",
					},
				},
			},
		},
	}

}
