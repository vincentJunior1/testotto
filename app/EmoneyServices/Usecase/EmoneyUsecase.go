package usecase

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	emoneyRepo "github.com/vincentJunior1/test-kriya/app/EmoneyServices/Repository"
	"github.com/vincentJunior1/test-kriya/app/EmoneyServices/dbEntity"
	"github.com/vincentJunior1/test-kriya/app/EmoneyServices/httpEntity"
	"github.com/vincentJunior1/test-kriya/helper"
	"golang.org/x/crypto/bcrypt"
)

type EmoneyServiceUsecase struct {
	ctx        context.Context
	EmoneyRepo emoneyRepo.EmoneyRepository
}

func NewEmoneyUsecase(
	ctx context.Context,
	EmoneyRepo emoneyRepo.EmoneyRepository,
) *EmoneyServiceUsecase {
	return &EmoneyServiceUsecase{
		ctx:        ctx,
		EmoneyRepo: EmoneyRepo,
	}
}

func (e *EmoneyServiceUsecase) Login(c *gin.Context) string {
	return "dapet nih"
}

func (e *EmoneyServiceUsecase) Register(c *gin.Context) *helper.Response {
	payload := &httpEntity.UserRequest{}

	errMessage := ""

	if err := c.BindJSON(&payload); err != nil {
		return &helper.Response{
			Code:    http.StatusBadRequest,
			Message: errMessage,
			Status:  false,
			Data:    nil,
		}
	}

	if payload.Name == "" {
		errMessage = "Name Can't be empty"
	} else if payload.Email == "" {
		errMessage = "Email Can't be empty"
	} else if payload.Password == "" {
		errMessage = "Password Can't be empty"
	}

	if errMessage != "" {
		return &helper.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: errMessage,
			Status:  false,
			Data:    nil,
		}
	}

	user := dbEntity.User{}
	if err := e.EmoneyRepo.GetUserByParams(map[string]interface{}{
		"email": payload.Email,
	}, &user); err != nil {
		return &helper.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error create user",
			Status:  false,
			Data:    nil,
		}
	}

	if user.ID != 0 {
		return &helper.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Email already created",
			Status:  false,
			Data:    nil,
		}
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 6)

	newUser := &dbEntity.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
	}

	if err := e.EmoneyRepo.CreateUser(newUser); err != nil {
		return &helper.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error create user",
			Status:  false,
			Data:    nil,
		}
	}

	return &helper.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Status:  false,
		Data:    newUser,
	}

}
