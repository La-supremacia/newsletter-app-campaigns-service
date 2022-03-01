package services

import "campaigns-service/pkg/models"

func New_Campaign(organizationId string, campaignName string, contacts []string) *models.Campaign {
	return &models.Campaign{
		OrganizationId: organizationId,
		CampaignName:   campaignName,
		Contacts:       contacts,
	}
}

func New_CreateCampaign_Response(campaignId string, campaignName string) *models.CreateCampaign_Response {
	return &models.CreateCampaign_Response{
		CampaignID:   campaignId,
		CampaignName: campaignName,
	}
}

/*
func New_EditOrganization_Response(organizationId string, organizationName string) *models.CreateOrganization_Response {
	return &models.CreateOrganization_Response{
		OrganizationId:   organizationId,
		OrganizationName: organizationName,
	}
} */
func New_DeleteOrganization_Response(message string, success bool) *models.DeleteOrganization_Response {
	return &models.DeleteOrganization_Response{
		Message: message,
		Success: success}
}
func New_DeleteOrganization_Request(organizationId string) *models.DeleteOrganization_Request {
	return &models.DeleteOrganization_Request{
		OrganizationId: organizationId,
	}
}
