package constant

// content_type
const (
	TEXT = iota
	FILE
	IMAGE
	AUDIO
	VIDEO
	AUDIO_ONLINE
	VIDEO_ONLINE
)

// chat_type
const (
	SINGLE = iota
	GROUP
)

const (
	HEART_BEAT = "heartbeat"
	PONG       = "pong"
)

const (
	DISTRIBUTE_PREFIX = "distribute:"
	MAX_RETRY         = 5
	RETRY_INTERVAL    = 200
)
