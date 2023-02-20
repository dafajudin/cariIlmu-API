package controller

import (
	"net/http"
	"strconv"
	"cariIlmu-API/models"
	"github.com/labstack/echo/v4"
)

func CreateUsersCourses(c echo.Context) error {
	user_id := c.FormValue("user_id")
	course_id := c.FormValue("course_id")

	conv_id, err := strconv.Atoi(user_id) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_course_id, err := strconv.Atoi(course_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.PostUserCourses(conv_id, conv_course_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadAllUsersCourses(c echo.Context) error {
	result, err := models.FindAllUserCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadUsersCoursesById(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FindUserCoursesById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsersCourses(c echo.Context) error {
	id := c.FormValue("id")
	users_id := c.FormValue("users_id")
	course_id := c.FormValue("course_id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_users_id, err := strconv.ParseInt(users_id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_course_id, err := strconv.ParseInt(course_id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.PutUserCourses(conv_id, int(conv_users_id), int(conv_course_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUsersCourses(c echo.Context) error {
	//accept id from client
	id := c.FormValue("id")

	//convert id to int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// call function Delete_UsersCourses from models
	result, err := models.DeleteUsersCourses(conv_id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	return c.JSON(http.StatusOK, result)
}

