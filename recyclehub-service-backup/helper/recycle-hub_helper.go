package helper

import (
	"recyclehub-service/model/domain"
	"recyclehub-service/model/web/response"
)

func RecycleHubToResponse(recycleHub domain.RecycleHub) response.RecycleHubResponse {
	return response.RecycleHubResponse{
		ID:          recycleHub.ID.Hex(),
		Name:        recycleHub.Name,
		Phone:       recycleHub.Phone,
		AddressID:   recycleHub.AddressID,
		WasteTypeID: recycleHub.WasteTypeID,
	}
}

func RecycleHubToResponses(recycleHubs []domain.RecycleHub) []response.RecycleHubResponse {
	var responses []response.RecycleHubResponse
	for _, recycleHub := range recycleHubs {
		responses = append(responses, RecycleHubToResponse(recycleHub))
	}
	return responses
}
