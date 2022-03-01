package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type Campaign struct {
	mgm.DefaultModel `bson:",inline"`
	CampaignName     string   `db:"campaign_name" json:"campaign_name" validate:"required"`
	OrganizationId   string   `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	Description      string   `db:"description" json:"description" validate:"required"`
	Contacts         []string `db:"contacts" json:"contacts" validate:"" default:"[]"`
	CronjobPattern   string   `db:"cronjob_pattern" json:"cronjob_pattern" bson:"cronjob_pattern" validate:""`
	TemplateId       string   `db:"template_id" json:"template_id" bson:"template_id" validate:""`
}

type CreateCampaign_Request struct {
	CampaignName   string   `db:"campaign_name" json:"campaign_name" validate:"required"`
	CronjobPattern string   `db:"cronjob_pattern" json:"cronjob_pattern" bson:"cronjob_pattern" validate:""`
	OrganizationId string   `db:"organization_id" json:"organization_id" bson:"organization_id" validate:""`
	Contacts       []string `db:"contacts" json:"contacts" validate:""`
	TemplateId     string   `db:"template_id" json:"template_id" bson:"template_id" validate:""`
}

type CreateCampaign_Response struct {
	CampaignID   string `db:"campaign_id" json:"campaign_id" validate:""`
	CampaignName string `db:"name" json:"name" validate:"required"`
}

type EditCampaign_Request struct {
	CampaignName   string `db:"campaign_name" json:"campaign_name" validate:""`
	CronjobPattern string `db:"cronjob_pattern" json:"cronjob_pattern" bson:"cronjob_pattern" validate:""`
}

type DeleteOrganization_Request struct {
	OrganizationId string `db:"organization_id" json:"organization_id" validate:""`
}

type GetOrganization_Response struct {
	Id               string    `db:"_id" json:"_id" validate:""`
	CreatedAt        time.Time `db:"created_at" json:"created_at" validate:"" bson:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at" validate:""`
	OrganizationName string    `db:"organization_name" json:"organization_name" validate:"required"`
	UserId           string    `db:"user_id" json:"user_id" validate:"required"`
}

type DeleteOrganization_Response struct {
	Message string `db:"message" json:"message" validate:""`
	Success bool   `db:"success" json:"success" validate:""`
}
