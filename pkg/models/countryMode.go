package models

type CountryModel struct {
	Id        int    `json:"id"`
	Iso       string `json:"iso"`
	Name      string `json:"name"`
	NiceName  string `json:"nicename"`
	Iso3      string `json:"iso3"`
	NumCode   int    `json:"numcode"`
	PhoneCode string `json:"phonecode"`
}
