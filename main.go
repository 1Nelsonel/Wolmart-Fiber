package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/1Nelsonel/Wolmart-Fiber/routes"
)	

// Middleware to initialize db connections
func init() {
	database.ConnectDB()
}


func main() {
	// Database connect
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connection")
	}

	defer sqlDb.Close()

	// Create a new engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "partials/layouts",
	})

	app.Static("assets", "./assets")

	// Initialize default config
	app.Use(logger.New())

	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
