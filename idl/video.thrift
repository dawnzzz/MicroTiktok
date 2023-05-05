namespace go video

include "base.thrift"

/*
* 点赞视频
*/
struct feed_request {
    1: i64 latest_time,
    2: i64 viewer_id,
}

struct feed_response {
    1: base.base_response base_resp,
    2: list<base.Video> video_list,
    3: i64 next_time,
}

/*
* 发布视频
*/
struct publish_request {
    1: i64 user_id,
    2: string play_url,
    3: string cover_url,
    4: string title,
}

struct publish_response {
    1: base.base_response base_resp,
}

/*
* 获取发布视频列表
*/
struct get_published_list_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_published_list_response {
    1: base.base_response base_resp,
    2: list<base.Video> video_list,
}

/*
* 获取点赞视频列表
*/
struct get_favorite_list_request {
    1: i64 viewer_id,
    2: i64 owner_id,
}

struct get_favorite_list_response {
    1: base.base_response base_resp,
    2: list<base.Video> video_list,
}

service VideoService {
    feed_response Feed(1: feed_request req),                                                                    // 获取视频
    publish_response PublishVideo(1: publish_request req),                                                      // 发布视频
    get_published_list_response GetPublishedVideoList(1: get_published_list_request req),                       // 获取发布列表
    get_favorite_list_response GetFavoriteVideoList(1: get_favorite_list_request req),                          // 获取点赞视频列表
}