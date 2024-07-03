package parser

import (
	"reflect"
	"testing"
)

func TestParserMsg(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Msg
	}{

		{
			name: "first line",
			args: `@msg xiasmddam`,
			want: &Msg{
				ID:     "xiasmddam",
				NoBind: false,
			},
		},
		{
			name: "first line with no bind",
			args: `@msg xiasmddam
@nobind`,
			want: &Msg{
				ID:     "xiasmddam",
				NoBind: true,
			},
		},
		{
			name: "not first line with no bind",
			args: `xiamsdiaxmasmdafds
@msg xiasmddam
@nobind`,
			want: &Msg{
				ID:     "xiasmddam",
				NoBind: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MsgParser(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonsParser(t *testing.T) {

	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "first line",
			args: "@common xiasmddam",
			want: []string{"xiasmddam"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonsParser(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonsParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
