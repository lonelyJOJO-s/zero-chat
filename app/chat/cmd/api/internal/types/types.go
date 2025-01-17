// Code generated by goctl. DO NOT EDIT.
package types

type Message struct {
	From         int64  `json:"from"`
	Content      string `json:"content"`
	SendTime     int64  `json:"send_time"`
	ContentType  int32  `json:"content_type"` // text or audio or video
	File         string `json:"file"`
	ChatType     int32  `json:"chat_type"` // group or single
	FromUsername string `json:"fromUsername"`
	Avatar       string `json:"avatar"`
}

type GetHistoryMessageReq struct {
	ChatType int32 `query:"chat_type"`
	Id       int32 `query:"id"`
	Cnt      int64 `query:"cnt"`
	Offset   int64 `query:"offset"`
}

type GetHistoryMessageResp struct {
	Msgs          []Message `json:"msgs"`
	NextReadIndex int64     `json:"next_read_index"`
}

type SearchHistoryMessageReq struct {
	ChatType int32  `query:"chat_type"`
	Id       int32  `query:"id"`
	Limit    int64  `query:"limit"`
	Offset   int64  `query:"offset"`
	Keyword  string `query:"keyword"`
}

type SearchHistoryMessageResp struct {
	Msgs []Message `json:"msgs"`
}
