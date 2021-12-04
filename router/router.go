package router

import (
	"Gobang-2004A/config"
	"Gobang-2004A/controller"
	"Gobang-2004A/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	addr := ":" + strconv.Itoa(config.Config.Get("server.port").(int))
	router := gin.Default()
	//使用中间件
	router.Use(middleware.Cors())
	//加载静态资源
	router.LoadHTMLFiles("./static/index.html")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	//请求
	router.GET("/", controller.Index)
	router.GET("/createroom", controller.CreateRoom)
	router.GET("/enterroom", controller.EnterRoom)

	router.Run(addr)
}
