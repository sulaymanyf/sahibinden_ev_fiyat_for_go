package models

import "time"

type Satikik struct {
	ID                     uint `gorm:"primary_key"`
	EvId                   int
	ShortName              string
	CategoryID             string
	Title                  string
	Status                 string
	Live                   bool
	Price                  float64
	Currency               string
	FormattedPrice         string
	ClassifiedDate         time.Time
	ExpiryDate             time.Time
	ImageURL               string
	ImageURLLargeThumbnail string
	ImageMainURL           string
	BanyoSayisi            string
	Fullm2                 string
	Netm2                  string
	OdaSayi                string
	BinaYasi               string
	BuluduguKat            string
	KatSayi                string
	Isitma                 string
	Balkon                 string
	Esyali                 string
	Aidat                  string
	CreatedAt              time.Time
}

type Kiralik struct {
	ID                     uint `gorm:"primary_key"`
	EvId                   int
	ShortName              string
	CategoryID             string
	Title                  string
	Status                 string
	Live                   bool
	Price                  float64
	Currency               string
	FormattedPrice         string
	ClassifiedDate         time.Time
	ExpiryDate             time.Time
	ImageURL               string
	ImageURLLargeThumbnail string
	ImageMainURL           string
	BanyoSayisi            string
	Fullm2                 string
	Netm2                  string
	OdaSayi                string
	BinaYasi               string
	BuluduguKat            string
	KatSayi                string
	Isitma                 string
	Balkon                 string
	Esyali                 string
	Aidat                  string
	CreatedAt              time.Time
}

type FiyatForSatlik struct {
	ID        uint `gorm:"primary_key"`
	EvId      int
	Fiyat     float64
	CreatedAt time.Time
}

type FiyatForKiralik struct {
	ID        uint `gorm:"primary_key"`
	EvId      int
	Fiyat     float64
	CreatedAt time.Time
}
