package router

import (
	_ "ChatSpace/docs"
	"ChatSpace/service"
	"ChatSpace/service/Message"
	"ChatSpace/service/ReletionShip"
	"ChatSpace/service/User"
	"ChatSpace/service/WebSocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	corsConf := cors.DefaultConfig()
	corsConf.AddAllowHeaders("Authorization")
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	//配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//配置路由规则

	r.GET("/ping", service.Ping)

	r.POST("/user/register", User.RegisterUser)
	r.POST("/user/login", User.Login)
	r.POST("/user/update_password", User.UpdatePassword)
	r.POST("/user/update_username", User.UpdateUsername)
	r.POST("/user/avatar", User.SetAvatar)
	r.GET("/user/my_info", User.MyInfo)

	r.POST("relationship/add_friend", ReletionShip.AddFriend)
	r.GET("relationship/friend_request_list", ReletionShip.FriendRequestList)
	r.GET("relationship/friend_list", ReletionShip.FriendList)
	r.POST("relationship/verify", ReletionShip.FriendVerify)

	manager := WebSocket.NewClientManager()
	go manager.Start()

	r.GET("/ws", func(c *gin.Context) {
		WebSocket.WebSocket(manager, c)
	}) //websocket入口

	r.POST("/message/files", Message.UploadFiles)
	r.POST("/message/videos", Message.UploadVideos)
	r.POST("/message/images", Message.UploadImages)
	r.POST("/message/audios", Message.UploadAudios)

	return r
}
