package router

import (
	"net/http"

	"apiservice/handler/sd"
	"apiservice/handler/user"
	"apiservice/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// pprof router
	pprof.Register(g)

	// api for authentication functionlities
	g.POST("/login", user.Login)

	// 用户路由设置
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
	return g
}
