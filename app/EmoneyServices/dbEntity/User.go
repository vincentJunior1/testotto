package dbEntity

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrementIncrement;type:int(10)unsigned" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(45)default null" json:"name"`
	Email     string     `gorm:"column:email;type:varchar(45)default null" json:"email"`
	Password  string     `gorm:"column:password;type:text default null" json:"password"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updated_at"`
}

func (e *User) TableName() string {
	return "users"
}
