package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandlerProducts(c *fiber.Ctx) error {

	if c.Params("category") != "" {

		idCategory, err := strconv.ParseUint(c.Params("category"), 10, 32)

		if err != nil {
			c.Status(fiber.ErrBadRequest.Code).JSON("Params dont admited")
		}

		list, err := GetByCategory(idCategory)

		if err != nil {
			c.Status(fiber.ErrConflict.Code).JSON("Error while fetching the data")
		}
		return c.Status(fiber.StatusOK).JSON(list)

	}
	list, _ := GetProducts()
	return c.Status(fiber.StatusOK).JSON(list)
}
