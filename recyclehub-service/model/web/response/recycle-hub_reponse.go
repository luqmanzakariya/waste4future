package response

// RecycleHubResponse represents the response payload for a single recycle hub.
// @Description RecycleHubResponse represents a recycle hub with its ID, name, phone, address ID, and waste type IDs.
type RecycleHubResponse struct {
	ID          string   `json:"id"`            // Unique identifier for the recycle hub
	Name        string   `json:"name"`          // Name of the recycle hub
	Phone       string   `json:"phone"`         // Phone number of the recycle hub
	AddressID   string   `json:"address_id"`    // Address ID of the recycle hub
	WasteTypeID []string `json:"waste_type_id"` // Waste type ID associated with the recycle hub
}

// RecycleHubListResponse represents the response payload for a list of recycle hubs.
type RecycleHubListResponse struct {
	RecycleHubs []RecycleHubResponse `json:"recycle_hubs"` // List of recycle hubs
}
