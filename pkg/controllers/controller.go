package controllers

import (
	"campaigns-service/pkg/models"
	"campaigns-service/pkg/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"go.mongodb.org/mongo-driver/bson"
)

type HTTPError struct {
	Status  int    `db:"status" json:"status" validate:"" bson:"status"`
	Message string `db:"message" json:"message" validate:"message"`
}

// PostCreateCampaign func creates a new campaign.
// @Description  Create and associate a new campaign associated to an org by given params.
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Summary      Create a campaign and associate it to a given org
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param campaign body models.CreateCampaign_Request true "campaign info"
// @Success      202  {object} models.CreateCampaign_Response
// @Router       /campaign/ [post]
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

	createdCampaign := services.New_Campaign(u.OrganizationId, u.CampaignName, u.Description, u.TemplateId)
	err := mgm.Coll(createdCampaign).Create(createdCampaign)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully created a new Campaign", err)
	return c.Status(fiber.StatusCreated).JSON(createdCampaign)
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
	baseModel.Description = u.Description
	baseModel.TemplateId = u.TemplateId

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
// @Success      200
// @Router       /v1/campaign/:id [DELETE]
func DeleteRemoveCampaign(c *fiber.Ctx) error {
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

	err = coll.Delete(baseModel)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(fiber.Map{
		"sucess":  true,
		"message": "The Campaign was successfully deleted",
	})
}

// RetrieveCampaign func retrieve an existing campaign.
// @Description  Lookup a campaign based on a given ID.
// @Summary      Retrieve a campaign by it's ID
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.Campaign
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
// @Summary      Create a relationship between a campaign and a contact
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.EditCampaign_Response
// @Router       /v1/campaign/:id [POST]
func AppendContactToCampaign(c *fiber.Ctx) error {
	id := c.Params("id")
	u := new(models.Campaign_Pivot_Contact_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}
	if u.ContactId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a contact ID is required")
	}

	pivotModel := &models.Campaign_Pivot_Contact{}
	coll := mgm.Coll(pivotModel)
	_, err := coll.UpdateOne(c.Context(), bson.M{"contact_id": u.ContactId, "campaign_id": id}, bson.D{
		{Key: "$set", Value: bson.D{
			bson.E{Key: "campaign_id", Value: id},
			bson.E{Key: "contact_id", Value: u.ContactId},
		},
		},
	}, mgm.UpsertTrueOption())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	fmt.Println("Successfully added a contact to a Campaign")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "The contact was successfully append to the campaign",
	})
}

// RemoveContactFromCampaign remove a  an existing campaign.
// @Description  Lookup a campaign by it's ID and remove a contact from the suscription list.
// @Summary      Delete the relationship between a campaign and a contact
// @Tags         Campaign
// @Accept       json
// @Produce      json
// @Param        id   path   string  true  "Campaign ID"
// @Success      200  {object} models.EditCampaign_Response
// @Router       /v1/campaign/:id [DELETE]
func RemoveContactFromCampaign(c *fiber.Ctx) error {
	id := c.Params("id")
	contactId := c.Params("contact")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a campaign ID is required")
	}
	if contactId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "a contact ID is required")
	}

	pivotModel := &models.Campaign_Pivot_Contact{}
	coll := mgm.Coll(pivotModel)
	coll.FindOneAndDelete(c.Context(), bson.M{"contact_id": contactId, "campaign_id": id})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "The contact was successfully removed from the campaign",
	})
}
