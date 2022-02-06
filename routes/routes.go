package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/controllers"
	md "github.com/vincentJunior1/test-kriya/middleware"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", md.JwtAuth("Admin"), controllers.GetUsers)
		users.GET("/:id", md.JwtAuth("Admin"), controllers.GetUser)
		users.POST("/", md.JwtAuth("Admin"), controllers.CreateUser)
		users.PATCH("/:id", md.JwtAuth("Admin"), controllers.UpdateUser)
		users.DELETE("/:id", md.JwtAuth("Admin"), controllers.DeleteUser)
	}
	login := r.Group("/login")
	{
		login.POST("/", controllers.Login)
	}

	return r
}
