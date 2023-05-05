package main

import (
	society "github.com/dawnzzz/MicroTiktok/kitex_gen/society/societyservice"
	"log"
)

func main() {
	svr := society.NewServer(new(SocietyServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
