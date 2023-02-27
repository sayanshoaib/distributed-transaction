package service

import (
	"distributed-transaction/apimodels"
	dao2 "distributed-transaction/dao"
	"distributed-transaction/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductService struct {
	repo *dao2.ProductDao
}

func NewProductService(productDAO *dao2.ProductDao) *ProductService {
	return &ProductService{
		repo: productDAO,
	}
}

func (s *ProductService) CreateProduct(ctx echo.Context, productId uuid.UUID, req *apimodels.ReqProductCreate) (*entity.Product, error) {

	res, err := s.repo.Create(ctx, productId, req)
	return res, err
}
