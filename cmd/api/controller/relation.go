package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wt993638658/simpletk/cmd/api/rpc"
	"github.com/wt993638658/simpletk/dal/pack"
	"github.com/wt993638658/simpletk/kitex_gen/relation"
	"github.com/wt993638658/simpletk/pkg/errno"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	var paramVar RelationActionParam
	token := c.Query("token")
	to_user_id := c.Query("to_user_id")
	action_type := c.Query("action_type")

	tid, err := strconv.Atoi(to_user_id)
	if err != nil {
		SendResponse(c, pack.BuildRelationActionResp(errno.ErrBind))
		return
	}

	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildRelationActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.ToUserId = int64(tid)
	paramVar.ActionType = int32(act)

	rpcReq := relation.DouyinRelationActionRequest{
		ToUserId:   paramVar.ToUserId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	}
	ctx := context.Background()
	resp, err := rpc.RelationAction(ctx, &rpcReq)
	if err != nil {
		SendResponse(c, pack.BuildRelationActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	var paramVar UserParam
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFollowingListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(uid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFollowingListResp(errno.ErrBind))
		return
	}
	ctx := context.Background()
	resp, err := rpc.RelationFollowList(ctx, &relation.DouyinRelationFollowListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFollowingListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	var paramVar UserParam
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFollowerListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(uid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFollowerListResp(errno.ErrBind))
		return
	}
	ctx := context.Background()
	resp, err := rpc.RelationFollowerList(ctx, &relation.DouyinRelationFollowerListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFollowerListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
