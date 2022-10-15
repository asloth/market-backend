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
	Discount  uint64
	Category  uint64
}

func GetProducts() (*[]Product, error) {
	var listProduct []Product

	result := database.Handler.Db.Table("product").Find(&listProduct)

	fmt.Println(listProduct)

	return &listProduct, result.Error
}

func GetByCategory(idCategory uint64) (*[]Product, error) {
	var listProduct []Product

	result := database.Handler.Db.Table("product").Find(&listProduct, Product{Category: idCategory})

	return &listProduct, result.Error
}

func GetByName(category string) (*[]Product, error) {

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
