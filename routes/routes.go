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
		users.POST("/", md.JwtAuth("Admin"), controllers.CreateUser)
	}
	login := r.Group("/login")
	{
		login.POST("/", controllers.Login)
	}
	merchant := r.Group("/merchant")
	{
		merchant.GET("/omzet/:merchant_id", md.JwtAuth("Admin"), controllers.GetOmzet)
	}

	return r
}
