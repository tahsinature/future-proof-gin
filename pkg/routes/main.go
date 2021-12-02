package routes

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/controllers"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/middlewares"
)

var (
	userController   = new(controllers.UserController)
	secretController = new(controllers.SecretController)
)

func Setup() *gin.Engine {
	fmt.Println(config.DB)
	routes := gin.Default()
	binding.Validator = new(forms.DefaultValidator)

	routes.Use(middlewares.Cors)
	routes.Use(middlewares.RequestID)
	routes.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := routes.Group("/v1")
	secretRoutes := v1.Group("/secret")

	AddPingRoutes(routes.Group("/ping"))
	AddUserRoutes(v1.Group("/users"))
	AddSecretRoutes(secretRoutes)

	routes.LoadHTMLGlob("./pkg/public/html/*")
	routes.Static("/public", "./pkg/public")

	routes.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"goVersion":             runtime.Version(),
			"ginBoilerplateVersion": "v0.03",
		})
	})

	routes.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return routes
}
