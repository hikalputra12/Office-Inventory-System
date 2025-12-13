package model

import "time"

/* Inventory represents an inventory item. */
type Inventory struct {
	Model
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
}
