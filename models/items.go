package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode" form:"itemCode"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
	OrderID     int
}
