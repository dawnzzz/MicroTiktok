namespace go comment

include "base.thrift"

/*
* 评论
*/
struct comment_request {
    1: i64 user_id,
    2: i64 video_id,
    3: string comment_text,
}

struct comment_response {
    1: base.base_response base_resp,
    2: base.Comment comment,
}

/*
* 删除评论
*/
struct comment_delete_request {
    1: i64 user_id,
    2: i64 video_id,
    3: i64 comment_id,
}

struct comment_delete_response {
    1: base.base_response base_resp,
}

/*
* 获取评论列表
*/
struct get_comment_list_request {
    1: i64 video_id,
}

struct get_comment_list_response {
    1: base.base_response base_resp,
    2: list<base.Comment> comment_list,
}

service CommentService {
    comment_response Comment(1: comment_request req),                           // 评论
    comment_delete_response CommentDelete(1: comment_delete_request req),       // 删除评论
    get_comment_list_response GetCommentList(1: get_comment_list_request req),  // 获取视频评论列表
}