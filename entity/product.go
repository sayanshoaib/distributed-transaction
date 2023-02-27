package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Product struct {
	bun.BaseModel `bun:"products"`
	ID            int64     `bun:",pk,autoincrement"`
	ProductID     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name          string    `bun:",notnull"`
	Price         int64     `bun:",notnull"`
	Quantity      int64     `bun:",notnull"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `json:"-" bun:",soft_delete,nullzero"`
}
