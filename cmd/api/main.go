package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wt993638658/simpletk/cmd/api/controller"
	"github.com/wt993638658/simpletk/cmd/api/rpc"
	"github.com/wt993638658/simpletk/pkg/jwt"
	"github.com/wt993638658/simpletk/pkg/ttviper"
)

var (
	Config = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	Jwt    *jwt.JWT
)

func Init(r *gin.Engine) {

	rpc.InitRPC(&Config)
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)                //
	apiRouter.GET("/user/", controller.UserInfo)            //
	apiRouter.POST("/user/register/", controller.Register)  //1
	apiRouter.POST("/user/login/", controller.Login)        //1
	apiRouter.POST("/publish/action/", controller.Publish)  //
	apiRouter.GET("/publish/list/", controller.PublishList) //

	//extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction) //
	apiRouter.GET("/favorite/list/", controller.FavoriteList)      //
	apiRouter.POST("/comment/action/", controller.CommentAction)   //
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)     //
	apiRouter.GET("/relation/follow/list/", controller.FollowList)     //
	apiRouter.GET("/relation/follower/list/", controller.FollowerList) //
	apiRouter.GET("/relation/friend/list/", controller.FollowerList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}

func main() {

	// 初始化 API 配置
	r := gin.Default()
	Init(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
