package setup

import (
	"ydsd_gin/config"
)

func Initialize() {
	// 加载配制信息
	config.LoadConfig()
}
