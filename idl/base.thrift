namespace go base

struct base_response {
    1: i32 status_code  // 0 is success, other is failed
    2: string status_msg
}

// 用户
struct User {
    1: i64 id                                               // 用户id
    2: string name                                          // 用户名
    3: SocialInfo socialInfo(api.json="inline")             // 关注数量、粉丝数量、是否被关注
    4: string avatar,                                       // 用户头像
    5: string background_image,                             // 用户背景图
    6: string signature,                                    // 用户签名
    7: UserInteractInfo userInteracInfo(api.json="inline")  // 用户被赞、发布作品数量、点赞数量
}

// 关注数量、粉丝数量、是否被关注
struct SocialInfo {
    1: i64 follow_count
    2: i64 follower_count
    3: bool is_follow
}

// 用户被赞、点赞数量、发布作品数量
struct UserInteractInfo {
    1: i64 total_favorited,     // 被赞数量
    2: i64 work_count,          // 发布作品数量
    3: i64 favorite_count,      // 点赞数量
}

// 视频
struct Video {
    1: i64 id                                                   // 视频id
    2: User author                                              // 视频作者
    3: string play_url                                          // 视频url
    4: string cover_url                                         // 封面
    5: VideoInteractInfo videoInteractInfo(api.json="inline")   // 视频的点赞数量、评论数量、是否点赞
    6: string title                                             // 视频标题
}

// 视频的点赞数量、评论数量、是否点赞
struct VideoInteractInfo {
    1: i64 favorite_count
    2: i64 comment_count
    3: bool is_favorite
}

// 评论
struct Comment {
    1: i64 id               // 评论id
    2: User user            // 评论者信息
    3: string content       // 内容
    4: string create_date   // 评论发布时间
}

struct FriendUser {
    1: User user(api.json="inline") // 用户信息
    2: string message               // 最后一次消息
    3: i64 msgType                  // 消息类型：0表示接收消息，1表示发送消息
}

// 消息
struct Message {
    1: i64 id           // 消息ID
    2: i64 to_user_id   // 接收者id
    3: i64 from_user_id // 发送者id
    4: string content   // 消息内容
    5: i64 create_time  // 消息发送时间
}