package jwt

import (
	"github.com/gtkit/encry/jwt"

	"my_gin/config"
)

func InitJwt() {
	jwt.SetSignKey(config.GetString("jwt.secret"))
}
