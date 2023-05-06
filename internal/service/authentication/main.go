package main

import (
	authentication "github.com/dawnzzz/MicroTiktok/kitex_gen/authentication/authenticationservice"
	"log"
)

func main() {
	svr := authentication.NewServer(new(AuthenticationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
