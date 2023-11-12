package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/1Nelsonel/Wolmart-Fiber/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.HomePage)
	app.Get("/shop/", controllers.Shop)
	app.Get("/product/", controllers.Product)
	app.Get("/about/", controllers.AboutUs)
	app.Get("/blogs/", controllers.Blogs)
	app.Get("/contact/", controllers.Contact)

	// Auth router
	app.Get("/login/", controllers.Login)
}
