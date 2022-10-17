package category

import "github.com/asloth/market-backend/database"

type Category struct {
	Id   int
	Name string
}

func GetCategories() (*[]Category, error) {
	var listCategory []Category

	result := database.Handler.Db.Table("category").Find(&listCategory)

	return &listCategory, result.Error
}
