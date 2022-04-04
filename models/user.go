package models

import (
	"time"

	"github.com/vincentJunior1/test-kriya/httpEntity"
)

// User is the main user model.
type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrementIncrement;type:int(10)unsigned" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(45)default null" json:"name"`
	UserName  string     `gorm:"column:user_name;type:varchar(45)default null" json:"user_name"`
	Password  string     `gorm:"column:password;type:varchar(100)default null" json:"password"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at"`
	CreatedBy int        `gorm:"column:created_by;type:bigint(20)not null" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updated_at"`
	UpdatedBy int        `gorm:"column:updated_by;type:bigint(20)not null" json:"updated_by"`
}

// GetUsers queries the database for all users.
func GetUsers(users *[]User, pagination *httpEntity.Pagination) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	DB.Table("users").Count(&pagination.TotalData)
	if err = DB.Joins("Role").Limit(pagination.Limit).Offset(offset).Find(users).Error; err != nil {
		return err
	}

	return nil
}

// CreateUser creates a new user.
func CreateUser(user *User) (err error) {
	if err = DB.Create(user).Error; err != nil {
		return err
	}
	DB.Joins("Role")
	DB.Find(user)
	return nil
}

// UpdateUser creates a new user.
func UpdateUser(user *User) (err error) {
	if err = DB.Joins("Role").Save(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUser queries the database for all users.
func GetUser(user *User, id string) (err error) {
	if err = DB.Joins("Role").Where("users.id = ?", id).Find(user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser queries the database for all users.
func DeleteUser(user *User, id string) (err error) {
	if err = DB.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}

// get user by params
func GetUserByParams(params map[string]interface{}, user *User) (err error) {
	if err = DB.Where(params).Find(user).Error; err != nil {
		return err
	}

	return nil
}
