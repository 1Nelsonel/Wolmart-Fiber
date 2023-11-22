package routes

import (
	"github.com/1Nelsonel/Wolmart-Fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthSetupRoutes(app *fiber.App) {
	// Auth router
	app.Get("/login/", controllers.Login)

}