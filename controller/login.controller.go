package controller

import (
	"time"
	"net/http"
	"cariIlmu-API/helper"
	"cariIlmu-API/models"	
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckLogin(c echo.Context) error {

	//accept data from client
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	//call function CheckLogin from models
	res, err := models.CheckLogin(name, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// return c.JSON(http.StatusOK, "Login Success")

	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)

	// payload
	claims["name"] = name
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// //generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
