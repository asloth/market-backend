package product

import (
	"fmt"

	"github.com/asloth/market-backend/database"
)

type Product struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Url_image string
	Price     float64
	Discount  uint
	Category  uint
}

func GetProducts() (*[]Product, error) {
	var listProduct []Product

	result := database.Handler.Db.Table("product").Find(&listProduct)

	fmt.Println(listProduct)

	return &listProduct, result.Error
}

func GetByName(name string) (*[]Product, error) {
	var listProducts []Product
	p := Product{
		Id:        1,
		Name:      "por nombre",
		Url_image: "sdda",
		Price:     90.9,
		Discount:  12,
		Category:  2,
	}
	listProducts = append(listProducts, p)

	return &listProducts, nil
}

func GetByCategory(category string) (*[]Product, error) {
	var listProducts []Product
	p := Product{
		Id:        1,
		Name:      "categoriiis",
		Url_image: "sdda",
		Price:     90.9,
		Discount:  12,
		Category:  2,
	}
	listProducts = append(listProducts, p)

	return &listProducts, nil
}
