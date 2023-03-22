package jwt

import (
	"github.com/gtkit/encry/jwt"

	"ydsd_gin/config"
)

func InitJwt() {
	jwt.SetSignKey(config.GetString("jwt.secret"))
}
