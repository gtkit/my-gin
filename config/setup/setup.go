package setup

import (
	"my_gin/config"
)

func Initialize() {
	// 加载配制信息
	config.LoadConfig()
}
