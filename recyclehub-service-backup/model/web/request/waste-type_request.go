package request

// GetWasteTypeByIDRequest represents the request payload for retrieving a waste type by ID.
// @Description GetWasteTypeByIDRequest represents the ID of the waste type to retrieve.
type GetWasteTypeByIDRequest struct {
	ID string `json:"id" validate:"required"` // The ID of the waste type to retrieve
}
