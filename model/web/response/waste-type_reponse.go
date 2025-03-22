package response

// WasteTypeResponse represents the response payload for a single waste type.
type WasteTypeResponse struct {
	ID    string  `json:"id"`    // Unique identifier for the waste type
	Name  string  `json:"name"`  // Name of the waste type
	Price float64 `json:"price"` // Price of the waste type
}

// WasteTypeListResponse represents the response payload for a list of waste types.
type WasteTypeListResponse struct {
	WasteTypes []WasteTypeResponse `json:"waste_types"` // List of waste types
}
