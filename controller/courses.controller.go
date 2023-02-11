package controller

import (
	"cariIlmu-API/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func ReadAllCourses(c echo.Context) error {
	result, err := models.ReadAllCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCourses(c echo.Context) error {
	title := c.FormValue("title") 
	course_categories_id := c.FormValue("course_categories_id")

	result, err := models.PostCourses(title, course_categories_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCourses(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	course_categories_id := c.FormValue("course_categories_id")

	//convert string to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.PutCourses(conv_id, title, course_categories_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCourses(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.Delete_Courses(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}