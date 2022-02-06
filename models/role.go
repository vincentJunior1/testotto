package models

type Roles struct {
	ID   uint   `gorm:"->;primaryKey" json:"id" `
	Name string `gorm:"column:name;type:varchar(100)default null" json:"name"`
}
