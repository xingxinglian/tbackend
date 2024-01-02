package data

type CollectionMetadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	CoverImage  string   `json:"cover_image"`
	SocialLinks []string `json:"social_links"`
}

type NftItemData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ExternalUrl string `json:"external_url"`
	Image       string `json:"image"`
	Marketplace string `json:"getgems.io"`
}
