package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yigitnuhuz/gotodo/services"
)

func init() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoints
	e.GET("/", services.Hello)
	e.GET("/todos", services.AllTodos)
	e.POST("/todos", services.CreateTodo)

	e.GET("/todos/:id", services.GetTodo)
	e.PUT("/todos/:id/complete", services.UpdeteTodoIsComplete)
	e.PUT("/todos/:id/uncomplete", services.UpdeteTodoIsUncomplete)
	e.DELETE("/todos/:id", services.DeleteTodo)

	// Start server
	e.Logger.Fatal(e.Start(":3200"))
}

func main() {

}
