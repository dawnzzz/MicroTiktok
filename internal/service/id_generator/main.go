package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/internal/service/id_generator/config"
	id_generator "github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator/idgeneratorservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	// 初始化id generator config
	InitIdGeneratorConfig()

	// 服务注册
	r, err := etcd.NewEtcdRegistry(global.EtcdEndpoints)
	if err != nil {
		panic(err)
	}

	//rpc info
	info := &rpcinfo.EndpointBasicInfo{
		ServiceName: global.RpcIdGeneratorServiceName,
		Tags:        map[string]string{},
	}

	svr := id_generator.NewServer(
		new(IdGeneratorServiceImpl), server.WithServerBasicInfo(info), server.WithRegistry(r),
		server.WithServiceAddr(utils.NewNetAddr("tcp", fmt.Sprintf(":%v", config.IdGeneratorConfigObj.Port))),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
