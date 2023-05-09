package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/service/authentication/config"
	authentication "github.com/dawnzzz/MicroTiktok/kitex_gen/authentication/authenticationservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	InitAuthenticationConfig()

	// 服务注册
	r, err := etcd.NewEtcdRegistry(global.EtcdEndpoints)
	if err != nil {
		panic(err)
	}

	//rpc info
	info := &rpcinfo.EndpointBasicInfo{
		ServiceName: global.RpcAuthenticationServiceName,
	}

	svr := authentication.NewServer(new(AuthenticationServiceImpl),
		server.WithServerBasicInfo(info), server.WithRegistry(r),
		server.WithServiceAddr(utils.NewNetAddr("tcp", fmt.Sprintf(":%v", config.AuthenticationConfigObj.Port))),
		server.WithMuxTransport(),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
