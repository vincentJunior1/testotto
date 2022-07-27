package EmoneyHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/app/EmoneyServices"
	// "github.com/vincentJunior1/test-kriya/middleware"
)

type EmoneyHandler struct {
	emoneyUsecase EmoneyServices.EmoneyServiceUsecase
}

func EmoneyRoutes(emoneyUsecases EmoneyServices.EmoneyServiceUsecase, r *gin.Engine) {
	emoneyUC := EmoneyHandler{
		emoneyUsecase: emoneyUsecases,
	}

	r.GET("/login", emoneyUC.LoginHanlder)
	r.POST("/register", emoneyUC.Register)

}

func (e *EmoneyHandler) LoginHanlder(c *gin.Context) {
	login := e.emoneyUsecase.Login(c)

	c.JSON(http.StatusOK, login)
}

func (e *EmoneyHandler) Register(c *gin.Context) {
	register := e.emoneyUsecase.Register(c)
	c.JSON(register.Code, register)
}
