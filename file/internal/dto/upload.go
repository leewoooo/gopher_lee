package dto

// UploadResponse for upload file response
type UploadResponse struct {
	Error string   `json:"error"` // error Message
	Names []string `json:"names"` // uploaded File Names
}
