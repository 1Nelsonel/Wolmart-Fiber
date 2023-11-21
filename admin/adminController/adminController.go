package admincontroller

import "github.com/gofiber/fiber/v2"

// Dashboard
func Dashboard(c *fiber.Ctx) error  {
	return c.Render("dashboard", c, "partials/adminLayout")
}

// Add Product
func AddProduct(c *fiber.Ctx) error  {
	return c.Render("addProduct", c, "partials/adminLayout")
}

// adminProcuct.html
func AdminProducts(c *fiber.Ctx) error  {
	return c.Render("adminProducts", c, "partials/adminLayout")
}