package controllers

import (
	"net/http"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/httpEntity"
	"github.com/vincentJunior1/test-kriya/models"
	"golang.org/x/crypto/bcrypt"
)

type JwtStruct struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type JwtClaims struct {
	NewData *models.User
	jwt.StandardClaims
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	payload := &httpEntity.UserRequest{}
	errMessage := ""

	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if payload.Name == "" {
		errMessage = "Name can't be empty"
	} else if payload.UserName == "" {
		errMessage = "User name can't be empty"
	} else if payload.Password == "" {
		errMessage = "Password can't be empty"
	}

	if errMessage != "" {
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: errMessage,
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	findUserName := &models.User{}
	if err := models.GetUserByParams(map[string]interface{}{
		"user_name": payload.UserName,
	}, findUserName); err != nil {
		log.Fatal(err.Error())
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Failed Find Email",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	if findUserName.ID != 0 {
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Email already created",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	user := &models.User{
		UserName:  payload.UserName,
		Name:      payload.Name,
		Password:  string(newPassword),
		CreatedBy: 1,
		UpdatedBy: 1,
	}

	err := models.CreateUser(user)

	if err != nil {
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error create data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &httpEntity.Response{
			Code:    http.StatusOK,
			Message: "Success",
			Status:  false,
			Data:    user,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
}

// login user
func Login(c *gin.Context) {
	payload := &httpEntity.LoginUser{}

	if err := c.BindJSON(payload); err != nil {
		response := &httpEntity.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	user := &models.User{}
	if err := models.GetUserByParams(map[string]interface{}{
		"user_name": payload.UserName,
	}, user); err != nil {
		log.Fatal(err.Error())
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Failed Find user name",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if user.ID == 0 {
		log.Fatal("error: user not found")
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "user name not found",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		log.Fatal("error: user not found")
		response := &httpEntity.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "wrong password",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	claims := JwtClaims{
		NewData:        user,
		StandardClaims: jwt.StandardClaims{},
	}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	response := &httpEntity.Response{
		Code:    http.StatusUnprocessableEntity,
		Message: "wrong password",
		Status:  false,
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	}
	c.JSON(http.StatusOK, response)
	c.Abort()
}
