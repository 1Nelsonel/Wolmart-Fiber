package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string         `json:"name"`
	Price         float64        `json:"price"`
	Description   string         `json:"description"`
	CategoryID    uint           `json:"category_id"`
	Category      Category       `json:"category" gorm:"foreignKey:CategoryID"`
	ManufactureID uint           `json:"manufacture_id"`
	Manufacture   Manufacture    `json:"manufacture" gorm:"foreignKey:ManufactureID"`
	VendorID      uint           `json:"vendor_id"`
	Vendor        Vendor         `json:"vendor" gorm:"foreignKey:VendorID"`
	Images        []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}

type Category struct {
	gorm.Model
	Name   string          `json:"name"`
	Images []CategoryImage `json:"images" gorm:"foreignKey:CategoryID"`
}

type Manufacture struct {
	gorm.Model
	Name   string             `json:"name"`
	Images []ManufactureImage `json:"images" gorm:"foreignKey:ManufactureID"`
}

type Vendor struct {
	gorm.Model
	Name   string        `json:"name"`
	Images []VendorImage `json:"images" gorm:"foreignKey:VendorID"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	FilePath  string `json:"file_path"`
}

type CategoryImage struct {
	gorm.Model
	CategoryID uint   `json:"category_id"`
	FilePath   string `json:"file_path"`
}

type ManufactureImage struct {
	gorm.Model
	ManufactureID uint   `json:"manufacture_id"`
	FilePath      string `json:"file_path"`
}

type VendorImage struct {
	gorm.Model
	VendorID uint   `json:"vendor_id"`
	FilePath string `json:"file_path"`
}
