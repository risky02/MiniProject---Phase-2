package middleware

import (
	"fmt"
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

func (udb UserDB) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please login to Access this page")
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("")
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("SECRET"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["id"].(int)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID in token")
			}
			var user int
			if err := udb.DB.Where("id = ?", int(userID)).First(&user).Error; err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user data")
			}

			// Set user data to the context
			c.Set("user", user)
		}

		return c.JSON(http.StatusUnauthorized, "Please login to Access this page")
	}
}