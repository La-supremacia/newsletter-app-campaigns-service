package main

import (
	mid "campaigns-service/pkg/middleware"
	"campaigns-service/pkg/routes"
	"campaigns-service/platform/database"

	_ "campaigns-service/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title           Campaigns microservice API
// @version         1.0
// @description     Create, Edit, Delete, Update Campaigns
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  lasupremaciadelpuntoycoma@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath  /api/v1
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	mid.FiberMiddleware(app)
	database.Init()
	routes.PublicRoutes(app)
	app.Listen(":3000")
}
