package main

import (
	"app/cache"
	"app/controllers"
	"app/middleware"
	"app/utils"

	"github.com/gofiber/fiber/v2"
)

func getApp() *fiber.App {
	utils.LoadEnv()
	cache.LoadRedis()

	app := fiber.New()

	// Middleware
	app.Use(middleware.RateLimiterMiddleware(5, 60))
	app.Use(middleware.RecoverMiddleware())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.HelmetMiddleware())
	app.Use(middleware.IdempotencyMiddleware())

	// Routes
	app.Get("/api/primes", controllers.PrimalityTest)

	return app
}

func main() {
	app := getApp()
	app.Listen(":8080")
}
