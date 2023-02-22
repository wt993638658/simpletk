namespace go relation

include "user.thrift"

//struct User {
//    1:i64 id // 用户id
//    2:string name // 用户名称
//    3:optional i64 follow_count // 关注总数
//    4:optional i64 follower_count // 粉丝总数
//    5:bool is_follow // true-已关注，false-未关注
//}

struct douyin_relation_action_request {
    1:i64 user_id  // 用户id
    2:string token  // 用户鉴权token
    3:i64 to_user_id  // 对方用户id
    4:i32 action_type  // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1:i32 status_code   // 状态码，0-成功，其他值-失败
    2:optional string status_msg  // 返回状态描述
}

struct douyin_relation_follow_list_request {
    1:i64 user_id  // 用户id
    2:string token  // 用户鉴权token
}

struct douyin_relation_follow_list_response {
    1:i32 status_code // 状态码，0-成功，其他值-失败
    2:optional string status_msg // 返回状态描述
    3:list<user.User> user_list // 用户信息列表
}

struct douyin_relation_follower_list_request {
    1:i64 user_id  // 用户id
    2:string token  // 用户鉴权token
}

struct douyin_relation_follower_list_response {
    1:i32 status_code  // 状态码，0-成功，其他值-失败
    2:optional string status_msg // 返回状态描述
    3:list<user.User> user_list // 用户列表
}

service RelationSrv{
    douyin_relation_action_response RelationAction(1:douyin_relation_action_request req)
    douyin_relation_follow_list_response RelationFollowList(1:douyin_relation_follow_list_request req)//获取已关注用户的列表
    douyin_relation_follower_list_response RelationFollowerList(1:douyin_relation_follower_list_request req)//获取粉丝用户列表
}