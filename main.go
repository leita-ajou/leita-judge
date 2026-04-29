package main

import (
	"os"

	"leita/src/route"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

// @title			Leita API Docs
// @version		2.0.0
// @description	Leita Judge System API Documentation
// @BasePath		/api
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(swagger.New(swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "/api/swagger",
	}))

	if err := route.RegisterRoutes(app); err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(app.Listen(":" + os.Getenv("JUDGE_PORT")))
}
