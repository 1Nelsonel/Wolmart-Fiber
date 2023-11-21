package controllers

import (
	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/1Nelsonel/Wolmart-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

// ===========GET ALL PRODUCTS==========
func Shop(c *fiber.Ctx) error{
	db := database.DBConn
	var products []models.Product
	db.Find(&products)

	// c.JSON(http.StatusOK, products)
	return c.Render("shop", c, "partials/layout")
}

// Product
func Product(c *fiber.Ctx) error  {
	return c.Render("product", c, "partials/layout")
}

// Cart
func Cart(c *fiber.Ctx) error {
	return c.Render("cart", c, "partials/layout")
}

// checkout
func Checkout(c *fiber.Ctx) error {
	return c.Render("checkout", c, "partials/layout")
}

// compare
func Compare(c *fiber.Ctx) error {
	return c.Render("compare", c, "partials/layout")
}

// Whishlist
func Whishlist(c *fiber.Ctx) error {
	return c.Render("whishlist", c, "partials/layout")
}