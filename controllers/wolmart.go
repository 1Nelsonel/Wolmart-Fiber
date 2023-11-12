package controllers

import "github.com/gofiber/fiber/v2"

// HOME
func HomePage(c *fiber.Ctx) error {
	return c.Render("home", c, "partials/layout")
}

// About
func AboutUs(c *fiber.Ctx) error {
	return c.Render("about", c , "partials/layout")	
}

func Contact(c *fiber.Ctx) error {
	return c.Render("contact", c, "partials/layout")
}

// Blog
func Blogs(c *fiber.Ctx) error {
	return c.Render("blogs", c , "partials/layout")	
}