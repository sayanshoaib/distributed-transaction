package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Cart struct {
	bun.BaseModel `bun:"carts"`
	ID            int       `json:"id" bun:"id,pk,autoincrement"`
	CustomerID    string    `json:"customer_id" bun:"customer_id"`
	BranchID      int       `json:"branch_id" bun:"branch_id"`
	Products      []Product `json:"products" bun:"products"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `json:"-" bun:",soft_delete,nullzero"`
}
