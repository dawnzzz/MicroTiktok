namespace go message

include "base.thrift"

/*
* 消息记录
*/
struct message_chat_request {
    1: i64 to_user_id
    2: i64 pre_msg_time
}

struct message_chat_response {
    1: base.base_response base_resp(api.json="inline")
    2: list<base.Message> message_list
}

/*
* 发送消息
*/
struct message_send_request {
    1: i64 to_user_id
    2: i64 from_user_id
    4: string content
}

struct message_send_response {
    1: base.base_response base_resp
}

service MessageService {
    message_chat_response History(1: message_chat_request req)
    message_send_response SendMessage(1: message_send_request req)
}