package product

import (
	"github.com/gofiber/fiber/v2"
)

func HandlerProducts(c *fiber.Ctx) error {

	if c.Params("category") != "" {
		list, _ := GetByCategory(c.Params("category"))
		return c.Status(fiber.StatusOK).JSON(list)
		// => Hello john
	}
	list, _ := GetProducts()
	return c.Status(fiber.StatusOK).JSON(list)
}
