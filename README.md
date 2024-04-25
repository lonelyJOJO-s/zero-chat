# zero-chat
this is a project of socical meidia, using go-zero to build it
the Preliminary services are as belows:
1. usercenter
2. chat
3. video
4. twiiter
5. storage

TODO:
- [ ] redis执行cnt和content记录中加入事务机制 


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
