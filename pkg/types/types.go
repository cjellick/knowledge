package types

// Dataset represents a new knowledge vector space
type Dataset struct {
	// Dataset ID - must be a valid RFC 1123 hostname
	ID             string `json:"id" format:"hostname_rfc1123" binding:"required,hostname_rfc1123" example:"asst-12345"`
	EmbedDimension *int   `json:"embed_dim" example:"1536" default:"1536" swaggertype:"integer"`
}

// Query represents an incoming user query
type Query struct {
	Prompt string `json:"prompt" binding:"required"`
	TopK   *int   `json:"topk" example:"5" swaggertype:"integer"`
}

// Ingest represents incoming content that should be ingested
type Ingest struct {
	Filename *string `json:"filename" `
	FileID   *string `json:"file_id"`
	Content  string  `json:"content" binding:"required,base64"`
}
