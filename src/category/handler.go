package category

import "github.com/gofiber/fiber/v2"

func HandlerCategories(c *fiber.Ctx) error {

	list, err := GetCategories()

	if err != nil {
		c.Status(fiber.ErrConflict.Code).JSON("Error while fetching the data")
	}

	return c.Status(fiber.StatusOK).JSON(list)
}
