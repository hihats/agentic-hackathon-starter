package models

// ShuffleRequest represents the incoming request for shuffling participants
type ShuffleRequest struct {
	Participants []string `json:"participants" binding:"required"`
	GroupSize    *int     `json:"group_size,omitempty"`
	NumGroups    *int     `json:"num_groups,omitempty"`
}

// ShuffleResponse represents the response with grouped participants
type ShuffleResponse struct {
	Groups [][]string `json:"groups"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
