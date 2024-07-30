package model

const (
	TimelineId   = "TimelineId"
	Sequence     = "Sequence"
	Id           = "Id"
	Conversation = "Conversation"
	Content      = "Content"
	MsgType      = "MsgType"
	SendTime     = "Timestamp"
	Sender       = "Sender"
	Type         = "Type"
	File         = "File"
	UserId       = "user_id"
)

func NewPrimaryKeys(timeLineId string, sequence int64) map[string]any {
	return map[string]any{
		TimelineId: timeLineId,
		Sequence:   sequence,
	}
}

type Opt func(*map[string]any)

func WithId(val string) Opt {
	return func(c *map[string]any) {
		(*c)[Id] = val
	}
}

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

func WithFile(val string) Opt {
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
