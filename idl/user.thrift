namespace go user

include "base.thrift"

/*
* 注册
*/
struct user_register_request {
    1: string username,
    2: string password,
}

struct user_register_response {
    1: base.base_response base_resp,
    2: i64 user_id,
    3: string token,
}

/*
* 登录
*/
struct user_login_request {
    1: string username,
    2: string password,
}

struct user_login_response {
    1: base.base_response base_resp,
    2: i64 user_id,
    3: string token,
}

/*
* 获取用户信息
*/
struct get_user_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_user_response {
    1: base.base_response base_resp,
    2: base.User user,
}

/*
* 批量获取用户信息
*/
struct batch_get_user_request {
    1: i64 viewer_id,
    2: list<i64> owner_id_list,
}

struct batch_get_user_resonse {
    1: base.base_response base_resp,
    2: list<base.User> user_list,
}

service UserService {
    user_register_response Register(1: user_register_request req),                                      // 用户注册
    user_login_response Login(1: user_login_request req),                                               // 用户登录
    get_user_response GetUserInfo(1: get_user_request req),                                             // 获取用户信息
    batch_get_user_resonse BatchGetUserInfo(1: batch_get_user_request req),                             // 批量获取用户信息
}