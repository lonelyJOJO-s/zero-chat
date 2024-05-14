# zero-chat
this is a project of socical media, using go-zero to build it
## function
- login/register
- user/group
- chat: text,pic,video, audio
- instant chat: video, audio, creen-sharing
## back-end
### micro-service:
- user server
- chat server(need to break into Instant Message and non-instant Message)
### tech stack
- go-zero 
- redis
- kafka
- mysql
- aliyun oss
- aliyun storetable
- etcd
- jaeger
- nginx
- docker
- websocket
- webrtc

## front-end
link: https://github.com/lonelyJOJO-s/zero-chat-web


Note:
1. 阿里云timeline-message结构体修改：
```
var DefaultStreamAdapter = &StreamMessageAdapter{
	IdKey:        "Id",
	ContentKey:   "Content",
	TimestampKey: "Timestamp",
	AttrPrefix:   "",
}
```
