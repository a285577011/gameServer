package msg

import (
	"github.com/name5566/leaf/network/binaryjson"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = binaryjson.NewProcessor()

func init() {

}
