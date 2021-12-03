package routes

import (
	"net/http"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tahsinature/future-proof-gin/pkg/controllers"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/middlewares"
)

var (
	userController   = new(controllers.UserController)
	secretController = new(controllers.SecretController)
)

func Setup() *gin.Engine {
	engine := gin.Default()

	binding.Validator = new(forms.DefaultValidator)

	engine.Use(middlewares.Cors)
	engine.Use(middlewares.RequestID)
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := engine.Group("/v1")
	secretRoutes := v1.Group("/secret")

	AddPingRoutes(engine.Group("/ping"))
	AddUserRoutes(v1.Group("/users"))
	AddSecretRoutes(secretRoutes)

	engine.LoadHTMLGlob("./pkg/public/html/*")
	engine.Static("/public", "./pkg/public")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"goVersion":             runtime.Version(),
			"ginBoilerplateVersion": "v0.03",
		})
	})

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return engine
}
