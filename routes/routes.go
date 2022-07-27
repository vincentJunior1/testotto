package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	EmoneyH "github.com/vincentJunior1/test-kriya/app/EmoneyServices/Handler"
	EmoneyRepo "github.com/vincentJunior1/test-kriya/app/EmoneyServices/Repository"
	EmoneyUc "github.com/vincentJunior1/test-kriya/app/EmoneyServices/Usecase"
	"gorm.io/gorm"
)

type Handlers struct {
	Ctx context.Context
	DB  *gorm.DB
	R   *gin.Engine
}

func NewHandlers(ctx context.Context, DB *gorm.DB, r *gin.Engine) *Handlers {
	return &Handlers{
		Ctx: ctx,
		DB:  DB,
		R:   r,
	}
}

// SetupRouter sets up the router.
func (h *Handlers) SetupRouter() {
	emoneyRepo := EmoneyRepo.NewEmoneyRepo(h.DB)

	emoneyUsecase := EmoneyUc.NewEmoneyUsecase(h.Ctx, *emoneyRepo)

	EmoneyH.EmoneyRoutes(emoneyUsecase, h.R)

}
