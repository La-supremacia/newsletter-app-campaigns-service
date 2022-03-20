package controllers

import (
	"campaigns-service/pkg/models"
	"campaigns-service/pkg/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

type HTTPError struct {
	Status  int    `db:"status" json:"status" validate:"" bson:"status"`
	Message string `db:"message" json:"message" validate:"message"`
}

// PostCreateCampaign func creates a new campaign.
// @Description  Create and associate a new campaign associated to an org by given params.
// @Summary      Create a campaign and associate it to a given org
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      202  {object} models.CreateCampaign_Response
// @Router       /v1/campaign/ [post]
func PostCreateCampaign(c *fiber.Ctx) error {
	u := new(models.CreateCampaign_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if u.CampaignName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign name is required")
	}
	if u.OrganizationId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a organization id is required")
	}

	createdCampaign := services.New_Campaign(u.OrganizationId, u.CampaignName, u.Contacts)
	err := mgm.Coll(createdCampaign).Create(createdCampaign)
	responseCampaign := services.New_CreateCampaign_Response(createdCampaign.ID.Hex(), createdCampaign.CampaignName)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully created a new Campaign", responseCampaign)
	return c.Status(fiber.StatusCreated).JSON(responseCampaign)
}

// PutEditCampaign func edit an existing campaign.
// @Description  Update an existing campaign associated, found by it's ID.
// @Summary      Update a campaign
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.EditCampaign_Response
// @Router       /v1/campaign/:id [PUT]
func PutEditCampaign(c *fiber.Ctx) error {
	id := c.Params("id")

	u := new(models.EditCampaign_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}

	baseModel := &models.Campaign{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	baseModel.CampaignName = u.CampaignName
	baseModel.CronjobPattern = u.CronjobPattern

	err = coll.Update(baseModel)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	fmt.Println("Successfully edited a Campaign", baseModel)
	return c.Status(fiber.StatusOK).JSON(baseModel)
}

// DeleteCampaign func remove an existing campaign.
// @Description  Remove an existing campaign associated, found by it's ID.
// @Summary      Remove a campaign by it's ID
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.DeleteCampaign_Response
// @Router       /v1/campaign/:id [DELETE]
func DeleteRemoveCampaign(c *fiber.Ctx) error {
	id := c.Params("id")

	u := new(models.DeleteCampaign_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}

	baseModel := &models.Campaign{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	err = coll.Delete(baseModel)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(fiber.Map{
		"sucess":  true,
		"message": "The contact was successfully deleted",
	})
}

// RetrieveCampaign func retrieve an existing campaign.
// @Description  Lookup a campaign based on a given ID.
// @Summary      Retrieve a campaign by it's ID
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.DeleteCampaign_Response
// @Router       /v1/campaign/:id [GET]
func GetRetrieveCampaign(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}

	baseModel := &models.Campaign{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(baseModel)
}

// PutAddContactToCampaign Add a contact to an existing campaign.
// @Description  Lookup a campaign by it's ID and append a contact to the suscription list.
// @Summary      Push a contact to a campaign contacts field
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.EditCampaign_Response
// @Router       /v1/campaign/:id [PUT]
func AppendContactToCampaign(c *fiber.Ctx) error {
	id := c.Params("id")

	u := new(models.EditCampaign_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}

	baseModel := &models.Campaign{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	baseModel.CampaignName = u.CampaignName
	baseModel.CronjobPattern = u.CronjobPattern

	err = coll.Update(baseModel)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	fmt.Println("Successfully added a contact to a Campaign", baseModel)
	return c.Status(fiber.StatusOK).JSON(baseModel)
}
