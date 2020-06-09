package models

import "time"

type ResponseRes struct {
	Success  bool `json:"success"`
	Response struct {
		ParentFacilityCategoryFacets []interface{} `json:"parentFacilityCategoryFacets"`
		PremiumClassified            interface{}   `json:"premiumClassified"`
		ShowPremiumBanner            bool          `json:"showPremiumBanner"`
		ShowProjectsClassified       bool          `json:"showProjectsClassified"`
		ShowTopListSeparator         bool          `json:"showTopListSeparator"`
		RealEstateProjects           interface{}   `json:"realEstateProjects"`
		NativeAdContent              interface{}   `json:"nativeAdContent"`
		FavoriteClassifieds          interface{}   `json:"favoriteClassifieds"`
		AttributeInfo                []interface{} `json:"attributeInfo"`
		TagAttributeInfo             []struct {
			Name          string `json:"name"`
			Label         string `json:"label"`
			Formatting    string `json:"formatting"`
			DisplayOrder  int    `json:"displayOrder"`
			CategoryLevel int    `json:"categoryLevel"`
		} `json:"tagAttributeInfo"`
		SortingInfo []struct {
			ID               string `json:"id"`
			Label            string `json:"label"`
			RequiresLocation bool   `json:"requiresLocation"`
		} `json:"sortingInfo"`
		CurrentSortingID string    `json:"currentSortingId"`
		LastUpdateTime   time.Time `json:"lastUpdateTime"`
		Paging           struct {
			PagingOffset int `json:"pagingOffset"`
			PagingSize   int `json:"pagingSize"`
			TotalResults int `json:"totalResults"`
		} `json:"paging"`
		Classifieds []struct {
			ID                     string      `json:"id"`
			Shortname              string      `json:"shortname"`
			CategoryID             string      `json:"categoryId"`
			Title                  string      `json:"title"`
			Description            interface{} `json:"description"`
			Status                 string      `json:"status"`
			Live                   bool        `json:"live"`
			Flags                  []string    `json:"flags"`
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
			ImageInfo              struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"imageInfo"`
			Latitude         float64 `json:"latitude"`
			Longitude        float64 `json:"longitude"`
			UserPenaltyScore int     `json:"userPenaltyScore"`
			Attributes       struct {
				A22     string `json:"a22"`
				A24     string `json:"a24"`
				A23     string `json:"a23"`
				A111578 string `json:"a111578"`
				A27     string `json:"a27"`
				A108238 string `json:"a108238"`
				A103651 string `json:"a103651"`
				A106960 string `json:"a106960"`
				A104239 string `json:"a104239"`
				A812    string `json:"a812"`
				A810    string `json:"a810"`
				A107889 string `json:"a107889"`
				A811    string `json:"a811"`
				A98426  string `json:"a98426"`
				A20     string `json:"a20"`
			} `json:"attributes"`
			CategoryBreadcrumb []struct {
				ID    int    `json:"id"`
				Label string `json:"label"`
			} `json:"categoryBreadcrumb"`
			TagAttributes struct {
				EmlakTipi string `json:"Emlak Tipi"`
			} `json:"tagAttributes"`
			Locations []struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Level int    `json:"level"`
			} `json:"locations"`
			Location         string `json:"location"`
			ActivePromotions []int  `json:"activePromotions"`
			Store            struct {
				ID   int         `json:"id"`
				Name string      `json:"name"`
				URL  string      `json:"url"`
				Logo interface{} `json:"logo"`
			} `json:"store"`
			OwnerID                    int         `json:"ownerId"`
			GeoRegionID                interface{} `json:"geoRegionId"`
			CategoryShortName          string      `json:"categoryShortName"`
			AuthorizedDealerClassified bool        `json:"authorizedDealerClassified"`
			AuthorizedBrands           interface{} `json:"authorizedBrands"`
			FromBankClassified         bool        `json:"fromBankClassified"`
			Comparable                 bool        `json:"comparable"`
			DetailedListViewAttributes []string    `json:"detailedListViewAttributes"`
		} `json:"classifieds"`
		CategoryBreadcrumbFacets []interface{} `json:"categoryBreadcrumbFacets"`
		StorecategoryFacets      []interface{} `json:"storecategoryFacets"`
		CategoryFacets           []interface{} `json:"categoryFacets"`
		SearchMeta               struct {
			Sections []struct {
				ID         interface{} `json:"id"`
				Name       string      `json:"name"`
				Label      string      `json:"label"`
				Formatting string      `json:"formatting"`
				Elements   []struct {
					Name                string        `json:"name"`
					CanonicalName       interface{}   `json:"canonicalName"`
					TypeID              interface{}   `json:"typeId"`
					Label               string        `json:"label"`
					Description         interface{}   `json:"description"`
					PostClassifiedLabel interface{}   `json:"postClassifiedLabel"`
					ElementType         string        `json:"elementType"`
					InputType           string        `json:"inputType"`
					DataType            string        `json:"dataType"`
					MultiSelect         bool          `json:"multiSelect"`
					MultiSelectSize     int           `json:"multiSelectSize"`
					MinLength           int           `json:"minLength"`
					MaxLength           int           `json:"maxLength"`
					MinValue            interface{}   `json:"minValue"`
					MaxValue            interface{}   `json:"maxValue"`
					EnumValues          []interface{} `json:"enumValues"`
					MapValues           struct {
					} `json:"mapValues"`
					TriggersRefresh             bool        `json:"triggersRefresh"`
					RefreshTarget               interface{} `json:"refreshTarget"`
					DependsOn                   interface{} `json:"dependsOn"`
					SummaryAttribute            bool        `json:"summaryAttribute"`
					ReadOnly                    bool        `json:"readOnly"`
					SortOrder                   int         `json:"sortOrder"`
					ShowOrder                   int         `json:"showOrder"`
					Filtered                    bool        `json:"filtered"`
					ShowOtherOption             bool        `json:"showOtherOption"`
					FacetType                   string      `json:"facetType"`
					Required                    bool        `json:"required"`
					ShowInSearchPage            bool        `json:"showInSearchPage"`
					WidgetName                  interface{} `json:"widgetName"`
					TimeLimitedEditableField    bool        `json:"timeLimitedEditableField"`
					SearchDocumentAttributeName interface{} `json:"searchDocumentAttributeName"`
					ExpandItsValues             bool        `json:"expandItsValues"`
					PriceEvaluationField        bool        `json:"priceEvaluationField"`
					Visible                     bool        `json:"visible"`
					CheckBox                    bool        `json:"checkBox"`
					SelectedUnitLabel           string      `json:"selectedUnitLabel"`
					DetailAttribute             bool        `json:"detailAttribute"`
					Address                     bool        `json:"address"`
					QueryText                   bool        `json:"queryText"`
					Boolean                     bool        `json:"boolean"`
					UnitSuffix                  string      `json:"unitSuffix,omitempty"`
					UnitEnumValues              []struct {
						ID                  string      `json:"id"`
						Label               string      `json:"label"`
						Selected            bool        `json:"selected"`
						DefaultName         interface{} `json:"defaultName"`
						PostClassifiedLabel interface{} `json:"postClassifiedLabel"`
						ElementMeta         float64     `json:"elementMeta"`
					} `json:"unitEnumValues,omitempty"`
				} `json:"elements"`
				ParentSection      interface{} `json:"parentSection"`
				SelectionCompleted bool        `json:"selectionCompleted"`
				Main               bool        `json:"main"`
				Group              bool        `json:"group"`
				Step               interface{} `json:"step"`
			} `json:"sections"`
			TagAttributes          interface{} `json:"tagAttributes"`
			CanonicalURL           string      `json:"canonicalUrl"`
			SearchURL              string      `json:"searchUrl"`
			Title                  string      `json:"title"`
			Description            string      `json:"description"`
			H1                     string      `json:"h1"`
			DefaultSorting         string      `json:"defaultSorting"`
			ApproximateResult      bool        `json:"approximateResult"`
			CommonAncestorCategory interface{} `json:"commonAncestorCategory"`
			CriteriaCategory       struct {
				ID                   int    `json:"id"`
				ParentID             int    `json:"parentId"`
				Name                 string `json:"name"`
				DisplayOrder         int    `json:"displayOrder"`
				Count                int    `json:"count"`
				CriteriaCategoryID   int    `json:"criteriaCategoryId"`
				CategoryIDBreadcrumb []int  `json:"categoryIdBreadcrumb"`
			} `json:"criteriaCategory"`
			GeoIPCityInfo         interface{} `json:"geoIPCityInfo"`
			AllowedAtCanonicalURL string      `json:"allowedAtCanonicalUrl"`
			BannedURLComponents   interface{} `json:"bannedUrlComponents"`
			RealCanonicalURL      interface{} `json:"realCanonicalUrl"`
			FormData              struct {
				AddressQuarter []string `json:"address_quarter"`
				PriceMin       []string `json:"price_min"`
				PagingSize     []string `json:"pagingSize"`
				A812           []string `json:"a812"`
				Sorting        []string `json:"sorting"`
				A811           []string `json:"a811"`
				Category       []string `json:"category"`
				PriceMax       []string `json:"price_max"`
			} `json:"formData"`
			StickyAttributes        []interface{} `json:"stickyAttributes"`
			RestrictedCategory      bool          `json:"restrictedCategory"`
			IgnoredFormElementNames []string      `json:"ignoredFormElementNames"`
		} `json:"searchMeta"`
		Warnings             interface{}   `json:"warnings"`
		MapData              interface{}   `json:"mapData"`
		Flags                []interface{} `json:"flags"`
		AdditionalParameters struct {
			ShoppingClassifiedsNearMeSectionTitle string `json:"ShoppingClassifiedsNearMeSectionTitle"`
		} `json:"additionalParameters"`
		CategoryTree []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			ClassifiedCount int    `json:"classifiedCount"`
			Children        []struct {
				ID              int    `json:"id"`
				Name            string `json:"name"`
				ClassifiedCount int    `json:"classifiedCount"`
				Children        []struct {
					ID              int    `json:"id"`
					Name            string `json:"name"`
					ClassifiedCount int    `json:"classifiedCount"`
					Children        []struct {
						ID              int           `json:"id"`
						Name            string        `json:"name"`
						ClassifiedCount int           `json:"classifiedCount"`
						Children        []interface{} `json:"children"`
					} `json:"children"`
				} `json:"children"`
			} `json:"children"`
		} `json:"categoryTree"`
		CanonicalUrls struct {
		} `json:"canonicalUrls"`
		VisibleComponents []string `json:"visibleComponents"`
	} `json:"response"`
}

type TekEvRes struct {
	Success  bool `json:"success"`
	Response struct {
		ID             int       `json:"id"`
		ShortName      string    `json:"shortName"`
		Title          string    `json:"title"`
		Price          float64   `json:"price"`
		Currency       string    `json:"currency"`
		ClassifiedDate time.Time `json:"classifiedDate"`
		Latitude       float64   `json:"latitude"`
		Longitude      float64   `json:"longitude"`
		Location       string    `json:"location"`
		URL            string    `json:"url"`
		Images         []struct {
			Thumbnail  string `json:"thumbnail"`
			Normal     string `json:"normal"`
			Large      string `json:"large"`
			X16        string `json:"x16"`
			W480       string `json:"w480"`
			ThreeSixty string `json:"three_sixty"`
			Main       bool   `json:"main"`
		} `json:"images"`
		Description string `json:"description"`
	} `json:"response"`
}

type Response struct {
	ID             int       `json:"id"`
	ShortName      string    `json:"shortName"`
	Title          string    `json:"title"`
	Price          float64   `json:"price"`
	Currency       string    `json:"currency"`
	ClassifiedDate time.Time `json:"classifiedDate"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Location       string    `json:"location"`
	URL            string    `json:"url"`
	Images         []struct {
		Thumbnail  string `json:"thumbnail"`
		Normal     string `json:"normal"`
		Large      string `json:"large"`
		X16        string `json:"x16"`
		W480       string `json:"w480"`
		ThreeSixty string `json:"three_sixty"`
		Main       bool   `json:"main"`
	} `json:"images"`
	Description string `json:"description"`
}
