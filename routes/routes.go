package routes

import (
	"cariIlmu-API/controller"
	"cariIlmu-API/middleware"
	"github.com/labstack/echo/v4"
)

func GetApiRoutes() *echo.Echo {
	e := echo.New()

	//routes for course categories
	e.GET("/api/course_categories", controller.GetAll, middleware.IsAuthenticated)
	e.POST("/api/course_categories", controller.PostCourseCategories)
	e.PUT("/api/course_categories", controller.UpdateCourseCategories)
	e.DELETE("/api/course_categories", controller.DeleteCourse_Categories)

	//create course
	e.GET("/courses", controller.ReadAllCourses, middleware.IsAuthenticated)
	e.POST("/courses", controller.CreateCourses)
	e.PUT("/courses", controller.UpdateCourses)
	e.DELETE("/api/courses", controller.DeleteCourses)

	//routes for login
	e.POST("/login", controller.CheckLogin)
	e.GET("/generate-hash/:password", controller.GenerateHashPassword)

	//routes for users_courses
	e.GET("/api/users_courses", controller.ReadAllUsersCourses, middleware.IsAuthenticated)
	e.POST("api/users_courses", controller.CreateUsersCourses)
	e.PUT("/api/users_courses", controller.UpdateUsersCourses)
	e.DELETE("/api/users_courses/:id", controller.DeleteUsersCourses)

	//routes for users
	e.GET("/api/users", controller.ReadAllUsers, middleware.IsAuthenticated)
	e.POST("/api/users", controller.CreateUsers)
	e.PUT("/api/users", controller.UpdateUsers)
	e.DELETE("/api/users", controller.DeleteUsers)

	//start server
	return e
}
