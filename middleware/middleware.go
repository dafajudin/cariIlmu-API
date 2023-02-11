package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

// To check if the user is authenticated or not
var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
