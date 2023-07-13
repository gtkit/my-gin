// @Author xiaozhaofu 2023/7/12 21:18:00
package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/gtkit/verify"

	"ydsd_gin/internal/pkg/response"
)

type OrderParams struct {
	OrderShopId string `form:"order_shop_id"  binding:"required,numeric"`
	Type        string `form:"type"  binding:"required,numeric"`
}

func Order(c *gin.Context) {
	var p OrderParams
	if err := c.ShouldBindQuery(&p); err != nil {

		response.Error(c, verify.ErrorStruct(err))
		return
	}
	// shopId := c.Query("order_shop_id")
	// err := verify.Validate().Var(shopId, "required,numeric")
	// if err != nil {
	// 	response.Error(c, verify.ErrorInfo("order_shop_id", err))
	// 	return
	// }
	//
	// payType := c.Query("type")
	// err = verify.Validate().Var(payType, "required,numeric")
	// if err != nil {
	// 	response.Error(c, verify.ErrorInfo("type", err))
	// 	return
	// }
	//

	response.Ok(c, gin.H{
		"order_token": "",
	})

}
