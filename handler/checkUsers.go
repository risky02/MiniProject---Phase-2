package handler

import (
	"Phase2/dto"
	"Phase2/entity"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func RegisterValidate(c echo.Context, validRegist entity.User) error {
	if len(validRegist.FullName) < 3 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Minimal Fullname 3 characters",
		})
	}

	if !isEmailValid(validRegist.Email) {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid email format",
		})
	}

	if len(validRegist.Username) < 3 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Minimal username 3 characters",
		})
	}

	if len(validRegist.Password) < 6 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Minimal Password 6 characters",
		})
	}
	return nil
}