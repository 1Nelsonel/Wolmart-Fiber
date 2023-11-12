package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/1Nelsonel/Wolmart-Fiber/models"
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



