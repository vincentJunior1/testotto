package models

import (
	"time"

	"github.com/vincentJunior1/test-kriya/httpEntity"
)

// User is the main user model.
type Transaction struct {
	ID           uint       `gorm:"primaryKey;autoIncrementIncrement;type:int(10)unsigned" json:"id"`
	MerchantID   int        `gorm:"column:merchant_id;type:int(40)not null" json:"merchant_id"`
	MerchantName string     `gorm:"column:merchant_name" json:"Merchant"`
	Omzet        int        `gorm:"column:omzet" json:"omzet"`
	OutletID     int        `gorm:"column:outlet_id;type:int(40)not null" json:"merchant_name"`
	BillTotal    int        `gorm:"column:bill_total;type:double not null" json:"bill_total"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at"`
	CreatedBy    int        `gorm:"column:created_by;type:bigint(20)not null" json:"created_by"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updated_at"`
	UpdatedBy    int        `gorm:"column:updated_by;type:bigint(20)not null" json:"updated_by"`
}

func (Transaction) TableName() string {
	return "Transactions"
}

func GetMerchantOmzet(merchantID int, startDate time.Time, endDate time.Time, userID uint, transaction *[]Transaction, pagination *httpEntity.Pagination) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	query := DB.Model(transaction)
	query = query.Select("SUM(Transactions.bill_total) as omzet,Transactions.*, Merchants.merchant_name as merchant_name, Merchants.user_id = user_id")
	query = query.Joins("left join Merchants ON merchant_id = Merchants.id")
	query = query.Where("merchant_id = ? AND Transactions.created_at BETWEEN ? AND ? AND user_id = ?", merchantID, startDate, endDate, userID)
	query = query.Limit(pagination.Limit).Offset(offset).Group("DATE(Transactions.created_at)").Group("Transactions.created_at ASC")
	pagination.TotalData = query.Find(transaction).RowsAffected
	pagination.Sort = "transaction created_at ASC"
	query.Find(transaction)
	return query.Error
}
