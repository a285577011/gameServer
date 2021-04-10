package msg

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&UserRedDot{}, "u.q")
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type UserRedDot struct {
	Common
}
