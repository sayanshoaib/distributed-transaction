package service

import (
	"distributed-transaction/db"
	"distributed-transaction/entity"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetCustomer(ctx echo.Context) error {
	customer := new(entity.Customer)

	db := db.DatabaseConn()
	tx := db.NewSelect().Model(customer).Where("id = ?", ctx.Param("id"))
	return ctx.String(http.StatusOK, tx.String())
}

func SaveCustomer(ctx echo.Context) error {
	customer := new(entity.Customer)
	db := db.DatabaseConn()

	res, err := db.NewInsert().Model(customer).Exec(ctx.Request().Context())
	if err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(http.StatusCreated, res)
}
