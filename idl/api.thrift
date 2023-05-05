namespace go api

include "base.thrift"

/*
* 用户注册
*/
struct user_register_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // 用户名，最多32个字符
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // 密码，最多32个字符
}

struct user_register_reponse {
    1: base.base_response base_resp(api.json="inline")
    2: i64 user_id
    3: string token
}

/*
* 用户登录
*/
struct user_login_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // 用户名，最多32个字符
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // 密码，最多32个字符
}

struct user_login_response {
    1: base.base_response base_resp(api.json="inline")
    2: i64 user_id
    3: string token
}

/*
* 获取用户信息
*/
struct user_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct user_response {
    1: base.base_response base_resp(api.json="inline")
    2: base.User user
}

/*
* 获取视频
*/
struct feed_request {
    1: i64 latest_time(api.query="latest_time") // 获取视频列表的最晚时间，默认为当前时间
    2: string token(api.query="token")
}

struct feed_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Video> video_list  // 视频列表
    3: i64 next_time // video_list中的最早时间，作为下一次请求的latest_time
}

/*
* 发布视频
*/
struct publish_action_request {
    1: string token(api.form="token")
    2: string title(api.form="title", api.vd="len($)>0 && len($)<33")
}

struct publish_action_response {
    1: base.base_response base_resp(api.json="inline")
}

/*
* 获取发布视频列表
*/
struct publish_list_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct publish_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Video> video_list
}

/*
* 点赞/取消点赞
*/
struct favorite_action_request {
    1: string token(api.query="token")
    2: i64 video_id(api.query="video_id")
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-点赞, 2-取消
}

struct favorite_action_response {
    1: base.base_response base_resp(api.json="inline")
}

/*
* 用户点赞视频列表
*/
struct favorite_list_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct favorite_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Video> video_list
}

/*
* 评论/删除评论
*/
struct comment_action_request {
    1: string token(api.query="token")
    2: i64 video_id(api.query="video_id")
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-评论, 2-删除
    4: string comment_text(api.query="comment_text")
    5: i64 comment_id(api.query="comment_id")
}

struct comment_action_response {
    1: base.base_response base_resp(api.json="inline")
    2: base.Comment comment
}

/*
* 视频评论列表
*/
struct comment_list_request {
    1: string token(api.query="token")
    2: i64 video_id(api.query="video_id")
}

struct comment_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Comment> comment_list
}

/*
* 关注/取关
*/
struct relation_action_request {
    1: string token(api.query="token")
    2: i64 to_user_id(api.query="to_user_id")
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-关注, 2-去管
}

struct relation_action_response {
    1: base.base_response base_resp(api.json="inline")
}

/*
* 关注列表
*/
struct relation_follow_list_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct relation_follow_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.User> user_list
}

/*
* 粉丝列表
*/
struct relation_follower_list_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct relation_follower_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.User> user_list
}

/*
* 朋友列表
*/
struct relation_friend_list_request {
    1: i64 user_id(api.query="user_id")
    2: string token(api.query="token")
}

struct relation_friend_list_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.FriendUser> user_list
}

/*
* 消息记录
*/
struct message_chat_request {
    1: string token(api.query="token")
    2: i64 to_user_id(api.query="to_user_id")
    3: i64 pre_msg_time(api.query="pre_msg_time")// The time of time of last latest message
}

struct message_chat_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Message> message_list
}

/*
* 发送消息
*/
struct message_action_request {
    1: string token(api.query="token")
    2: i64 to_user_id(api.query="to_user_id")
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-发送消息 2-撤回消息
    4: string content(api.query="content", api.vd="len($)>0 && len($)<255")
}

struct message_action_response {
    1: base.base_response base_resp(api.json="inline")
}

// API网关
service APIGatewayService {
    // 用户API
    user_register_reponse Register(1: user_register_request req)(api.post="/douyin/user/register/");    // 用户注册
    user_login_response Login(1: user_login_request req)(api.post="/douyin/user/login/");               // 用户登录
    user_response GetUserInfo(1: user_request req)(api.get="/douyin/user/");                            // 查询用户信息

    // 视频API
    feed_response Feed (1: feed_request req)(api.get="/douyin/feed/");                                          // 推送视频
    publish_action_response PublishVideo (1: publish_action_request req)(api.post="/douyin/publish/action/");   // 发布视频
    publish_list_response VideoList (1: publish_list_request req)(api.get="/douyin/publish/list/");             // 获取用户视频列表

    // 点赞API
    favorite_action_response Favorite(1: favorite_action_request req)(api.post="/douyin/favorite/action/"); // 点赞/取消视频
    favorite_list_response FavoriteList(1: favorite_list_request req)(api.get="/douyin/favorite/list/");    // 获取点赞列表

    // 评论API
    comment_action_response Comment(1: comment_action_request req)(api.post="/douyin/comment/action/");     // 评论视频
    comment_list_response CommentList(1: comment_list_request req)(api.get="/douyin/comment/list/");        // 获取视频评论列表

    // 关系API
    relation_action_response Action(1: relation_action_request req)(api.post="/douyin/relation/action/");                           // 关注/取关
    relation_follow_list_response FollowingList(1: relation_follow_list_request req)(api.get="/douyin/relation/follow/list/");      // 获取关注列表
    relation_follower_list_response FollowerList(1: relation_follower_list_request req)(api.get="/douyin/relation/follower/list/"); // 获取粉丝列表
    relation_friend_list_response FriendList(1: relation_friend_list_request req)(api.get="/douyin/relation/friend/list/");         // 获取朋友列表（互关）

    // 消息API
    message_chat_response ChatHistory(1: message_chat_request req)(api.get="/douyin/message/chat/");        // 获取聊天记录
    message_action_response SentMessage(1: message_action_request req)(api.post="/douyin/message/action/"); // 发送消息
}