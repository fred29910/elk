package pk

type Handle interface {
	CmdHeartbeat(pb.Context, *proto.HeartbeatReq) error
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
	}
	return &gr, nil
}
