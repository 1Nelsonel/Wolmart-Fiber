package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	BrandID     uint      `json:"brand_id"`
	Brand       Brand     `json:"brand" gorm:"foreignKey:BrandID"`
	VendorID    uint      `json:"vendor_id"`
	Vendor      Vendor    `json:"vendor" gorm:"foreignKey:VendorID"`
	Images      []Image   `json:"images" gorm:"foreignKey:ProductID"`
}

type Category struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images" gorm:"foreignKey:CategoryID"`
}

type Brand struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images" gorm:"foreignKey:BrandID"`
}

type Vendor struct {
	gorm.Model
	Name   string  `json:"name"`
	Images []Image `json:"images" gorm:"foreignKey:VendorID"`
}

type Image struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	CategoryID uint    `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	BrandID    uint    `json:"brand_id"`
	Brand      Brand   `json:"brand" gorm:"foreignKey:BrandID"`
	VendorID   uint    `json:"vendor_id"`
	Vendor     Vendor  `json:"vendor" gorm:"foreignKey:VendorID"`
	URL        string  `json:"url"`
}
