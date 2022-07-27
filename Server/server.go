package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vincentJunior1/test-kriya/models"
	"github.com/vincentJunior1/test-kriya/routes"
)

func Run() {
	r := gin.Default()
	c := context.Background()
	db := models.SetupDatabase()

	rh := &routes.Handlers{
		Ctx: c,
		DB:  db,
		R:   r,
	}
	rh.SetupRouter()

	r.Run()
}
