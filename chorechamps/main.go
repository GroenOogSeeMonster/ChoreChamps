package main

import (
    "github.com/GroenOogSeeMonster/ChoreChamps/chorechamps/config"
    "github.com/GroenOogSeeMonster/ChoreChamps/chorechamps/handler"
    customMiddleware "github.com/GroenOogSeeMonster/ChoreChamps/chorechamps/middleware"
    "github.com/GroenOogSeeMonster/ChoreChamps/chorechamps/model"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

func main() {
	// Load configuration
	config.Init()

    // Add a delay before connecting to the database
    time.Sleep(10 * time.Second)

	// Set up the database
	db, err := gorm.Open(mysql.Open(config.DBConnection), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the database schema
	db.AutoMigrate(&model.Kid{}, &model.Rule{})

	// Set up Echo
	e := echo.New()

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(customMiddleware.AuthMiddleware(db))

	// Set up handlers
	kidHandler := handler.NewKidHandler(db)
	ruleHandler := handler.NewRuleHandler(db)

	// Define routes
	e.POST("/kids", kidHandler.Create)
	e.GET("/kids/:id", kidHandler.Get)
	e.GET("/kids", kidHandler.GetAll)
	e.PUT("/kids/:id", kidHandler.Update)
	e.DELETE("/kids/:id", kidHandler.Delete)
	e.POST("/rules", ruleHandler.Create)
	e.GET("/rules/:id", ruleHandler.Get)
	e.GET("/rules", ruleHandler.GetAll)
	e.PUT("/rules/:id", ruleHandler.Update)
	e.DELETE("/rules/:id", ruleHandler.Delete)

	// Start the server
	e.Start(":8282")
}