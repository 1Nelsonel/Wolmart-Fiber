package controllers

import (
	"github.com/1Nelsonel/Wolmart-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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


// =============================FAQ======================================
type FAQController struct {
	DB *gorm.DB
}

// NewFAQController creates a new instance of FAQController
func NewFAQController(db *gorm.DB) *FAQController {
	return &FAQController{DB: db}
}

// FAQ
func (c *FAQController) Faqs(ctx *fiber.Ctx) error {
	var faqs []models.FAQ

	// Retrieve all FAQs from the database
	if err := c.DB.Find(&faqs).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to retrieve FAQs"})
	}
	context := fiber.Map{
		"Faqs":    faqs,
	}
	return ctx.Render("faqs", context , "partials/layout")	
}


// CreateFAQ creates a new FAQ entry
func (c *FAQController) CreateFAQ(ctx *fiber.Ctx) error {
	var faq models.FAQ

	if err := ctx.BodyParser(&faq); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	c.DB.Create(&faq)

	return ctx.JSON(fiber.Map{"message": "FAQ created successfully", "data": faq})
}

// GetFAQs retrieves all FAQs
func (c *FAQController) GetFAQs(ctx *fiber.Ctx) error {
	var faqs []models.FAQ

	c.DB.Find(&faqs)

	return ctx.JSON(fiber.Map{"data": faqs})
}

// GetFAQ retrieves a single FAQ by ID
func (c *FAQController) GetFAQ(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var faq models.FAQ

	if err := c.DB.First(&faq, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "FAQ not found"})
	}

	return ctx.JSON(fiber.Map{"data": faq})
}

// UpdateFAQ updates a FAQ by ID
func (c *FAQController) UpdateFAQ(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var existingFAQ models.FAQ

	if err := c.DB.First(&existingFAQ, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "FAQ not found"})
	}

	var updatedFAQ models.FAQ

	if err := ctx.BodyParser(&updatedFAQ); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	c.DB.Model(&existingFAQ).Updates(updatedFAQ)

	return ctx.JSON(fiber.Map{"message": "FAQ updated successfully", "data": existingFAQ})
}

// DeleteFAQ deletes a FAQ by ID
func (c *FAQController) DeleteFAQ(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var faq models.FAQ

	if err := c.DB.First(&faq, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "FAQ not found"})
	}

	c.DB.Delete(&faq)

	return ctx.JSON(fiber.Map{"message": "FAQ deleted successfully"})
}