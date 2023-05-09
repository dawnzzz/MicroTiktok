package main

import (
	id_generator "github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator/idgeneratorservice"
	"log"
)

func main() {
	// 初始化id generator config
	InitIdGeneratorConfig()

	svr := id_generator.NewServer(new(IdGeneratorServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
