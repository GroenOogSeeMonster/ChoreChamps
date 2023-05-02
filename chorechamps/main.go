package main

import (
	"github.com/GroenOogSeeMonster/chorechamps/config"
	"github.com/GroenOogSeeMonster/chorechamps/handler"
	"github.com/GroenOogSeeMonster/chorechamps/middleware"
	"github.com/GroenOogSeeMonster/chorechamps/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Init()

	// Set up the database
	db, err := gorm.Open(mysql.Open(config.DBConnectionString()), &gorm.Config{})
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
	e.Use(ChoreChamps_middleware.AuthMiddleware(db))

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