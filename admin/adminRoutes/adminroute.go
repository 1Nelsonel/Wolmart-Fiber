package adminroutes

import (
	admincontroller "github.com/1Nelsonel/Wolmart-Fiber/admin/adminController"
	"github.com/gofiber/fiber/v2"
)

func AdminSetupRoutes(app *fiber.App) {
	app.Get("/dashboard/", admincontroller.Dashboard)
	app.Get("/addProduct/", admincontroller.AddProduct)
	app.Get("/adminProducts/", admincontroller.AdminProducts)
}