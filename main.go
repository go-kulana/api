package main

import (
	"encoding/base64"
	"fmt"
	"github.com/go-kulana/core"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
	"os"
	"time"
)

const logFile = "errors.log"

func writeToLog(e error) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	millis := now.Nanosecond() / 1000

	msg := fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d.%06d] %s\n", year, month, day, hour, minute, second, millis, e.Error())
	if _, err := f.WriteString(msg); err != nil {
		log.Println(err)
	}
}

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
		writeToLog(err)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	info, err := core.GetAll(decodedDomain)
	if err != nil {
		writeToLog(err)
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

	err := app.Listen(":7000")
	if err != nil {
		writeToLog(err)
		return
	}
}
