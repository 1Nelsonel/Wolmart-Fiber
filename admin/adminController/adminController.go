package admincontroller

import (
	"fmt"
	

	"github.com/1Nelsonel/Wolmart-Fiber/database"
	"github.com/1Nelsonel/Wolmart-Fiber/models"
	"github.com/gofiber/fiber/v2"
)

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


// Category
func AddCategory(c *fiber.Ctx) error {
	db := database.DBConn
	// Get category name from form
	categoryName := c.FormValue("categoryName")

	// Get category image from form
	file, err := c.FormFile("categoryImage")
	if err != nil {
		return err
	}

	// Save category image
	filePath := fmt.Sprintf("/adminAssets/uploads/category%s", file.Filename)

	// Save category image
	if err := c.SaveFile(file, "."+filePath); err != nil {
		return err
	}

	// Create a new category instance
	newCategory := models.Category{
		Name: categoryName,
	}

	// Save the category to the database
	if err := db.Create(&newCategory).Error; err != nil {
		return err
	}

	// Create a new category image instance
	newCategoryImage := models.CategoryImage{
		CategoryID: newCategory.ID,
		FilePath:   filePath,
	}

	// Save the category image to the database
	if err := db.Create(&newCategoryImage).Error; err != nil {
		return err
	}

	// Redirect or respond as needed
	return c.Redirect("/Category/")
}


func Category(c *fiber.Ctx) error {
	// Get the database connection
	db := database.DBConn // Assuming you have a database connection set up

	// Fetch all categories with their associated images from the database
	var categories []models.Category
	if err := db.Preload("Images").Find(&categories).Error; err != nil {
		return err
	}


	context := fiber.Map{
		"Categories": categories,
	}

	return c.Render("Category", context, "partials/adminLayout")
}