package routes

import (
	"github.com/MEVLEPEIZ/Mini-Project/controllers"
	"github.com/MEVLEPEIZ/Mini-Project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes() *echo.Echo {

	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	// Login
	e.POST("/login", controllers.LoginsUsersController)

	// Logger
	m.LogMiddleware(e)

	// Auth
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUsersController)

	// JWT
	eJwt := e.Group("/api/v2")
	eJwt.GET("/users", controllers.GetUserDetailController)
	eJwt.GET("/users/:id", controllers.GetUserDetailController)
	eJwt.DELETE("/users/:id", controllers.GetUserDetailController)
	eJwt.PUT("/users/:id", controllers.GetUserDetailController)

	/*
	
	return e
}