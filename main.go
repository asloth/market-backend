package main

import (
	"os"

	"github.com/asloth/market-backend/database"
	"github.com/asloth/market-backend/src/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "bsale_test:bsale_test@tcp(mdb-test.c6vunyturrl6.us-west-1.rds.amazonaws.com)/bsale_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.Handler = database.NewBaseHandler(db)

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/products/:category?", product.HandlerProducts)

	app.Post("/products/search", product.HandleSearchProducts)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Listen("0.0.0.0" + port)
}
