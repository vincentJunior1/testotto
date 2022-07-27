package Repository

import (
	"github.com/vincentJunior1/test-kriya/app/EmoneyServices/dbEntity"
	"gorm.io/gorm"
)

type EmoneyRepository struct {
	DB *gorm.DB
}

func NewEmoneyRepo(DB *gorm.DB) *EmoneyRepository {
	return &EmoneyRepository{
		DB: DB,
	}
}

func (e *EmoneyRepository) GetUserByParams(params map[string]interface{}, user *dbEntity.User) error {
	if err := e.DB.Where(params).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (e *EmoneyRepository) CreateUser(user *dbEntity.User) (err error) {
	if err = e.DB.Create(user).Error; err != nil {
		return err
	}
	e.DB.Find(user)
	return nil
}
