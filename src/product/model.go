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

func GetByName(name string) (*[]Product, error) {

	var listProduct []Product

	result := database.Handler.Db.Table("product").Find(&listProduct, "name like ?", "%"+name+"%")

	return &listProduct, result.Error
}

func GetByNameAndCategory(name string, idCategory uint64) (*[]Product, error) {

	var listProduct []Product

	result := database.Handler.Db.Table("product").Find(&listProduct, "name like ? AND category = ?", "%"+name+"%", idCategory)

	return &listProduct, result.Error
}
