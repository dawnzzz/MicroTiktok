namespace go interact

include "base.thrift"

/*
* 获取视频的点赞数量、是否点赞、评论数量
*/
struct get_video_interact_info_request{
    1: i64 video_id, // Video Id
    2: i64 viewer_id, // viewer_id,
}

struct get_video_interact_info_response {
    1: base.base_response base_resp,
    2: base.VideoInteractInfo interact_info,
}

/*
* 批量获取视频的点赞数量、是否点赞、评论数量
*/
struct batch_get_video_interact_info_request{
    1: list<i64> video_id_list, // Video Id list.
    2: i64 viewer_id, // viewer_id,
}

struct batch_get_video_interact_info_response {
    1: base.base_response base_resp,
    2: list<base.VideoInteractInfo> interact_info_list,
}

/*
* 获取用户的被赞数量、点赞数量、作品数量
*/
struct get_user_interact_info_request{
    1: i64 user_id,
}

struct get_user_interact_info_response {
    1: base.base_response base_resp,
    2: base.UserInteractInfo interact_info,
}

/*
* 批量获取用户的被赞数量、点赞数量、作品数量
*/
struct batch_get_user_interact_info_request{
    1: list<i64> user_id_list,
}

struct batch_get_user_interact_info_response {
    1: base.base_response base_resp,
    2: list<base.UserInteractInfo> interact_info_list,
}

service InteractService {
    // 获取视频的交互信息（点赞数量、评论数量、是否点赞）
    get_video_interact_info_response GetVideoInteractInfo (1:get_video_interact_info_request req),
    batch_get_video_interact_info_response BatchGetVideoInteractInfo (1: batch_get_video_interact_info_request req),

    // 获取用户交互信息（获取用户的被赞数量、点赞数量、作品数量）
    get_user_interact_info_response GetUserInteractInfo (1: get_user_interact_info_request req),
    batch_get_user_interact_info_response BatchGetUserInteractInfo (1: batch_get_user_interact_info_request req),
}