package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	adminroutes "github.com/1Nelsonel/Wolmart-Fiber/admin/adminRoutes"
	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/1Nelsonel/Wolmart-Fiber/routes"
	// "github.com/qinains/fastergoding"
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

	// fastergoding.Run() // hot reload

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "partials/layouts",
	})

	app.Static("assets", "./assets")
	app.Static("adminAssets", "./adminAssets")

	// Initialize default config
	app.Use(logger.New())

	app.Use(cors.New())

	// overide method delete middleware
	app.Use(func (c *fiber.Ctx) error {
		if c.Method()== fiber.MethodPost {
			if  override := c.FormValue("_method");
			override != "" {
				c.Method(override)
			}
		}
		return c.Next()
	})

	// Routes
	routes.SetupRoutes(app)
	routes.AuthSetupRoutes(app)
	adminroutes.AdminSetupRoutes(app)

	// app.Listen(":8080")
	app.Listen("0.0.0.0:8080")

}
