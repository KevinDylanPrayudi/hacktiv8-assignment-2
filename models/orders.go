package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	gorm.Model
	CustomerName string    `json:"customerName" form:"customerName"`
	Items        []Item    `json:"items"`
	CreatedAt    time.Time `json:"orderedAt" gorm:"column:orderedAt"`
}

func (order *Order) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("order_id = ?", order.ID).Delete(&Item{})
	return
}
