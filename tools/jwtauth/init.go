package jwt

import (
	"gitlab.superjq.com/go-tools/encry/jwt"

	"ydsd_gin/config"
)

func InitJwt() {
	jwt.SetSignKey(config.GetString("jwt.secret"))
}
