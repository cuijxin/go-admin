package router

import (
	"go-admin/middleware"
	"log"

	"github.com/gin-gonic/gin"

	"go-admin/handler"
	"go-admin/handler/sd"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())
	r.Use(middleware.DemoEnv())

	r.GET("/info", handler.Ping)

	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()

	if err != nil {
		log.Fatalln("JWT Error", err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/routes", Dashboard)

	log.Println("路由加载成功！")
	return r
}

func Dashboard(c *gin.Context) {

	var user = make(map[string]interface{})
	user["login_name"] = "admin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data

	c.JSON(200, r)
}
