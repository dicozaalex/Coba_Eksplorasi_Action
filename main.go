package main

import (
	"log"

	"github.com/Latihan/Eksplorasi/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New() // Echo instance
	a := 23

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.InsertUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	// Start server
	log.Println("Connected to port :6061")
	e.Logger.Fatal(e.Start(":6061"))asd
}
