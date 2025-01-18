package main

import (
	"fmt"
	"iot-logging/configs"
	"log"
	"net/http"

	"iot-logging/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	TemperatureJSON struct {
		Temp float32 `json:"temp" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	conf := configs.GetConfig().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		conf.Host,
		conf.User,
		conf.Password,
		conf.DBName,
		conf.Port,
		conf.SSLMode,
		conf.Schema,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/temperature", func(c echo.Context) (err error) {
		t := new(TemperatureJSON)
		if err = c.Bind(t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		}
		if err = c.Validate(t); err != nil {
			return err
		}
		tempRecord := models.Temperature{Temperature: t.Temp}
		result := db.Create(&tempRecord)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		return c.JSON(http.StatusOK, t)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
