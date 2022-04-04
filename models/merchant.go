package models

import (
	"time"
)

// User is the main user model.
type Merchant struct {
	ID           uint       `gorm:"primaryKey;autoIncrementIncrement;type:int(10)unsigned" json:"id"`
	UserID       int        `gorm:"column:user_id;type:int(40)not null" json:"user_id"`
	MerchantName string     `gorm:"column:merchant_name;type:varchar(45)default null" json:"merchant_name"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at"`
	CreatedBy    int        `gorm:"column:created_by;type:bigint(20)not null" json:"created_by"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updated_at"`
	UpdatedBy    int        `gorm:"column:updated_by;type:bigint(20)not null" json:"updated_by"`
}
