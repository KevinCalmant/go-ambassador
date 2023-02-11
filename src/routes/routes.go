package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)

	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Post("/products", controllers.CreateProduct)

	adminAuthenticated.Get("/ambassadors", controllers.Ambassadors)
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Get("/products", controllers.Products)
	adminAuthenticated.Get("/products/:id", controllers.GetProduct)

	adminAuthenticated.Put("/update-info", controllers.UpdateInfo)
	adminAuthenticated.Put("/update-password", controllers.UpdatePassword)
	adminAuthenticated.Put("/product/:id", controllers.UpdateProduct)

	adminAuthenticated.Delete("/products/:id", controllers.DeleteProduct)
}
