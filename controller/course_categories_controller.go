package controller

import (
	"net/http"
	"strconv"
	"cariIlmu-API/models"
	"github.com/labstack/echo/v4"
)

func CreateCourseCategories(c echo.Context) error {

	//accept data from client
	name 		:= c.FormValue("name")

	//call function PostHero from models
	result, err := models.PostCourseCategories(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadAllCourseCategories(c echo.Context) error {
	result, err := models.FindAllCourseCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadAllCourseCategoriesById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindCourseCategoriesById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

	result, err := models.PutCourseCategories(conv_id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCourseCategories(c echo.Context) error{

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
	result, err := models.DeleteCourseCategories(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
