package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Customer struct {
	bun.BaseModel `bun:"customers"`
	ID            int64     `bun:",pk,autoincrement"`
	Name          string    `bun:",notnull"`
	Balance       int64     `bun:",notnull"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `json:"-" bun:",soft_delete,nullzero"`
}
