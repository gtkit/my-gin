package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gtkit/json"
)

func (h *handler) Ping(c *gin.Context) {

	// m, err = h.repository.Ping()
	// if err != nil {
	// 	// resp.Error(c, err)
	// }
	// fmt.Println("------------ping ----------", m)
	msg := NewMessage(1, []byte("this is xiaozhaofu test protoc 2"))
	pbmsg := msg.ToProtoc()
	fmt.Printf("----pbmsg2: %+v\n", &pbmsg)
	// res, err := msg.FromProtoc(pbmsg)
	// if err != nil {
	// 	fmt.Println("---msg FromProtoc error:", err)
	// }
	// fmt.Printf("----decode pbmsg: %+v\n", res)
	// c.ProtoBuf(200, &msg)
	// data := res

	// response.Ok(c, m)
	// c.JSON(200, gin.H{
	// 	"MsgType":    2,
	// 	"MsgContent": string(res.MsgContent),
	// })
	// c.ProtoBuf(200, res)
	var structMsg = struct {
		MsgType    int    `json:"msg_type"`
		MsgContent string `json:"msg_content"`
		MsgFrom    string `json:"msg_from"`
	}{
		MsgType:    1,
		MsgContent: "this is xiaozhaofu test protoc 1",
		MsgFrom:    "xiaozhaofu",
	}
	jsonMsg, _ := json.Marshal(structMsg)
	fmt.Println("----jsonMsg:", string(jsonMsg))
	json.CheckJSON()

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "pong",
	})

}

func (h *handler) DoPing(c *gin.Context) {
	fmt.Println("-----do ping -----")
}
