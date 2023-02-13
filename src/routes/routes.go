package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	/**
	 * Admin routes
	 */
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	/**
	 * Admin authenticated routes
	 */
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Post("products", controllers.CreateProduct)

	adminAuthenticated.Get("ambassadors", controllers.Ambassadors)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Get("users/:id/links", controllers.Link)
	adminAuthenticated.Get("products", controllers.Products)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Get("orders", controllers.Orders)

	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)
	adminAuthenticated.Put("product/:id", controllers.UpdateProduct)

	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)

	/**
	 * Ambassador routes
	 */
	ambassador := api.Group("ambassador")
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)

	/**
	 * Ambassador authenticated routes
	 */
	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Post("logout", controllers.Logout)

	ambassadorAuthenticated.Get("users", controllers.User)

	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
}
