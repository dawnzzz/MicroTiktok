package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/service/user/config"
	user "github.com/dawnzzz/MicroTiktok/kitex_gen/user/userservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	InitUserConfig()

	InitDB()

	InitRpcClient()

	// 服务注册
	r, err := etcd.NewEtcdRegistry(global.EtcdEndpoints)
	if err != nil {
		panic(err)
	}

	//rpc info
	info := &rpcinfo.EndpointBasicInfo{
		ServiceName: global.RpcUserServiceName,
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(info), server.WithRegistry(r),
		server.WithServiceAddr(utils.NewNetAddr("tcp", fmt.Sprintf(":%v", config.UserConfigObj.Port))),
		server.WithMuxTransport(),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
