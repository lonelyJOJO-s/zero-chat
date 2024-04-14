package model

const (
	TimeLineId   = "timeline_id"
	SequenceId   = "sequence_id"
	Conversation = "conversation"
	Content      = "content"
	MsgType      = "msg_type"
	SendTime     = "send_time"
	Sender       = "sender"
	Type         = "type"
	File         = "file"
	UserId       = "user_id"
)

func NewPrimaryKeys(timeLineId string, sequenceId int64) map[string]any {
	return map[string]any{
		TimeLineId: timeLineId,
		SequenceId: sequenceId,
	}
}

type Opt func(*map[string]any)

func WithConversation(val string) Opt {
	return func(c *map[string]any) {
		(*c)[Conversation] = val
	}
}

func WithSender(val int64) Opt {
	return func(c *map[string]any) {
		(*c)[Sender] = val
	}
}

func WithContent(val string) Opt {
	return func(c *map[string]any) {
		(*c)[Content] = val
	}
}

func WithMsgType(val int32) Opt {
	return func(c *map[string]any) {
		(*c)[MsgType] = int64(val)
	}
}

func WithType(val int32) Opt {
	return func(c *map[string]any) {
		(*c)[Type] = int64(val)
	}
}

func WithFile(val []byte) Opt {
	return func(c *map[string]any) {
		(*c)[File] = val
	}
}

func WithSendTime(val int64) Opt {
	return func(c *map[string]any) {
		(*c)[SendTime] = val
	}
}

func WithUserId(val int64) Opt {
	return func(c *map[string]any) {
		(*c)[UserId] = val
	}
}

func NewColumns(opts ...Opt) map[string]any {
	columns := map[string]any{}
	for _, opt := range opts {
		opt(&columns)
	}
	return columns
}
