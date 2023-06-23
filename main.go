package main

import (
	"encoding/base64"
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

	// base64 decode domain
	base64DecodedDomain, err := base64.StdEncoding.DecodeString(domain)

	// decode domain
	decodedDomain, err := url.QueryUnescape(string(base64DecodedDomain))
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

	// configure cors
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, Token")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	})

	app.Get("/", index)
	app.Get("/ping", ping)
	app.Get("/fetch/:domain", fetch)

	app.Listen(":7000")
}
