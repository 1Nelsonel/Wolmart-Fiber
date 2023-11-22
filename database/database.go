package database

import (
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/1Nelsonel/Wolmart-Fiber/models"
)

var DBConn *gorm.DB

func ConnectDB() {
	// var err error

	// db, err := gorm.Open(sqlite.Open("wolmart.db"), &gorm.Config{})

	dsn := "host=localhost user=nelsonel01 password=0101 dbname=DbWolmart port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		danger := color.New(color.FgHiRed)
		danger.Println("Error connecting to the database.....")
	}
	green := color.New(color.FgGreen)
	green.Println("Database connected successfully.......")

	// Auto Migration
	if err := db.AutoMigrate(
		&models.Product{},
		&models.ProductImage{},
		&models.Category{},
		&models.CategoryImage{},
		&models.FAQ{},
		&models.Manufacture{},
		&models.ManufactureImage{},
		&models.Vendor{},
		&models.VendorImage{}); err != nil {
		danger := color.New(color.FgHiRed)
		danger.Println("Auto migration failed:: ", err)
	} else {
		cyan := color.New(color.FgHiCyan)
		cyan.Println("Auto migration successful......")
	}

	DBConn = db
}
