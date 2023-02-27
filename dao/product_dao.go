package dao

import (
	"distributed-transaction/apimodels"
	"distributed-transaction/entity"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IProductDao interface {
	Tx(tx *bun.Tx) (IProductDao, *bun.Tx, error)
	NewProductDao() *ProductDao
	CreateProduct(ProductID uuid.UUID, req *apimodels.ReqProductCreate) (*entity.Product, error)
	GetProductByID(ProductID uuid.UUID) (*entity.Product, error)
}
