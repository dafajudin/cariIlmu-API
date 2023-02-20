package routes

import (
	"cariIlmu-API/controller"
	"cariIlmu-API/middleware"
	"github.com/labstack/echo/v4"
)

func GetApiRoutes() *echo.Echo {

	//create echo instance
	e := echo.New()

	//routes for course categories
	e.GET("/api/course_categories", controller.ReadAllCourseCategories, middleware.IsAuthenticated)
	e.GET("/api/course_categories/:id", controller.ReadAllCourseCategoriesById)
	e.POST("/api/course_categories", controller.CreateCourseCategories)
	e.PUT("/api/course_categories", controller.UpdateCourseCategories)
	e.DELETE("/api/course_categories", controller.DeleteCourseCategories)

	//create course
	e.GET("/api/courses", controller.ReadAllCourses, middleware.IsAuthenticated)
	e.GET("/api/courses/:id", controller.ReadCoursesById)
	e.POST("/api/courses", controller.CreateCourses)
	e.PUT("/api/courses", controller.UpdateCourses)
	e.DELETE("/api/courses", controller.DeleteCourses)

	//routes for login
	e.POST("/login", controller.CheckLogin)
	e.GET("/generate-hash/:password", controller.GenerateHashPassword)

	//routes for users_courses
	e.GET("/api/users_courses", controller.ReadAllUsersCourses, middleware.IsAuthenticated)
	e.GET("/api/users_courses/:id", controller.ReadUsersCoursesById)
	e.POST("api/users_courses", controller.CreateUsersCourses)
	e.PUT("/api/users_courses", controller.UpdateUsersCourses)
	e.DELETE("/api/users_courses", controller.DeleteUsersCourses)

	//routes for users
	e.GET("/api/users", controller.ReadAllUsers, middleware.IsAuthenticated)
	e.GET("/api/users/:id", controller.ReadUsersById)
	e.POST("/api/users", controller.CreateUsers)
	e.PUT("/api/users", controller.UpdateUsers)
	e.DELETE("/api/users", controller.DeleteUsers)
	
	return e
}
