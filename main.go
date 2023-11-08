package main

import (
	"Phase2/config"
	"Phase2/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	// db, err := config.InitDB()
	// if err != nil {
	// 	log.Fatalf("Error initializing the database: %s", err.Error())
	//}
	db := config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	log := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	userController := controller.NewUserDB(db)
	// actionController := controller.NewPostDB(db)

	g := e.Group("/user")
	{
		g.POST("/register", userController.Register)
		g.POST("/login", userController.Login)
	}

	// h := e.Group("/post")
	// h.Use(middleware.Authenticate)
	// {
	// 	h.POST("", actionController.Posting)
	// 	h.GET("", actionController.GetPosting)
	// 	h.GET(":id", actionController.GetPostingByID)
	// 	h.DELETE(":id", actionController.DeletePostByID)
	// }

	e.Start(":8080")
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}