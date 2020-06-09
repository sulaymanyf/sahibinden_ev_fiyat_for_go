package models

import "time"

type Ev struct {
	ID                     string      `json:"id"`
	Shortname              string      `json:"shortname"`
	CategoryID             string      `json:"categoryId"`
	Title                  string      `json:"title"`
	Status                 string      `json:"status"`
	Live                   bool        `json:"live"`
	Price                  float64     `json:"price"`
	Currency               string      `json:"currency"`
	OriginalPrice          interface{} `json:"originalPrice"`
	OriginalCurrency       interface{} `json:"originalCurrency"`
	FormattedPrice         string      `json:"formattedPrice"`
	ClassifiedDate         time.Time   `json:"classifiedDate"`
	ExpiryDate             time.Time   `json:"expiryDate"`
	ImageURL               string      `json:"imageUrl"`
	ImageURLLargeThumbnail string      `json:"imageUrlLargeThumbnail"`
	ImageMainURL           string      `json:"imageMainUrl"`
	Latitude               float64     `json:"latitude"`
	Longitude              float64     `json:"longitude"`
	UserPenaltyScore       int         `json:"userPenaltyScore"`
	Attributes             struct {
		A22     string `json:"banyo_sayi"`
		A24     string `json:"fullm2"`
		A23     string `json:"isitma"`
		A108238 string `json:"depozito"`
		A106960 string `json:"balkon"`
		A104239 string `json:"aidat"`
		A812    string `json:"bina_yasi"`
		A810    string `json:"kat_sayi"`
		A107889 string `json:"netm2"`
		A811    string `json:"buldugu_kat"`
		A20     string `json:"oda_sayi"`
	} `json:"attributes"`
	Location string `json:"location"`
	OwnerID  int    `json:"ownerId"`
}
