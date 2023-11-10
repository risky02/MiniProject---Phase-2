package controller

import (
	"Phase2/dto"
	"Phase2/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderDB struct {
	DB *gorm.DB
}

func NewOrderDB(db *gorm.DB) OrderDB {
	return OrderDB{DB: db}
}

func (DB OrderDB)Payment(c echo.Context) error {
	fmt.Println("masuk 1")
	user := c.Get("user").(entity.User)
	fmt.Println("masuk 2")
	isPayment := dto.Payment{}
	if err := c.Bind(&isPayment); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid payment Request",
		})
	}
	
	
	checkout := entity.Checkout{}
	if err := DB.DB.Where("CheckoutID", isPayment.CheckoutID).Error; err != nil {
		return c.JSON(http.StatusNotFound, dto.ResponFailed{
			Code:    http.StatusNotFound,
			Message: "Checkout not found",
		})
	}

	newPayment := entity.Payment{
		UserID: user.ID,
		CheckoutID: checkout.Id,
		RentalDays: checkout.RentalDays, 
		TotalCost:checkout.TotalCost,
		Payment: isPayment.Payment,
	}

	if user.Deposit < newPayment.TotalCost {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Insufficient deposit amount",
		})
	}
	
	user.Deposit -= newPayment.TotalCost
	if err := DB.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update user deposit",
		})
	}
	return c.JSON(http.StatusOK, dto.ResponSuccess{
			Code:    http.StatusOK,
			Message: "Payment Success",
			Data: newPayment,
		})
}