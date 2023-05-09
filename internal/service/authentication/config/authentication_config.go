package config

import "time"

type AuthenticationConfig struct {
	Port int

	TokenExpireDuration time.Duration // 过期时间
	Secret              string        // jwt密钥
}

var AuthenticationConfigObj *AuthenticationConfig
