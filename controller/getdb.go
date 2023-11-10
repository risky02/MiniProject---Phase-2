package controller

import (
	"Phase2/dto"
	"Phase2/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetDB struct {
	DB *gorm.DB
}

func GetEquipmentDB(db *gorm.DB) GetDB {
	return GetDB{DB: db}
}

func (g GetDB) GetEquipment(c echo.Context) error {
	var getEquipment []entity.Equipment
	if err := g.DB.Find(&getEquipment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get all data",
		})
	}

	return c.JSON(http.StatusOK, dto.ResponSuccess{
		Code:    http.StatusOK,
		Message: "Success to get all data",
		Data:    getEquipment,
	})
}

func (g GetDB) GetByID(c echo.Context) error {
    getID := c.Param("id")

    var idEquipment entity.Equipment
    if err := g.DB.Where("equipment_id = ?", getID).First(&idEquipment).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.JSON(http.StatusNotFound, dto.ResponFailed{
                Code:    http.StatusNotFound,
                Message: "Data not found",
            })
        }
        return c.JSON(http.StatusInternalServerError, dto.ResponFailed{
            Code:    http.StatusInternalServerError,
            Message: "Failed to get data",
        })
    }

    return c.JSON(http.StatusOK, dto.ResponSuccess{
        Code:    http.StatusOK,
        Message: "Success to get data",
        Data:    idEquipment,
    })
}
