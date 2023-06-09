package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/melatipus/UTS-EAI/internal/database"
	"github.com/melatipus/UTS-EAI/internal/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
