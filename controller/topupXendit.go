package controller

import (
	"Phase2/dto"
	"Phase2/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	xendit "github.com/xendit/xendit-go/v3"
	invoice "github.com/xendit/xendit-go/v3/invoice"
	"gorm.io/gorm"
)

type TopupDB struct {
	DB *gorm.DB
}

func NewTopupDB(db *gorm.DB) TopupDB {
	return TopupDB{DB: db}
}

type topupData  struct {
	Deposit float32 `json:"deposit"`
}

func (t TopupDB)Deposit(c echo.Context) error {
	user := c.Get("user").(entity.User)
	topUp := topupData{}
	// log := logrus.New()
	
	if err := c.Bind(&topUp); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid deposit request",
		})
	}

	if err := t.DB.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, dto.ResponFailed{
			Code:    http.StatusNotFound,
			Message: "User not found",
		})
	}

	if topUp.Deposit < 100 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Amount minimum 100 rupiah",
		})
	}
	
	MaxDepositAmount := float32(100000000)
	if topUp.Deposit > MaxDepositAmount {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Amount maksimum transaksi 100.000.000",
		})
	}

	if err := sendDepositToXendit(c, user.ID, topUp.Deposit); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Error send data to Xendit",
		})
	}

	user.Deposit += float32(topUp.Deposit)

	if err := t.DB.Save(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Error Database: Failed update field deposit",
		})
	}

	return c.JSON(http.StatusOK, dto.ResponSuccess{
		Code:    http.StatusOK,
		Message: "Deposit Berhasil",
		Data:    user,
	})
}

func sendDepositToXendit(c echo.Context, user int, depositAmount float32) error {
	judulInvoice := fmt.Sprintf("Invoice order user id = %v", user)
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(judulInvoice, depositAmount)

	xenditClient := xendit.NewClient(os.Getenv("Secret_API_Key"))

	resp, r, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from CreateInvoice: Invoice
	fmt.Fprintf(os.Stdout, "Response from InvoiceApi.CreateInvoice: %v\n", resp)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses xendit",
		"respon":  resp,
	})
}
