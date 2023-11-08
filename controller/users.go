package controller

import (
	"Phase2/dto"
	"Phase2/entity"
	"Phase2/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) UserDB {
	return UserDB{DB: db}
}

func (dbUser UserDB) Register(c echo.Context) error {
	registBody := dto.User{}
	if err := c.Bind(&registBody); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Invalid register Request",
		})
	}

	var eMail entity.User
	result := dbUser.DB.Where("email = ?", registBody.Email).First(&eMail)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar",
		})
	}

	var userName entity.User
	result = dbUser.DB.Where("username = ?", registBody.Username).First(&userName)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code:    http.StatusBadRequest,
			Message: "Username sudah terpakai",
		})
	}

	newRegister := entity.User{
		FullName: registBody.FullName,
		Email:    registBody.Email,
		Username: registBody.Username,
		Password: registBody.Password,
		Deposit: registBody.Deposit,
	}

	// validateRegist := handler.RegisterValidate(c, newRegister)
	// if validateRegist != nil {
	// 	return nil
	// }

	hashedPassword, err := helper.HashedPassword(newRegister.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Gagal Hash Password",
		})
	}

	newRegister.Password = hashedPassword

	err = dbUser.DB.Create(&newRegister).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menyimpan ke database",
		})
	} else {
		return c.JSON(http.StatusCreated, dto.ResponSuccess{
			Code:    http.StatusCreated,
			Message: "Registrasi Berhasil",
			Data:    newRegister,
		})
	}
}

func (dbUser UserDB) Login(c echo.Context) error {
	loginReq := dto.UserLogin{}
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Invalid Login Request",
		})
	}

	var loginUser entity.User
	result := dbUser.DB.Where("email = ?", loginReq.Email).First(&loginUser)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Email atau password anda salah",
		})
	}

	passwordCorrect := helper.CheckHashPassword(loginReq.Password, loginUser.Password)
	if passwordCorrect {
		return c.JSON(http.StatusBadRequest, dto.ResponFailed{
			Code: http.StatusBadRequest,
			Message: "Email atau password anda salah",
		})
	}

	token, err := helper.GenerateToken(jwt.MapClaims{
		"id": loginUser.ID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code: http.StatusInternalServerError,
			Message: "Gagal Generate Token",
		})
	}
	
	return c.JSON(http.StatusOK, dto.GetToken{
		Code: http.StatusOK,
		Message: "success Generate Token",
		Token: token,
	})
}