package main

import (
	user "github.com/dawnzzz/MicroTiktok/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	InitDB()

	InitRpcClient()

	if err != nil {
		log.Println(err.Error())
	}
}
