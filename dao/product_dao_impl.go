package dao

import (
	"distributed-transaction/apimodels"
	"distributed-transaction/entity"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"log"
)

type ProductDao struct {
	Client *bun.DB
}

func NewProductDao(dbClient *bun.DB) *ProductDao {
	return &ProductDao{
		Client: dbClient,
	}
}

func (d *ProductDao) Create(ctx echo.Context, productId uuid.UUID, req *apimodels.ReqProductCreate) (*entity.Product, error) {
	var product entity.Product
	res, err := d.Client.NewInsert().Model(product).Exec(ctx.Request().Context())
	if err != nil {
		log.Printf("Product create failed: %v", err.Error())
	}

	fmt.Println(res)

	return &product, nil
}

func (d *ProductDao) GetProductByID(ctx echo.Context, productId uuid.UUID) (*entity.Product, error) {
	return nil, nil
}
