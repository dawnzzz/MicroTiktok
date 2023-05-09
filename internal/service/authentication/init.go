package main

import (
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/service/authentication/config"
	"time"
)

func InitAuthenticationConfig() {
	config.AuthenticationConfigObj = &config.AuthenticationConfig{
		Port:                global.DefaultRpcAuthenticationServicePort,
		TokenExpireDuration: time.Hour * 24,
		Secret:              "dawnDawnDawn",
	}
}
