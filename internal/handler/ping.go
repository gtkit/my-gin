package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/model"
)

func (h *handler) Ping(c *gin.Context) {

	var (
		m   *model.AssistantMember
		err error
	)

	m, err = h.repository.Ping()
	if err != nil {
		// resp.Error(c, err)
	}
	fmt.Println("------------ping ----------", m)
	msg := NewMessage(1, []byte("this is xiaozhaofu test protoc 2"))
	pbmsg := msg.ToProtoc()
	fmt.Printf("----pbmsg2: %+v\n", string(pbmsg))
	res, err := msg.FromProtoc(pbmsg)
	if err != nil {
		fmt.Println("---msg FromProtoc error:", err)
	}
	fmt.Printf("----decode pbmsg: %+v\n", res)
	// c.ProtoBuf(200, &msg)
	// data := res

	// response.Ok(c, m)
	// c.JSON(200, gin.H{
	// 	"MsgType":    2,
	// 	"MsgContent": string(res.MsgContent),
	// })
	c.ProtoBuf(200, res)

}

func (h *handler) DoPing(c *gin.Context) {
	fmt.Println("-----do ping -----")
}
