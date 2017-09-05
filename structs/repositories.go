package structs

type Repository struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Title             string `json:"title"`
	ColorLabel        string `json:"color_label"`
	Type              string `json:"type"`
	Url               string `json:"url"`
	RefreshWebhookUrl string `json:"refresh_webhook_url"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Repositories struct {
	Meta
	Entries []Repositories `json:"entries"`
}
