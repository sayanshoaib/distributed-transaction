package app

import (
	"distributed-transaction/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	EchoApp *echo.Echo
	Handler *handlers.ProductHandler
}

func NewServer(handler *handlers.ProductHandler) *Server {
	return &Server{
		echo.New(),
		handler,
	}
}

func (s *Server) Start() {

	//e.POST("/customers", service.SaveCustomer)
	//e.GET("/customers/:id", service.GetCustomer)
	s.EchoApp.POST("/products", handlers.ProductHandler{}.CreateProduct)
	//e.GET("/products/:id", getProduct)
	//e.POST("/carts", saveCart)
	//e.GET("/carts/:id", getCart)
	s.EchoApp.Use(middleware.Logger())

	s.EchoApp.Logger.Fatal(s.EchoApp.Start(":8080"))
}
