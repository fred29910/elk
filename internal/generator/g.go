package generator

import "github.com/masseelch/elk/internal/parser"

type Package struct {
	Name    string   // 由 @package 定义与声明
	Imports []string // 需要引入的包

	Handle map[int32]HandleInfo

	Coms map[string]*parser.Msg // 当前包需要的消息映射

	Cms map[string]*parser.Msg

	OutDir        string // 可以是相对路径， 也可以是绝对路径， 控制handle 的输出位置。 由命令行参数控制， 如果不填写，默认输出到common包位置
	GopackageName string
}

type HandleInfo struct {
	MsgID    string
	MsgName  string
	IsCommon bool
}

type MsgDefine struct {
	Msgs map[string]int32
	Path string
	Pkg  string
	Cms  map[string]*parser.Msg
}
