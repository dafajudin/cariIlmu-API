package controller

import (
	"net/http"
	"strconv"
	"cariIlmu-API/models"
	"github.com/labstack/echo/v4"
)

//Create Courses
func CreateCourses(c echo.Context) error {
	title := c.FormValue("title") 
	course_categories_id := c.FormValue("course_categories_id")

	result, err := models.Post_Courses(title, course_categories_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

//Read All Courses
func ReadAllCourses(c echo.Context) error {
	result, err := models.Read_AllCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

//Update Courses
func UpdateCourses(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	course_categories_id := c.FormValue("course_categories_id")

	//convert string to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.Put_Courses(conv_id, title, course_categories_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

//Delete Courses
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