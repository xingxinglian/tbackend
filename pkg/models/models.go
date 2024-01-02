package models

import "gorm.io/gorm"

type CollectionMetadata struct {
	gorm.Model
	Name        string   `json:"name" gorm:"uniqueIndex"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	CoverImage  string   `json:"cover_image"`
	SocialLinks []string `json:"social_links" gorm:"type:text[]"`
}

type NftItem struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExternalUrl string `json:"external_url"`
	Marketplace string `json:"marketplace"`
}
