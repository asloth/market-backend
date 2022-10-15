package product

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandlerProducts(c *fiber.Ctx) error {

	if c.Params("category") != "" {
		list, err := GetByCategory(c.Params("category"))

		if err != nil {
			fmt.Println(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(list)

	}
	list, _ := GetProducts()
	return c.Status(fiber.StatusOK).JSON(list)
}
