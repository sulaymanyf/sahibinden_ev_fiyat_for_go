package drivers

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	_db *gorm.DB
	err error
)

func init() {
	log.Print("open")
	_db, err = gorm.Open("sqlite3", "sahibinden.db")
	evBase := models.Kiralik{}
	satikik := models.Satikik{}
	satlikf := models.FiyatForSatlik{}
	kiralikf := models.FiyatForKiralik{}

	if !_db.HasTable(&evBase) {
		_db.CreateTable(&evBase)
	}

	if !_db.HasTable(&satikik) {
		_db.CreateTable(&satikik)
	}
	if !_db.HasTable(&satlikf) {
		_db.CreateTable(&satlikf)
	}
	if !_db.HasTable(&kiralikf) {
		_db.CreateTable(&kiralikf)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Print("----------init----------")
	log.Print(_db)
	_db.LogMode(true)

	//base := models.EvBase{
	//	ID:                     0,
	//	EvId:                   "802575011",
	//	ShortName:              " emlak-konut-kiralik-istanbul-emlak-iskeleye-6-dak-5-yillik-asansorlu-3-plus1-125m-2.kat-802575011",
	//	CategoryID:             "16624",
	//	Title:                  "İstanbul Emlak İskeleye 6 dak 5 yıllık asansörlü 3+1 125m 2.kat",
	//	Status:                 " waiting_approval",
	//	Live:                   true,
	//	Price:                  3400,
	//	Currency:               "TL",
	//	FormattedPrice:         " 3.400 TL",
	//	ClassifiedDate:         time.Now(),
	//	ExpiryDate:             time.Now(),
	//	ImageURL:               " https://i0.shbdn.com/photos/57/50/11/thmb_802575011aep.jpg",
	//	ImageURLLargeThumbnail: " https://i0.shbdn.com/photos/57/50/11/lthmb_802575011aep.jpg",
	//	ImageMainURL:           " https://i0.shbdn.com/photos/57/50/11/x5_802575011aep.jpg",
	//	BanyoSayisi:            " 1",
	//	Fullm2:                 " 130",
	//	Netm2:                  " 120",
	//	OdaSayi:                " 3+1",
	//	BinaYasi:               " 5-10 arası",
	//	BuluduguKat:            " 3",
	//	KatSayi:                " 5",
	//	Isitma:                 " Doğalgaz (Kombi)",
	//	Balkon:                 " Var",
	//	Esyali:                 " ",
	//	Aidat:                  " 50",
	//}

}

func GetDB() *gorm.DB {
	return _db
}
