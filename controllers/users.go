package controllers

import (
	"net/http"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/httpEntity"
	"github.com/vincentJunior1/test-kriya/models"
	"github.com/vincentJunior1/test-kriya/utils"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Status     bool        `json:"status"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type JwtStruct struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type JwtClaims struct {
	NewData *models.User
	jwt.StandardClaims
}

// GetUsers gets all existing users.
func GetUsers(c *gin.Context) {
	var users []models.User
	pagination := utils.GeneratePaginationFromRequest(c)

	err := models.GetUsers(&users, &pagination)

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error get data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &Response{
			Code:       http.StatusOK,
			Message:    "Success",
			Status:     true,
			Data:       users,
			Pagination: pagination,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
	}
}

// GetUser finds a single user by ID.
func GetUser(c *gin.Context) {
	var user models.User

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error get data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &Response{
			Code:    http.StatusOK,
			Message: "Success",
			Status:  true,
			Data:    user,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
	}
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	payload := &httpEntity.UserRequest{}
	errMessage := ""

	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if payload.Email == "" {
		errMessage = "Email can't be empty"
	} else if payload.UserName == "" {
		errMessage = "User name can't be empty"
	} else if payload.RoleID != 1 || payload.RoleID != 0 {
		errMessage = "Please input the role id"
	} else if payload.Password == "" {
		errMessage = "Password can't be empty"
	} else if payload.Status != 1 || payload.Status != 0 {
		errMessage = "Status should be active and inactive"
	}

	if errMessage != "" {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: errMessage,
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	findEmail := &models.User{}
	if err := models.GetUserByParams(map[string]interface{}{
		"email": payload.Email,
	}, findEmail); err != nil {
		log.Fatal(err.Error())
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Failed Find Email",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	if findEmail.ID != 0 {
		response := &Response{
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
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: string(newPassword),
		RoleID:   payload.RoleID,
		Status:   payload.Status,
	}

	err := models.CreateUser(user)

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error create data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &Response{
			Code:    http.StatusOK,
			Message: "Email already created",
			Status:  false,
			Data:    user,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
}

// UpdateUser updates a new user by ID.
func UpdateUser(c *gin.Context) {
	var user models.User
	payload := &httpEntity.UserRequest{}
	errMessage := ""

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error get user",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if user.ID == 0 {
		response := &Response{
			Code:    http.StatusBadRequest,
			Message: "User not found",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	c.BindJSON(&payload)

	if errMessage != "" {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: errMessage,
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if payload.Email == "" {
		errMessage = "Email can't be empty"
	} else if payload.UserName == "" {
		errMessage = "User name can't be empty"
	} else if payload.RoleID != 1 || payload.RoleID != 0 {
		errMessage = "Please input the role id"
	} else if payload.Password == "" {
		errMessage = "Password can't be empty"
	} else if payload.Status != 1 || payload.Status != 0 {
		errMessage = "Status should be active and inactive"
	}

	err = models.UpdateUser(&user)

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error update data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &Response{
			Code:    http.StatusOK,
			Message: "Success Delete user",
			Status:  false,
			Data:    user,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
}

// DeleteUser deletes a user by ID.
func DeleteUser(c *gin.Context) {
	var user models.User

	err := models.GetUser(&user, c.Param("id"))

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error get data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	if user.ID == 0 {
		response := &Response{
			Code:    http.StatusBadRequest,
			Message: "User not found",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	}

	err = models.DeleteUser(&user, c.Param("id"))

	if err != nil {
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Error delete data",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
		return
	} else {
		response := &Response{
			Code:    http.StatusOK,
			Message: "Success Delete user",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
}

func Login(c *gin.Context) {
	payload := &httpEntity.LoginUser{}

	if err := c.BindJSON(payload); err != nil {
		response := &Response{
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
		"email": payload.Email,
	}, user); err != nil {
		log.Fatal(err.Error())
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Failed Find Email",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if user.ID == 0 {
		log.Fatal("error: user not found")
		response := &Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "email not found",
			Status:  false,
			Data:    nil,
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		c.Abort()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		log.Fatal("error: user not found")
		response := &Response{
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
	response := &Response{
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
