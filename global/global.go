package global

import (
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
)

// RPC Client
var (
	RpcUserClient           userservice.Client
	RpcVideoClient          videoservice.Client
	RpcFavoriteClient       favoriteservice.Client
	RpcCommentClient        commentservice.Client
	RpcRelationClient       relationservice.Client
	RpcMessageClient        messageservice.Client
	RpcAuthenticationClient authenticationservice.Client
	RpcInteractClient       interactservice.Client
	RpcSocietyClient        societyservice.Client
	RpcIdGeneratorClient    idgeneratorservice.Client
)

// RPC Service name
const (
	RpcUserServiceName           = "user"
	RpcVideoServiceName          = "video"
	RpcFavoriteServiceName       = "favorite"
	RpcCommentServiceName        = "comment"
	RpcRelationServiceName       = "relation"
	RpcMessageServiceName        = "message"
	RpcAuthenticationServiceName = "authentication"
	RpcInteractServiceName       = "interact"
	RpcSocietyServiceName        = "society"
	RpcIdGeneratorServiceName    = "id_generator"
)

var (
	EtcdEndpoints []string
)

const (
	HzViewerIDKey = "viewer_id"
)
