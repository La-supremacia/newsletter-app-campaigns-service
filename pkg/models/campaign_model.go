package models

import (
	"github.com/kamva/mgm/v3"
)

//Create campaign
type Campaign struct {
	mgm.DefaultModel `bson:",inline"`
	CampaignName     string `db:"campaign_name" json:"campaign_name" bson:"campaign_name" validate:"required"`
	OrganizationId   string `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	Description      string `db:"description" json:"description" validate:"required"`
	TemplateId       string `db:"template_id" json:"template_id" bson:"template_id" validate:""`
}

type CreateCampaign_Request struct {
	CampaignName   string `db:"campaign_name" json:"campaign_name" validate:"required"`
	OrganizationId string `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	TemplateId     string `db:"template_id" json:"template_id" bson:"template_id" validate:""`
	Description    string `db:"description" json:"description" validate:"required"`
}

type CreateCampaign_Response struct {
	CampaignID   string `db:"campaign_id" json:"campaign_id" validate:""`
	CampaignName string `db:"name" json:"name" validate:"required"`
}

//Create campaign

//Edit Campaign++

type EditCampaign_Request struct {
	CampaignName string `db:"campaign_name" json:"campaign_name" validate:""`
	TemplateId   string `db:"template_id" json:"template_id" bson:"template_id" validate:""`
	Description  string `db:"description" json:"description" validate:"required"`
}

type EditCampaign_Response struct {
	CampaignName   string `db:"campaign_name" json:"campaign_name" validate:"required"`
	OrganizationId string `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	Description    string `db:"description" json:"description" validate:"required"`
	TemplateId     string `db:"template_id" json:"template_id" bson:"template_id" validate:""`
}

//Edit Campaign--
type AppendConcactToCampaign struct {
	CampaignName   string `db:"campaign_name" json:"campaign_name" validate:"required"`
	OrganizationId string `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	TemplateId     string `db:"template_id" json:"template_id" bson:"template_id" validate:""`
}

type GetCampaign_Response struct {
	CampaignID string `db:"campaign_id" json:"campaign_id" validate:""`
}

type DeleteCampaign_Response struct {
	Message string `db:"message" json:"message" validate:""`
	Success bool   `db:"success" json:"success" validate:""`
}

type Campaign_Pivot_Contact struct {
	mgm.DefaultModel `bson:",inline"`
	CampaignId       string `db:"campaign_id" json:"campaign_id" bson:"campaign_id" validate:""`
	ContactId        string `db:"contact_id" json:"contact_id" bson:"contact_id" validate:""`
}
type Campaign_Pivot_Contact_Request struct {
	CampaignId string `db:"campaign_id" json:"campaign_id" bson:"campaign_id" validate:""`
	ContactId  string `db:"contact_id" json:"contact_id" bson:"contact_id" validate:""`
}

type Campaign_Pivot_Contact_Response struct {
	CampaignId string `db:"campaign_id" json:"campaign_id" bson:"campaign_id" validate:""`
	ContactId  string `db:"contact_id" json:"contact_id" bson:"contact_id" validate:""`
}
