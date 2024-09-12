// @Author 2023/11/3 15:10:00
package handler

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"ydsd_gin/protobuf/pb"
)

// var _ IProtoc = (*MessageBase)(nil)

// IProtoc 用于规范protoc序列化
type IProtoc interface {
	ToProtoc() []byte
	FromProtoc([]byte) error
}

// // FromProtoc 接受用于反序列化的字节数据和符合IProtoc的类型，功能是将msg数据赋值初始化
// func FromProtoc[T IProtoc](data []byte, msg T) error {
// 	err := msg.FromProtoc(data)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func ToProtoc[T IProtoc](msg T) []byte {
// 	return msg.ToProtoc()
//
// }

type MessageBase struct {
	Impl *pb.BaseMessage
}

func NewMessage(msgType int32, msgContent []byte) *MessageBase {
	return &MessageBase{
		Impl: &pb.BaseMessage{
			MsgType:    msgType,
			MsgContent: msgContent,
		},
	}
}

func (m *MessageBase) ToProtoc() []byte {
	msg, err := proto.Marshal(m.Impl)
	if err != nil {
		return nil
	}
	return msg
}

func (m *MessageBase) FromProtoc(msg []byte) (*pb.BaseMessage, error) {
	m.Impl = &pb.BaseMessage{}
	err := proto.Unmarshal(msg, m.Impl)
	if err != nil {
		return nil, err
	}
	fmt.Println("---proto.Unmarshal ", m.Impl)
	return m.Impl, nil
}
