package entities

type Countries struct {
	CountryID   int    `json:"country_id" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
	CountryName string `json:"country_name" binding:"required"`
}
