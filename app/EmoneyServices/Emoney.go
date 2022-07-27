package EmoneyServices

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/app/EmoneyServices/dbEntity"
	"github.com/vincentJunior1/test-kriya/helper"
)

type EmoneyServiceUsecase interface {
	Login(c *gin.Context) string
	Register(c *gin.Context) *helper.Response
}

type EmoneyServiceRepo interface {
	GetUserByParams(params map[string]interface{}, user *dbEntity.User) error
}
