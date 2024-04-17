package model

type LatestShortCut struct {
	SequenceId     int64  `structs:"sequence_id"`
	LastestContent string `structs:"lastest_content"`
	Cnt            int    `structs:"cnt"`
}
