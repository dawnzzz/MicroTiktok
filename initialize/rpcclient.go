package initialize

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/dawnzzz/MicroTiktok/global"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/authentication/authenticationservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/comment/commentservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/favorite/favoriteservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/id_generator/idgeneratorservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/interact/interactservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/message/messageservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/relation/relationservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/society/societyservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/user/userservice"
	"github.com/dawnzzz/MicroTiktok/kitex_gen/video/videoservice"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func rpcClientOptions() []client.Option {
	var options []client.Option
	// 添加服务发现
	options = append(options, client.WithResolver(newEtcdResolver()))
	// 添加多路复用
	options = append(options, client.WithMuxConnection(1))

	return options
}

func newEtcdResolver() discovery.Resolver {
	etcdResolver, err := etcd.NewEtcdResolver(global.EtcdEndpoints)
	if err != nil {
		panic(err)
	}

	return etcdResolver
}

func InitRpcUserClient() {
	var err error
	global.RpcUserClient, err = userservice.NewClient(global.RpcUserServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcVideoClient() {
	var err error
	global.RpcVideoClient, err = videoservice.NewClient(global.RpcVideoServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcFavoriteClient() {
	var err error
	global.RpcFavoriteClient, err = favoriteservice.NewClient(global.RpcFavoriteServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcCommentClient() {
	var err error
	global.RpcCommentClient, err = commentservice.NewClient(global.RpcCommentServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcRelationClient() {
	var err error
	global.RpcRelationClient, err = relationservice.NewClient(global.RpcRelationServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcMessageClient() {
	var err error
	global.RpcMessageClient, err = messageservice.NewClient(global.RpcMessageServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcAuthenticationClient() {
	var err error
	global.RpcAuthenticationClient, err = authenticationservice.NewClient(global.RpcAuthenticationServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcInteractClient() {
	var err error
	global.RpcInteractClient, err = interactservice.NewClient(global.RpcInteractServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcSocietyClient() {
	var err error
	global.RpcSocietyClient, err = societyservice.NewClient(global.RpcSocietyServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}

func InitRpcIdGeneratorClient() {
	var err error
	global.RpcIdGeneratorClient, err = idgeneratorservice.NewClient(global.RpcIdGeneratorServiceName, rpcClientOptions()...)
	if err != nil {
		panic(err)
	}
}
