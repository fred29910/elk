package pk

import (
	"context"
	"fmt"
	"github.com/masseelch/elk/test/proto"
	"github.com/masseelch/elk/test/proto/pk"
	pg "google.golang.org/protobuf/proto"
)

func Data(hand Handle, tp proto.MSG, data []byte) error {

	if tp == proto.MSG_CMD_HEARTBEAT_REQ {
		// 创建一个空的 ExampleMessage
		msg := &proto.HeartbeatReq{}

		// 将 []byte 数据反序列化为 ExampleMessage
		if err := pg.Unmarshal(data, msg); err != nil {
			return err
		}
		return hand.Heartbeat(nil, msg)
	}

	return fmt.Errorf("cat not parser data")
}

type GameReply struct {
	MsgID int
	Data  []byte
}

func Out(msg pg.Message, opts func([]byte) []byte) (*GameReply, error) {
	var gr GameReply
	bs, err := pg.Marshal(msg)
	if err != nil {
		return nil, err
	}
	if opts != nil {
		bs = opts(bs)
	}
	gr.Data = bs

	switch msg.(type) {
	case *proto.Echox:
		gr.MsgID = int(proto.MSG_CMD_HELLO_REQ)
	case *pk.Hne:
		gr.MsgID = int(proto.MSG_CMD_HELLO_REQ)
	default:
		return nil, fmt.Errorf("cat not support data to out")
	}
	return &gr, nil
}

type Handle interface {
	Heartbeat(context.Context, *proto.HeartbeatReq) error
}
