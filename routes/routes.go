package routes

import (
	"github.com/labstack/echo/v4"
	"cariIlmu-API/controller"
)

func GetApiRoutes() *echo.Echo {
	e := echo.New()

	//routes for course categories
	e.GET("/api/course_categories", controller.GetAll)
	e.POST("/api/course_categories", controller.PostCourseCategories)
	e.PUT("/api/course_categories", controller.UpdateCourseCategories)
	e.DELETE("/api/course_categories", controller.DeleteCourse_Categories)

	//create course 
	e.GET("/courses", controller.ReadAllCourses)
	e.POST("/courses", controller.CreateCourses)
	e.PUT("/courses", controller.UpdateCourses)
	e.DELETE("/api/courses", controller.DeleteCourses)

	//routes for login
	e.POST("/api/login", controller.CheckLoginAdmin)
	e.GET("/generate-hash/:password", controller.GenerateHashPassword)

	//routes for users_courses 
	e.GET("/api/users_courses", controller.ReadAllUsersCourses)
	e.POST("api/users_courses", controller.CreateUsersCourses)
	e.PUT("/api/users_courses", controller.UpdateUsersCourses)
	e.DELETE("/api/users_courses/:id", controller.DeleteUsersCourses)

	//routes for users
	e.GET("/api/users", controller.ReadAllUsers)
	e.POST("/api/users", controller.CreateUsers)
	e.PUT("/api/users", controller.UpdateUsers)
	e.DELETE("/api/users", controller.DeleteUsers)

	//start server

	return e
}