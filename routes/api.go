package routes

import (
	"go-api/src/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	}))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("assets"))))
	route.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	api := route.Group("/api")
	{
		api.GET("/user", controller.User)
		api.POST("/user", controller.User)

		subjectGroup := api.Group("/subject")
		{
			subjectGroup.GET("/", controller.Subject)
			subjectGroup.POST("/", controller.Subject)
			subjectGroup.GET("/:id", controller.SubjectByID)
			subjectGroup.PATCH("/:id", controller.SubjectByID)
		}
	}
	return route
}
