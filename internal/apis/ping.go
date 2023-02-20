package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"ydsd_gin/internal/code"
	"ydsd_gin/internal/pkg/response"
)

type data struct {
	Info string `json:"info"`
}

func Ping(c *gin.Context) *response.Response {
	d := data{Info: "this is ping info " + viper.GetString("application.domain")}
	return response.ResponseOk(code.AdminCreateError, d)

}
