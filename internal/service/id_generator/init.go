package main

import (
	"github.com/bwmarrin/snowflake"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/service/id_generator/config"
	"math/rand"
	"time"
)

func InitIdGeneratorConfig() {
	// 暂时随机初始化一个node id
	rand.Seed(time.Now().Unix())
	config.IdGeneratorConfigObj = &config.IdGeneratorConfig{
		Port:   global.DefaultRpcIdGeneratorServicePort,
		NodeID: rand.Int63() % (2 << snowflake.NodeBits),
	}
}
