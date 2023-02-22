package apis

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/config"
	"ydsd_gin/internal/pkg/response"
)

type data struct {
	Info string `json:"info"`
}

func Ping(c *gin.Context) {
	d := data{Info: "this is ping info " + config.GetString("application.domain")}
	response.Ok(c, d)
	// response.Fail(c)
	// err := goerr.Custom("参数错误")
	// if err != nil {
	// 	response.Error(c, goerr.ErrParams, err)
	// }

}
