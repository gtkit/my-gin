package setup

import (
	"ydsd_gin/config"
)

func Initialize() {
	addapp()
	adddb()    // 数据库
	addjwt()   // jwt
	addlog()   // 日志
	addnews()  // 消息发送
	addredis() // redis

	// 加载配制信息
	config.LoadConfig()
}
