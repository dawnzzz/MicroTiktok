package config

import "time"

type AuthenticationConfig struct {
	TokenExpireDuration time.Duration // 过期时间
	Secret              string        // jwt密钥
}

var AuthenticationConfigObj *AuthenticationConfig

func init() {
	AuthenticationConfigObj = &AuthenticationConfig{
		TokenExpireDuration: time.Hour * 24,
		Secret:              "dawnDawnDawn",
	}
}
