package models

type SahibinRes struct {
	Response struct {
		Classifieds []struct {
			ClassifiedDate             string   `json:"classifiedDate"`
			Comparable                 bool     `json:"comparable"`
			Currency                   string   `json:"currency"`
			DetailedListViewAttributes []string `json:"detailedListViewAttributes"`
			ExpiryDate                 string   `json:"expiryDate"`
			FormattedPrice             string   `json:"formattedPrice"`
			ID                         string   `json:"id"`
			ImageInfo                  struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"imageInfo"`
			ImageMainURL string  `json:"imageMainUrl"`
			OwnerID      int64   `json:"ownerId"`
			Price        float64 `json:"price"`
			Shortname    string  `json:"shortname"`
			Title        string  `json:"title"`
		} `json:"classifieds"`
	} `json:"response"`
}
