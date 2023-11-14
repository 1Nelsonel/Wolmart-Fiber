package routes

import (
	"github.com/1Nelsonel/Wolmart-Fiber/controllers"
	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/gofiber/fiber/v2"
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

	// FAQ routes
	faqController := controllers.NewFAQController(database.DBConn) 

	// CRUD operations for FAQs
	app.Get("/faqs-old/", faqController.Faqs)
	app.Post("/faqs/", faqController.CreateFAQ)
	app.Get("/faqs/", faqController.GetFAQs)
	app.Get("/faqs/:id", faqController.GetFAQ)
	app.Put("/faqs/:id", faqController.UpdateFAQ)
	app.Delete("/faqs/:id", faqController.DeleteFAQ)
}
