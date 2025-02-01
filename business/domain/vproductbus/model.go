package vproductbus

import (
	"time"

	"github.com/Klimentin0/courses-service1/business/types/money"
	"github.com/Klimentin0/courses-service1/business/types/name"
	"github.com/Klimentin0/courses-service1/business/types/quantity"
	"github.com/google/uuid"
)

// Product represents an individual product with extended information.
type Product struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        name.Name
	Cost        money.Money
	Quantity    quantity.Quantity
	DateCreated time.Time
	DateUpdated time.Time
	UserName    name.Name
}
