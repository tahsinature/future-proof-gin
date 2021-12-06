package routes

import (
	"net/http"
	"path"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/controllers"
	"github.com/tahsinature/future-proof-gin/pkg/exception"
	"github.com/tahsinature/future-proof-gin/pkg/middlewares"
)

var (
	userController   = new(controllers.User)
	secretController = new(controllers.Secret)
)

func Setup() *gin.Engine {
	engine := gin.New()

	engine.Use(middlewares.Cors)
	engine.Use(middlewares.RequestID)
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := engine.Group("/v1")

	new(Ping).setup(engine.Group("/ping"))
	new(User).setup(v1.Group("/users"))
	new(Secret).setup(v1.Group("/secret"))

	engine.LoadHTMLGlob(path.Join(config.Other.RootPath, "pkg", "public", "html/*"))
	engine.Static("/public", path.Join(config.Other.RootPath, "pkg", "public"))

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"goVersion":             runtime.Version(),
			"ginBoilerplateVersion": "v0.03",
		})
	})

	engine.NoRoute(func(c *gin.Context) {
		controllers.Response.FromError(c, exception.Response{HTTPCode: http.StatusNotFound, Message: "Seems like you are lost.", Flag: exception.Flags.Get("ROUTE_NOT_FOUNT")})
	})

	return engine
}
