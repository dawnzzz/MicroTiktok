namespace go society

include "base.thrift"

/*
* 获取关注数量、粉丝数量、是否被关注
*/
struct get_social_info_request{
    1: i64 viewer_id
    2: i64 owner_id
}

struct get_social_info_response{
    1: base.base_response base_resp
    2: base.SocialInfo social_info
}

/*
* 批量获取关注数量、粉丝数量、是否被关注
*/
struct batch_get_social_info_request{
    1: i64 viewer_id
    2: list<i64> owner_id_list
}

struct batch_get_social_info_response{
    1: base.base_response base_resp
    2: list<base.SocialInfo> social_info_list
}

service SocietyService {
    get_social_info_response GetSocialInfo(1:get_social_info_request req),
    batch_get_social_info_response BatchGetSocialInfo(1:batch_get_social_info_request req),
}