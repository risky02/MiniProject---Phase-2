package middleware

import (
	"Phase2/entity"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


type middlewareDB struct {
	DB *gorm.DB
}

func NewAuthDB(db *gorm.DB) middlewareDB {
	return middlewareDB{DB: db}
}

func (authDB middlewareDB) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
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
			userID, ok := claims["id"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID in token")
			}
			var user entity.User
			if err := authDB.DB.Where("id = ?", int(userID)).First(&user).Error; err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user data")
			}

			// Set user data to the context
			c.Set("user", user)
			
		}
		return next(c)
	}
}