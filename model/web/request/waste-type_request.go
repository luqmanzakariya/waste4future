package request

// GetWasteTypeByIDRequest represents the request payload for retrieving a waste type by ID.
type GetWasteTypeByIDRequest struct {
	ID string `json:"id" validate:"required"` // The ID of the waste type to retrieve
}
