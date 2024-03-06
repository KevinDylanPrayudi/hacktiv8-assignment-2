package models

import (
	"reflect"
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

func (order *Order) AfterDelete(tx *gorm.DB) error {
	tx.Clauses(clause.Returning{}).Where("order_id = ?", order.ID).Delete(&Item{})
	return nil
}

func (order *Order) BeforeUpdate(tx *gorm.DB) error {
	items := order.Items
	var updatedItem Item
	for _, item := range items {
		if reflect.ValueOf(item).FieldByName("ID").IsZero() {
			return gorm.ErrRecordNotFound
		} else if err := tx.First(&updatedItem, item.ID).Error; err != nil {
			return err
		}
	}
	return nil
}
