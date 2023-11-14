package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Description  string    `json:"description"`
	CategoryID   uint      `json:"category_id"`
	Category     Category  `json:"category"`
	BrandID      uint      `json:"brand_id"`
	Brand        Brand     `json:"brand"`
	VendorID     uint      `json:"vendor_id"`
	Vendor       Vendor    `json:"vendor"`
	Images       []Image   `json:"images"`
}

type Category struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type Brand struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type Vendor struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type Image struct {
	gorm.Model
	ProductID   uint   `json:"product_id"`
	Product     Product
	CategoryID  uint   `json:"category_id"`
	Category    Category
	BrandID     uint   `json:"brand_id"`
	Brand       Brand
	VendorID    uint   `json:"vendor_id"`
	Vendor      Vendor
	URL         string `json:"url"`
}