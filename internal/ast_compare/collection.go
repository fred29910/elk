package ast_compare

import "go/ast"

// 收集不同数据，分类
// 1. imports
// 2. global var/static
// 3. global func
// 4. struct
// 		4.1 struct func

// Compare , 对比两个文件
func (m *Compare) Compare() error {

	fileAst := FileAst{}

	fileAst.Package = PackageCompare{
		old: m.oldNodes.Name.Name,
		new: m.newNodes.Name.Name,
	}
	fileAst.Import = m.importCompare()
	return nil
}

func (m *Compare) importCompare() ImportCompare {

	newNodeMap := collectionImports(m.newNodes)
	oldNodeMap := collectionImports(m.oldNodes)

	var old []Import
	conflict := map[string][2]Import{}
	for s, i := range oldNodeMap {
		if item, ok := newNodeMap[s]; ok {
			if item.alias != i.alias {
				conflict[s] = [2]Import{i, item}
			} else {
				old = append(old, i)
			}
		} else {
			old = append(old, i)
		}
	}
	var newIm []Import
	for s, i := range newNodeMap {
		if _, ok := oldNodeMap[s]; !ok {
			newIm = append(newIm, i)
		}
	}
	return ImportCompare{
		new:      newIm,
		old:      old,
		conflict: conflict,
	}
}

func collectionImports(m *ast.File) map[string]Import {
	newNodeMap := map[string]Import{}
	for _, spec := range m.Imports {
		pkgPath := spec.Path.Value
		aliasName := ""
		if spec.Name != nil {
			aliasName = spec.Name.Name
		}
		newNodeMap[pkgPath] = Import{
			alias: aliasName,
			pkg:   pkgPath,
		}
	}
	return newNodeMap
}
