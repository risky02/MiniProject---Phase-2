package controller

import (
	"Phase2/dto"
	"Phase2/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CheckoutDB struct {
	DB *gorm.DB
}

func NewCheckoutDB(db *gorm.DB) CheckoutDB {
	return CheckoutDB{DB: db}
}

func (DB CheckoutDB)Checkout(c echo.Context) error {
	user := c.Get("user").(entity.User)
	isCheckout := dto.Checkout{}
	if err := c.Bind(&isCheckout); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid checkout Request",
		})
	}

	equipment := entity.Equipment{}
	if err := DB.DB.Where("equipment_id = ?", isCheckout.EquipmentId).First(&equipment).Error; err != nil {
		return c.JSON(http.StatusNotFound, dto.ResponFailed{
			Code:    http.StatusNotFound,
			Message: "Equipment not found",
		})
	}

	rentalDate, err := time.Parse("2006-01-02", isCheckout.RentalDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid Rental Date",
		})
	}

	returnDate, err := time.Parse("2006-01-02", isCheckout.ReturnDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid Return Date",
		})
	}

	rentalDays := int(returnDate.Sub(rentalDate).Hours() / 24)
	totalCost := float32(rentalDays) * equipment.Price
	rentalDaysString := strconv.Itoa(rentalDays)

	newCheckout := entity.Checkout{
		UserId: user.ID,
		EquipmentId: isCheckout.EquipmentId,
		RentalDate: isCheckout.RentalDate,
		ReturnDate: isCheckout.ReturnDate,
		RentalDays: rentalDaysString,
		TotalCost:totalCost,
	}

	err = DB.DB.Create(&newCheckout).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Gagal Checkout",
		})
	}
	return c.JSON(http.StatusCreated, dto.ResponSuccess{
		Code:    http.StatusCreated,
		Message: "Berhasil Checkout",
		Data:    newCheckout,
	})
}