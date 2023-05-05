package main

import (
	relation "github.com/dawnzzz/MicroTiktok/kitex_gen/relation/socialityservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(SocialityServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
