// @Author xiaozhaofu 2023/7/12 21:18:00
package examples

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/verify"

	"ydsd_gin/internal/pkg/response"
)

type OrderParams struct {
	OrderShopId string `form:"order_shop_id"  binding:"required,numeric"`
	Type        string `form:"type"  binding:"required,numeric"`
}

func Order(c *gin.Context) {
	// 1. 验证结构体参数
	var p OrderParams
	if err := c.ShouldBindQuery(&p); err != nil {
		resp.Error(c, verify.StructErr(err))
		return
	}

	// 2. 验证单个参数
	shopId := c.Query("order_shop_id")
	err := verify.VarField(shopId, "required,numeric")
	if err != nil {
		resp.Error(c, verify.ErrorInfo("order_shop_id", err))
		return
	}

	payType := c.Query("type")
	err = verify.VarField(payType, "required,numeric")
	if err != nil {
		resp.Error(c, verify.ErrorInfo("type", err))
		return
	}

	// 3. 验证map参数
	user := map[string]interface{}{
		"name":  "htttereeee",
		"emain": "hddd@google.com",
		// "email": "1",
	}

	rules := map[string]interface{}{
		"name":  "required,min=8,max=15",
		"email": "omitempty,email",
	}
	// 此处err为map[string]any
	if err := verify.VarMap(user, rules); len(err) > 0 {
		log.Printf("verify error: %+v, len: %d\n", err, len(err))
		resp.Error(c, verify.MapErr(err))
		return
	}

	resp.Ok(c, gin.H{
		"order_token": "",
	})

}
