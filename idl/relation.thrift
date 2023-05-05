namespace go relation

include "base.thrift"

/*
* 关注
*/
struct relation_request {
    1: i64 user_id,
    2: i64 to_user_id,
}

struct relation_response {
    1: base.base_response base_resp,
}

/*
* 取消关注
*/
struct relation_cancel_request {
    1: i64 user_id,
    2: i64 to_user_id,
}

struct relation_cancel_response {
    1: base.base_response base_resp,
}

/*
* 获取关注列表
*/
struct get_relation_follow_list_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_relation_follow_list_response {
    1: base.base_response base_resp,
    2: list<base.User> user_list,
}

/*
* 获取粉丝列表
*/
struct get_relation_follower_list_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_relation_follower_list_response {
    1: base.base_response base_resp,
    2: list<base.User> user_list,
}

/*
* 获取朋友列表
*/
struct get_relation_friend_list_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_relation_friend_list_response {
    1: base.base_response base_resp,
    2: list<base.FriendUser> user_list,
}


service SocialityService {
    relation_response Relation(1: relation_request req),                        // 关注
    relation_cancel_response RelationCancel(1: relation_cancel_request req),    // 取消关注

    get_relation_follow_list_response GetFollowList(1: get_relation_follow_list_request req),           // 获取关注列表
    get_relation_follower_list_response GetFollowerList(1: get_relation_follower_list_request req),     // 获取粉丝列表
    get_relation_friend_list_response GetFriendList(1: get_relation_friend_list_request req),           // 获取朋友列表
}