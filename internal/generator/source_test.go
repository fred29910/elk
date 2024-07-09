package generator

import (
	"github.com/masseelch/elk/internal/parser"
	"go/ast"
	"testing"
)

func Test_handleGen(t *testing.T) {
	type args struct {
		cn MsgDefine
		pg Package
		f  *ast.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "ssv",
			args: args{
				cn: MsgDefine{
					Msgs: map[string]int32{
						"Min_MSG":             0,
						"CMD_HEARTBEAT_REQ":   1,
						"CMD_HEARTBEAT_REPLY": 2,
						"CMD_HELLO_REQ":       3,
					},
					Path: "hello",
					Pkg:  "github.com/masseelch/elk/test/proto",
					Cms: map[string]*parser.Msg{
						"CMD_HEARTBEAT_REQ": &parser.Msg{
							ID:     "CMD_HEARTBEAT_REQ",
							NoBind: false,
							Handle: true,
						},
						"CMD_HELLO_REQ": {
							ID:     "CMD_HELLO_REQ",
							NoBind: false,
							Handle: false,
						},
					},
				},
				pg: Package{
					Name:    "pk",
					Imports: nil,
					Handle:  nil,
					Coms: map[string]*parser.Msg{
						"CMD_HEARTBEAT_REQ": {
							ID:     "CMD_HEARTBEAT_REQ",
							NoBind: false,
							Handle: false,
						},
					},
					Cms: map[string]*parser.Msg{
						"CMD_HEARTBEAT_REPLY": {
							ID:     "CMD_HEARTBEAT_REPLY",
							NoBind: false,
							Handle: false,
						},
					},
					OutDir:        "",
					GopackageName: "github.com/masseelch/elk/test/proto/pk",
				},
				f: &ast.File{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleGen(tt.args.cn, tt.args.pg, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("handleGen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_msgIDToHandleName(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sv ssa",
			args: args{
				"CMD_HEARTBEAT_REQ",
			},
			want: "CmdHeartbeat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := msgIDToHandleName(tt.args.id); got != tt.want {
				t.Errorf("msgIDToHandleName() = %v, want %v", got, tt.want)
			}
		})
	}
}
