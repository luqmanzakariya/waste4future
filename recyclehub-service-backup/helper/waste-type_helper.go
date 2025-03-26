package helper

import (
	"recyclehub-service/model/domain"
	"recyclehub-service/model/web/response"
)

func WasteTypeToResponse(wasteType domain.WasteType) response.WasteTypeResponse {
	return response.WasteTypeResponse{
		ID:    wasteType.ID.Hex(),
		Name:  wasteType.Name,
		Price: wasteType.Price,
	}
}

func WasteTypeToResponses(wasteTypes []domain.WasteType) []response.WasteTypeResponse {
	var responses []response.WasteTypeResponse
	for _, wasteType := range wasteTypes {
		responses = append(responses, WasteTypeToResponse(wasteType))
	}
	return responses
}
