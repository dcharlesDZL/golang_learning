## Async impl
- 利用两个带缓冲的channel来实现任务的异步处理
```go
type Service struct {
	RecChannel chan int
	ReqChannel chan int
}
```
- ReqChannel发送任务的channel
- RecChannel接收结果的channel
利用通道缓冲的个数来实现任务并发数量的控制.