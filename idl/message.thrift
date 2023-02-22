namespace go message


struct douyin_message_action_request {
    1:string token  // 用户鉴权token
    2:i64 to_user_id  // 对方用户id
    3:i32 action_type  // 1-关注，2-取消关注
    4:string content
}

struct douyin_message_action_response {
    1:i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
}
struct douyin_message_chat_request {
    1:string token  // 用户鉴权token
    2:i64 to_user_id  // 对方用户id
}

struct douyin_message_char_response {
    1:i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
    3:list<Message> message_list
}
struct Message{
    1:i64 id
    2:i64 to_user_id
    3:i64 from_user_id
    4:string content
    5:i64 create_time
}

service MessageSrv{
    douyin_message_action_response MessageAction(1:douyin_message_action_request req)
    douyin_message_char_response MessageChat(1:douyin_message_chat_request req)
}