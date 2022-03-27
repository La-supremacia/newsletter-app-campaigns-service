package services

import "campaigns-service/pkg/models"

func New_Campaign(organizationId string, campaignName string, description string, templateId string) *models.Campaign {
	return &models.Campaign{
		CampaignName:   campaignName,
		OrganizationId: organizationId,
		Description:    description,
		TemplateId:     templateId,
	}
}

func New_CreateCampaign_Response(campaignId string, campaignName string) *models.CreateCampaign_Response {
	return &models.CreateCampaign_Response{
		CampaignID:   campaignId,
		CampaignName: campaignName,
	}
}

/*
func New_Campaign_Pivot_Contact(contactId string, campaignId string) *models.Campaign_Pivot_Contact {
	return &models.Campaign_Pivot_Contact{
		ContactId:  contactId,
		CampaignId: campaignId,
	}
}


func New_Campaign_Pivot_Contact_Response(contactId string, campaignId string) *models.Campaign_Pivot_Contact {
	return &models.Campaign_Pivot_Contact{
		ContactId:  contactId,
		CampaignId: campaignId,
	}
}

func New_EditOrganization_Response(organizationId string, organizationName string) *models.CreateOrganization_Response {
	return &models.CreateOrganization_Response{
		OrganizationId:   organizationId,
		OrganizationName: organizationName,
	}
} */
/* func New_DeleteOrganization_Response(message string, success bool) *models.DeleteOrganization_Response {
	return &models.DeleteOrganization_Response{
		Message: message,
		Success: success}
}
func New_DeleteOrganization_Request(organizationId string) *models.DeleteOrganization_Request {
	return &models.DeleteOrganization_Request{
		OrganizationId: organizationId,
	}
} */
