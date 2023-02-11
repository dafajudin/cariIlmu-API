package controller

import (
	"cariIlmu-API/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	result, err := models.ReadAllCourse_Categories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func PostCourseCategories(c echo.Context) error {

	//accept data from client
	name 		:= c.FormValue("name")

	//call function PostHero from models
	result, err := models.PostCourse_Categories(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCourseCategories(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.PutCourse_Categories(conv_id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCourse_Categories(c echo.Context) error{

	//accept id from client
	id := c.FormValue("id")

	//convert id to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	//call function DeleteCoursesCategories from models
	result, err := models.DeleteCourse_Categories(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
