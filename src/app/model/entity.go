package model

import (
	"time"
)

// CREATE MODEL PRODUCT
type Product struct {
	ID        uint
	Name      string
	Qty       int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
