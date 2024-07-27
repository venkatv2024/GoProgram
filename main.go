package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server Running")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Venkat")
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Hellow Env! " + os.Getenv("Test_Env"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
