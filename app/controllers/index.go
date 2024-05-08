package controllers

import (
	"app/cache"
	"app/libs"
	"app/validators"
	"time"

	"github.com/gofiber/fiber/v2"
)

var isPrimeMemoized = cache.Memoize(libs.IsPrime, 86400 * time.Second)

func PrimalityTest(c *fiber.Ctx) error {
	var body validators.Number

	if err := c.QueryParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"error": "Please enter a positive integer!"},
		)
	}

	if err := validators.Validator.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"error": "Please enter a positive integer!"},
		)
	}

	if !validators.PositiveIntRegex.MatchString(body.Value) {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"error": "Please enter a positive integer!"},
		)
	}

	result, err := isPrimeMemoized(body.Value)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"error": "Please enter a positive integer!"},
		)
	}

	return c.JSON(result)
}
