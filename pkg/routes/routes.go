package routes

import (
	"campaigns-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/campaign", controllers.PostCreateCampaign).Name("Create Campaign")
	route.Put("/campaign/:id", controllers.PutEditCampaign).Name("Edit Campaign")

	/*
		route.Delete("/campaign/:id", controllers.DeleteRemoveOrganization).Name("Remove Campaign")
		route.Get("/campaign/:id", controllers.GetRetrieveOrganizationbyId).Name("Retrieve Campaign By Id")
			route.Get("/organization", controllers.GetRetrieveOrganizationbyUserId).Name("Retrieve Organization By User Id")
		   	route.Post("/test", controllers.TestToken)
		   	route.Get("/", controllers.GetRoutes)
	*/
}
