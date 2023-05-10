package initialize

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
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

func rpcClientOptions(serviceName string) []client.Option {
	var options []client.Option
	// 添加服务发现
	options = append(options, client.WithResolver(newEtcdResolver()))
	options = append(options, client.WithLoadBalancer(loadbalance.NewWeightedBalancer()))
	// 添加多路复用
	options = append(options, client.WithMuxConnection(2))
	// 设置TTHeader传输协议
	options = append(options, client.WithTransportProtocol(transport.TTHeader))
	options = append(options, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	//options = append(options, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}))

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
	global.RpcUserClient, err = userservice.NewClient(global.RpcUserServiceName, rpcClientOptions(global.RpcUserServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcVideoClient() {
	var err error
	global.RpcVideoClient, err = videoservice.NewClient(global.RpcVideoServiceName, rpcClientOptions(global.RpcVideoServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcFavoriteClient() {
	var err error
	global.RpcFavoriteClient, err = favoriteservice.NewClient(global.RpcFavoriteServiceName, rpcClientOptions(global.RpcFavoriteServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcCommentClient() {
	var err error
	global.RpcCommentClient, err = commentservice.NewClient(global.RpcCommentServiceName, rpcClientOptions(global.RpcCommentServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcRelationClient() {
	var err error
	global.RpcRelationClient, err = relationservice.NewClient(global.RpcRelationServiceName, rpcClientOptions(global.RpcRelationServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcMessageClient() {
	var err error
	global.RpcMessageClient, err = messageservice.NewClient(global.RpcMessageServiceName, rpcClientOptions(global.RpcMessageServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcAuthenticationClient() {
	var err error
	global.RpcAuthenticationClient, err = authenticationservice.NewClient(global.RpcAuthenticationServiceName, rpcClientOptions(global.RpcAuthenticationServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcInteractClient() {
	var err error
	global.RpcInteractClient, err = interactservice.NewClient(global.RpcInteractServiceName, rpcClientOptions(global.RpcInteractServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcSocietyClient() {
	var err error
	global.RpcSocietyClient, err = societyservice.NewClient(global.RpcSocietyServiceName, rpcClientOptions(global.RpcSocietyServiceName)...)
	if err != nil {
		panic(err)
	}
}

func InitRpcIdGeneratorClient() {
	var err error
	global.RpcIdGeneratorClient, err = idgeneratorservice.NewClient(global.RpcIdGeneratorServiceName, rpcClientOptions(global.RpcIdGeneratorServiceName)...)
	if err != nil {
		panic(err)
	}
}
