package main

import (
	"Phase2/config"
	"Phase2/controller"
	"Phase2/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %s", err.Error())
	}
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	middleware.LogrusConfig()

	userController := controller.NewUserDB(db)
	topUp := controller.NewTopupDB(db)
	inquiryProduct := controller.GetEquipmentDB(db)
	checkOut := controller.NewCheckoutDB(db)
	payMents := controller.NewOrderDB(db)
	authMiddleware := middleware.NewAuthDB(db)

	g := e.Group("/user")
	{
		g.POST("/register", userController.Register)
		g.POST("/login", userController.Login)
	}

	h := e.Group("/test")
	h.Use(authMiddleware.Authenticate)
	{
		h.POST("/deposit", topUp.Deposit)
		h.GET("/equipment", inquiryProduct.GetEquipment)
		h.GET("/equipment/:id", inquiryProduct.GetByID)
		h.POST("/equipment/checkout", checkOut.Checkout)
		h.POST("/equipment/payment", payMents.Payment)
	}

	e.Start(":8080")
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}