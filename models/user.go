package models

import "github.com/vincentJunior1/test-kriya/httpEntity"

// User is the main user model.
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrementIncrement;type:int(10)unsigned" json:"id" `
	UserName string `gorm:"column:user_name;type:varchar(45)default null" json:"user_name"`
	Email    string `gorm:"column:email;type:varchar(100)default null" json:"email"`
	Password string `gorm:"column:password;type:varchar(100)default null" json:"password"`
	Status   uint   `gorm:"column:status;type:tinyint(4)default 1" json:"status"`
	RoleID   uint   `gorm:"column:role_id;type:int(10)unsigned not null" json:"role_id"`
	Role     *Roles `gorm:"references:role_id;foreignkey:id" json:"role"`
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
	if err = DB.Joins("Role").Create(user).Error; err != nil {
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
	if err = DB.First(user, id).Error; err != nil {
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
