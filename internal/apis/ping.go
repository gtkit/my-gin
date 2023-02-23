package apis

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/config"
	"ydsd_gin/internal/model"
	"ydsd_gin/internal/pkg/response"
)

type data struct {
	Info string `json:"info"`
}

func (h *handler) Ping(c *gin.Context) {
	d := data{Info: "this is ping info " + config.GetString("application.domain")}
	var m *model.AssistantMember

	m, err := h.repository.Ping()
	if err != nil {
		h.log.Info("ping err-----", err)
	}
	// fmt.Println("member info------", m)
	h.log.Info("member info------", m)
	response.Ok(c, d)

}
