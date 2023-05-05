namespace go favorite

include "base.thrift"

/*
* 点赞
*/
struct favorite_request {
    1: i64 user_id,
    2: i64 video_id,
}

struct favorite_response {
    1: base.base_response base_resp,
}

/*
* 取消点赞
*/
struct favorite_cancel_request {
    1: i64 user_id,
    2: i64 video_id,
}

struct favorite_cancel_response {
    1: base.base_response base_resp,
}

/*
* 获取点赞视频列表
*/
struct get_favorite_video_list_request {
    1: i64 user_id,
}

struct get_favorite_video_list_response {
    1: base.base_response base_resp,
    2: list<base.Video> video_list,
}

// 点赞的RPC接口
service FavoriteService {
    favorite_response Favorite(1: favorite_request req),                                            // 点赞
    favorite_cancel_response FavoriteCancel(1: favorite_cancel_request req),                        // 取消点赞
    get_favorite_video_list_response GetFavoriteVideoList(1: get_favorite_video_list_request req),  // 获取点赞列表
}