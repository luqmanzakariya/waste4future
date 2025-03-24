package request

// RecycleHubCreateRequest represents the request payload for creating a new recycle hub.
// @Description RecycleHubCreateRequest represents the details required to create a recycle hub.
type RecycleHubCreateRequest struct {
	Name        string `json:"name" validate:"required"`          // Name of the recycle hub
	Phone       string `json:"phone" validate:"required"`         // Phone number of the recycle hub
	AddressID   string `json:"address_id" validate:"required"`    // Address ID of the recycle hub
	WasteTypeID string `json:"waste_type_id" validate:"required"` // Waste type ID associated with the recycle hub
}

// RecycleHubUpdateRequest represents the request payload for updating an existing recycle hub.
// @Description RecycleHubUpdateRequest represents the details required to update a recycle hub.
type RecycleHubUpdateRequest struct {
	Name        string `json:"name" validate:"required"`          // Updated name of the recycle hub
	Phone       string `json:"phone" validate:"required"`         // Updated phone number of the recycle hub
	AddressID   string `json:"address_id" validate:"required"`    // Updated address ID of the recycle hub
	WasteTypeID string `json:"waste_type_id" validate:"required"` // Updated waste type ID associated with the recycle hub
}
