package ast_compare

import (
	"fmt"
	"go/ast"
)

type Compare struct {
	oldNodes *ast.File
	newNodes *ast.File
}

func NewCompare(srcOld any, srcNew any) (*Compare, error) {
	var oldNodes, newNodes *ast.File
	var err error
	// 如果是[]byte , 那么直接解析， 如果是string, 当文件处理
	if fileName, ok := srcOld.(string); ok {
		_, oldNodes, err = parseFile(fileName)
		if err != nil {
			return nil, err
		}
	} else if fileBs, ok2 := srcOld.([]byte); ok2 {
		_, oldNodes, err = parseBytes(fileBs)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("cat not get file source %V", srcOld)
	}

	if fileName, ok := srcNew.(string); ok {
		_, newNodes, err = parseFile(fileName)
		if err != nil {
			return nil, err
		}
	} else if fileBs, ok2 := srcNew.([]byte); ok2 {
		_, newNodes, err = parseBytes(fileBs)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("cat not get file source %V", srcNew)
	}
	//fmt.Println(oldNodes, newNodes)
	return &Compare{oldNodes: oldNodes, newNodes: newNodes}, nil
}

// Merge , 合并两个文件
func (m *Compare) Merge() error {

	return nil
}
