package api

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/dtu-admin/storage"
	"net/http"
)

type paramSearch struct {
	Offset    int    `form:"offset"`
	Length    int    `form:"length"`
	SortKey   string `form:"sortKey"`
	SortOrder string `form:"sortOrder"`
}

type paramId struct {
	Id int64 `uri:"id"`
}

type paramId2 struct {
	Id  int64 `uri:"id"`
	Id2 int64 `uri:"id2"`
}

func RegisterRoutes(app *gin.RouterGroup) {
	//注册 User类型
	gob.Register(&storage.User{})
	//启用session
	app.Use(sessions.Sessions("dtu-admin", memstore.NewStore([]byte("dtu-admin-secret"))))

	app.POST("/login", authLogin)

	//检查登录
	app.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		if user := session.Get("user"); user != nil {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	})

	app.DELETE("/logout", authLogout)
	app.POST("/password", authPassword)

	//TODO 转移至子目录，并使用中间件，检查session及权限
	app.GET("/channels")
	app.POST("/channels")
	app.POST("/channel")
	app.DELETE("/channel/:id")
	app.PUT("/channel/:id")
	app.GET("/channel/:id")
	app.GET("/channel/:id/start")
	app.GET("/channel/:id/stop")

	app.GET("/channel/:id/connections")
	app.POST("/channel/:id/connections")
	app.DELETE("/channel/:id/connection/:id2") //关闭连接
	app.GET("/channel/:id/connection/:id2/statistic")
	app.GET("/channel/:id/connection/:id2/pipe") //转Websocket透传

}

func responseOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func responseError(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":    false,
		"error": err,
	})
}

func nop(c *gin.Context) {
	c.String(http.StatusForbidden, "Unsupported")
}

func mustLogin(c *gin.Context) {
	//测试
	session := sessions.Default(c)
	if user := session.Get("user"); user != nil {
		c.Next()
	} else {
		//c.Redirect(http.StatusSeeOther, "/login")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}