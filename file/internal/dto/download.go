package dto

// DownloadResponse for download file response
type DownloadResponse struct {
	Error string   `json:"error"` // error Message
	Names []string `json:"names"` // uploaded File Names
}
