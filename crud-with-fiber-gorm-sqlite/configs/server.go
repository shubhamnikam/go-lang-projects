package configs

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/controllers"
)

func SetupServer() {
	app := fiber.New()
	
	SetupRoutes(app)
	
	log.Fatal(app.Listen(os.Getenv("SERVER_URL")))
}

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	
	apiV1.Get("/welcome", controllers.Welcome)

	//routes - users
	apiV1.Get("/users", controllers.GetAllUsers)
	apiV1.Get("/users/:id", controllers.GetUser)
	apiV1.Post("/users", controllers.CreateUser)
	apiV1.Put("/users/:id", controllers.UpdateUser)
	apiV1.Delete("/users/:id", controllers.DeleteUser)
	// routes - product
	apiV1.Get("/products", controllers.GetAllProducts)
	apiV1.Get("/products/:id", controllers.GetProduct)
	apiV1.Post("/products", controllers.CreateProduct)
	apiV1.Put("/products/:id", controllers.UpdateProduct)
	apiV1.Delete("/products/:id", controllers.DeleteProduct)
	// routes - order
	apiV1.Get("/orders", controllers.GetAllOrders)
	apiV1.Get("/orders/:id", controllers.GetOrder)
	apiV1.Post("/orders", controllers.CreateOrder)
	apiV1.Delete("/orders/:id", controllers.DeleteOrder)
}