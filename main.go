package main

import (
	"github.com/go-kulana/core"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

func index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func fetch(c *fiber.Ctx) error {
	domain := c.Params("domain")

	// decode domain
	decodedDomain, err := url.QueryUnescape(domain)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	info, err := core.GetAll(decodedDomain)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(info)
}

func main() {
	app := fiber.New()

	app.Get("/", index)
	app.Get("/ping", ping)
	app.Get("/fetch/:domain", fetch)

	app.Listen(":7000")
}
