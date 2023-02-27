package handlers

import (
	"distributed-transaction/apimodels"
	"distributed-transaction/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) (err error) {
	productID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Cannot generate UUID: %v", err.Error())
	}

	var req *apimodels.ReqProductCreate
	err = json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		log.Printf("cannot decode request body: %v", err.Error())
	}

	res, err := h.productService.CreateProduct(c, productID, req)
	if err != nil {
		log.Printf("%v", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
