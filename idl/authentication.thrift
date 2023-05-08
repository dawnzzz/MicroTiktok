namespace go authentication

include "base.thrift"

/*
* 生成token
*/
struct generate_token_request {
    1: i64 user_id,
}

struct generate_token_response {
    1: base.base_response base_resp,
    2: string token,
}

/*
* 验证token是否合法
*/
struct check_token_request {
    1: string token
}

struct check_token_response {
    1: base.base_response base_resp,
    2: i64 user_id,
}

service AuthenticationService {
    generate_token_response GenerateToken(1: generate_token_request req)    // 获取token
    check_token_response CheckToken(1: check_token_request req)             // 验证token是否合法
}