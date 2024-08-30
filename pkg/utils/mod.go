package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetGoModPath 返回当前项目的 go.mod 文件的绝对路径
func GetGoModPath() (string, error) {
	// 获取当前工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 在工作目录及其父目录中查找 go.mod 文件
	for {
		modPath := filepath.Join(workingDir, "go.mod")
		if _, err := os.Stat(modPath); err == nil {
			// 如果找到 go.mod 文件，返回其绝对路径
			return filepath.Abs(modPath)
		}

		// 获取父目录
		parentDir := filepath.Dir(workingDir)

		// 如果当前目录已经是根目录，则退出循环
		if parentDir == workingDir {
			break
		}

		// 继续查找上一级目录
		workingDir = parentDir
	}

	return "", fmt.Errorf("go.mod not found in the current directory or any parent directory")
}

// GetModulePath 解析 go.mod 文件并返回模块路径
func GetModulePath(goModDir string) (string, error) {
	goModPath := filepath.Join(goModDir, "go.mod")
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module path not found in go.mod")
}

// GetImportPath 返回指定目录的 Go import 路径
func GetImportPath(targetDir string) (string, error) {
	goModDir, err := GetGoModPath()
	if err != nil {
		return "", err
	}

	modulePath, err := GetModulePath(goModDir)
	if err != nil {
		return "", err
	}

	// 计算目标目录相对于 go.mod 文件所在目录的相对路径
	relPath, err := filepath.Rel(goModDir, targetDir)
	if err != nil {
		return "", err
	}

	// 拼接模块路径和相对路径，形成完整的 import 路径
	if relPath == "." {
		return modulePath, nil
	}
	return filepath.ToSlash(filepath.Join(modulePath, relPath)), nil
}
