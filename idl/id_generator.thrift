namespace go id_generator

include "base.thrift"

/*
* 生成user id
*/
struct get_user_id_request {
}

struct get_user_id_response {
     1: base.base_response base_resp,
     2: i64 user_id,
}

/*
* 生成video id
*/
struct get_video_id_request {
}

struct get_video_id_response {
     1: base.base_response base_resp,
     2: i64 video_id,
}

/*
* 生成comment id
*/
struct get_comment_id_request {
}

struct get_comment_id_response {
     1: base.base_response base_resp,
     2: i64 comment_id,
}

service IdGeneratorService {
    get_user_id_response GetUserID(1: get_user_id_request req),
    get_video_id_request GetVideoID(1: get_video_id_request req),
    get_comment_id_request GetCommentID(1: get_comment_id_request req),
}