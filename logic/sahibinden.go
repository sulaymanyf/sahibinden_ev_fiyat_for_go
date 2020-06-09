package logic

import (
	"../drivers"
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetEvinFiyat() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.sahibinden.com/sahibinden-ral/rest/classifieds/search?a811=40598&a812=40604&price_currency=1&category=16633&price_max=1500000&a812=40728&a812=40602&a812=40603&a812=40605&a812=40606&a811=62367&a811=40594&a811=40596&a811=40600&a811=62364&a811=62366&a811=40593&a811=40595&a811=40597&a811=40599&a811=40601&a811=62365&address_quarter=51395&address_quarter=23114&address_quarter=51399&address_quarter=52036&address_quarter=51400&address_quarter=23097&sorting=price_desc&m%3AsearchMeta=true&m%3AincludeCategoryTree=false&m%3AincludeMatchingFavoriteClassifieds=true&m%3AsmartTextSearch=true&pagingOffset=0&pagingSize=200&language=tr", nil)
	CheckError(err)
	req.Header.Set("Host", "api.sahibinden.com")
	req.Header.Set("X-J3PopQvX-d", "o_0")
	req.Header.Set("X-J3PopQvX-c", "A28nYCVyAQAArBgpIQOHt07_tqEQnHt8XqC6ZHj0IKu7FpePpdxdohUx5OfbAR-NVzSfQDXk7xUO8YLqosIO8Q==")
	req.Header.Set("X-NewRelic-ID", "VQACUF9bCBAGXVNbAQQOUVc=")
	req.Header.Set("x-client-profile", "Generic_v2.1")
	req.Header.Set("x-timestamp", "1589765538355")
	req.Header.Set("X-J3PopQvX-b", "oi9as7")
	req.Header.Set("X-J3PopQvX-a", "m42X-5AiAbij0O5_qzkAAs2=c9rYSZOvoIxXLiTmj7arqo8=LYw0sJfNXO7SgSU-CgdlL4=EzmTmhIrvIrt9fwMSjzXDWjGIm4e=FtpgVBaQ=rLRtppZcw_wPmEo-Be6xZEEZHPHAjKLAAx3x8=JM0r5_4PZ6eijyM6sf6ztOl8QN3P0v5l8QcXL-eqHLlBax3Y7kz6=l7fINtEhDc0HooDBr7iqbCOp9RTDH1=cvOwWKbT3Q_lKsKCAVPKvoJOxvZ_cCN9IvY3I2JBSx0V9cK3JA-b_LZ4LIFHXmSkc1wT==69WieKPGnyFZ8xWr-VUTbk2JQVWfOF2eKmtXNivwQ-Rs-O3aY8B=W-N7js6N928qJLyEXsQgyq3iD-SVU1ygkxhdrQ9OkoModsMfbikvd2fa2CWaGmdb0zz_6y1tWH4lcoFAaStB7vbAdq6-pTbDkre5-C7MIqd7FdFD91c0FqX1qA87Q6eQPJ7HeYMrinWA8sLNRnvnXvJh=X2=UmM=kbFQIaokQ4JzjJJd5ZYHz92ZwKp_6aJ2UTYBR1NNn8B6W1TcC3LcQM3BdOimJh-MU8SRqqvTgpcFzDnAKUXrNjnKP83JUr1orsnzb9FDOoIA4amN2AO7-fGIjL7E8has2aKnVwCCUUshsGtXI7e618i3W11hNaJWCOaeQ5vsmR7e_BUfdXtMNXC9M-ZhG_OfWkNyIvkRxpBBB8cYIK_AbWPlyEflJZfaWBLW7s=JyjRfAGB0MGHHDOwyXrtebPosna9628I-k2=R4OXkOTW0jb1TyeJ6geD1KAijfIs0E9M4p1Bm3xyQ13PkvbHO_OSJgtq7WQd9nyelkBB=XLLz2qrn_3mm58U55exanPeFw8hdDoTNhc4o6PHjQTSSGY-B9psSeqaWUS0AHoJ7s51yvkdtyecJXxXxjerT_pri4dXR5NB_2eRFJQhE=Mwq_dhoGbpYer-iTJDTPWZZ6GJynNafEYGNqm-5LwQKqUi0H_nKisZM5AHg7Uj9JBLPZ8GwQOVt=BZR8QE8ecCUq4UgV6kH4Uj0vNp4=JnRj6m0wXQ2WwD6mKnwdPJrY8cAlfBIUnTWRncwf3sIG_yc6=rHUNWiGIBqV3F0d_r_m6V-xd7ZviF8wvxikERU5vhHizIo34shd2p4ZIb3i7s_x4rtelEZR8KYLLArE=xV37I7n9MMhJLNvoR21WK2wnC4LwLvw8YzQE5yCZezfD8RnxsNIFnaK-5B4CLM-iw_fzQFxDKIxZXgqCtX-4rrS8ohCEHDfq8rZxBdodgdywc2Oes7JPmApk9PSayKz=k63pdX7SBcI5fE=KXozwj-OCc2EjSjkZeoQsn0YwrhlcVABG9HJyna0=BmptHHwXrMA_ZMrmDIV42hOhJ5Qp1SiGeW--qstH2f=VzAgs1ay2jeHbNEQqTNQLRIDqGfO9wKAMyhzO=m-xQDnPDJCOaGqcdJ1cHZbsa8fVczp9jaSvIGKpAV=_0V0DDSZCoOt-FCJ1kk-tWtOn6BHHNDzkRI6FrnGKUnrYIrpUq0J6")
	req.Header.Set("User-Agent", "Sahibinden-IOS/4.7.0(2020/05/12,iPhone 6,iOS12.4.5)v3")
	req.Header.Set("x-api-hash", "925E6324164FACAD9ACE893223C2443C81DD9BF0")
	req.Header.Set("X-Device-Id", "8CD86B9902A14495B9880120076FE32C")
	req.Header.Set("X-J3PopQvX-z", "p")
	req.Header.Set("x-bundle-id", "com.sahibinden.sahibinden")
	req.Header.Set("x-app-version", "4.7.0")
	req.Header.Set("x-language", "tr")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("X-J3PopQvX-f", "A3upZCVyAQAAOWlZ7xCumu0gx2vqAvGj1pfqU3vrtDfBnKyxA2AAD7YFeMqAAWp6r-KfQDXk7xUO8YLqosIO8Q==")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-J3PopQvX-e", "a;vqNouuUHcJYfCyyZsd--FL4xkCouUV5UVtEBQrUSQSVaBwKOrLxtGdnrMMu6z2yVxG_D-XXbUuZuM1iOtdQJy88Omcz_kE3fQxnOdXbKTzbY3IYuwjAVYZyVKl1GyMJ2JgJFLpQ5mqisd5ulzv_xG9QdlOeOsyQmGSR3t7_UtzFZVn_gqeajUl9S9E_VLw1_E5gvinip2_Q4DtRffxhsXowj4C18SMVAlXl9SpdjMFLtY31uk9jo10kZ0aL7LAYQ5LxK9CQgEb-6h0WR_bz2jUjFCliVnYupn6xzANmufP9UgDgNB79h3KuXVTaiNYvSN13sEEAXoWcdo1hyTtL6ad8Ys9iIwn5U41kgArXzURXPJn6BLbID4CE3A0gbKUQldsAyNpMug5T5NNDYerKHl2KliwYwWFDRAhX8K2mCZPMS-GrUwJRz1fe814yaBi7SfKJzq46lGe9OB5C1JH9rdTzrA1aNPexTIFUPXpKH4sY2GON6d8ny7E8KqCaFXADkbXNdsc9KI8brBTzX5YfK4yYoxSi93U-bOw-1t3El9g9h4fZopJb2hP1Ri5gBuj1lbLAjcPa7hpUDAHlt7gNAdpga9h9-oVsnBf9bFPhYpIyrBWVMNpcnk6OicXkWrx1C5fn6NutGNqxCHI4AN1akvxFHQNNBzLWRURIo53SKP47gEVYSEUrZWW9NCcJ-uxLtQ40KH_EqZvly0jVNtMLx3E_Tm8iLtU017THpXYj43HF6bS4rQ-_JuZseUFkC4gx0Q9CT8KJURWVCupohychiBlFJIRY=;b2lU28_9AGDBvBdpLicbZ8JXOD7aS2okw8ClLYIJkWM=")
	req.Header.Set("x-api-key", "f51c6d9290ea9522aceac00f3b87a0002695d38a")
	req.Header.Set("idfa", "66642314-815E-4E43-BCE8-D1D8652315B6")
	resp, err := client.Do(req)
	CheckError(err)
	bodyText, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	fmt.Println(reflect.TypeOf(bodyText))
	toFile(bodyText)
	sahibin := models.SahibinRes{}
	err = json.Unmarshal(bodyText, &sahibin)
	CheckError(err)
	fmt.Println(reflect.TypeOf(sahibin))
	fmt.Println("m:", sahibin)
	fmt.Println("m.IP:", sahibin.Response.Classifieds)
	fmt.Println("m.Name:", sahibin.Response)
	b, err := json.Marshal(sahibin.Response.Classifieds)
	CheckError(err)
	fmt.Println(string(b))
	evler := make([]models.Daire, 0)
	daire := models.Daire{}
	for _, v := range sahibin.Response.Classifieds {
		daire.ID = v.ID
		daire.ClassifiedDate = v.ClassifiedDate
		daire.Comparable = v.Comparable
		daire.Currency = v.Currency
		daire.DetailedListViewAttributes = v.DetailedListViewAttributes
		daire.ExpiryDate = v.ExpiryDate
		daire.FormattedPrice = v.FormattedPrice
		daire.ImageInfo = v.ImageInfo
		daire.ImageMainURL = v.ImageMainURL
		daire.OwnerID = v.OwnerID
		daire.Price = v.Price
		daire.Shortname = v.Shortname
		daire.Title = v.Title

		//fmt.Println(daire)
		evler = append(evler, daire)
		//fmt.Println(len(evler))
	}

	//jsons, _ :=json.Marshal(evler)
	//fmt.Println(string(jsons))
	//toFile(bodyText)
	// fmt.Printf("%s\n", bodyText)
	sql := "insert into istanbul (ev_id, classified_date, comparable, currency, expiry_date, formatted_price, image_main_url, price, shortname, title,creat_time) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)"
	for _, v := range evler {
		stmt, err := drivers.MysqlDb.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)

		res, err := stmt.Exec(v.ID, v.ClassifiedDate, v.Comparable, v.Currency, v.ExpiryDate, v.FormattedPrice, v.ImageMainURL, v.Price, v.Shortname, v.Title, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	}
}

func toFile(bodyText []byte) {
	// fileName := time.Now().Format("2006-01-02")
	// filePath := "./" + fileName + ".json"
	//1、打开文件
	file, err := os.OpenFile("2020-05-18.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755) // For read access.
	//2、关闭文件
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	// 3、写入数据
	// data := []byte{115, 111, 109, 101, 10}
	// data := []byte(bodyText)

	file.Write(bodyText)

	// filePath = "bufio.txt"
	// filePath = "file.txt"
	//4、读取文件
	// res, err := ioutil.ReadFile(filePath)
	// CheckError(err)
	// fmt.Println(string(res))
}

//date_desc

func Satilik() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.sahibinden.com/sahibinden-ral/rest/classifieds/search?a811=62364&a812=40603&price_max=1500000&category=16633&price_min=500000&price_currency=1&a812=40728&a812=40602&a812=40604&a812=40605&a812=40606&a811=40599&a811=40594&a811=40597&a811=40600&a811=40595&a811=62365&a811=40598&a811=40593&a811=40601&a811=40596&address_quarter=52036&address_quarter=51395&address_quarter=51399&address_quarter=23114&address_quarter=51400&address_quarter=23097&m%3AsearchMeta=true&m%3AincludeCategoryTree=false&m%3AincludeMatchingFavoriteClassifieds=true&sorting=price_desc&m%3AsmartTextSearch=true&pagingOffset=0&pagingSize=200&language=tr", nil)
	CheckError(err)
	req.Header.Set("Host", "api.sahibinden.com")
	req.Header.Set("X-J3PopQvX-d", "o_0")
	req.Header.Set("X-J3PopQvX-c", "A2Fkm5FyAQAAo2gFqhedkwyFgGanWr_NMNRwqY5NYEGvXK70fo7IVDYUTedaAbw5PFKfQDXkMVIO8YLqosIO8Q==")
	req.Header.Set("X-NewRelic-ID", "VQACUF9bCBAGXVNbAQQOUVc=")
	req.Header.Set("x-client-profile", "Generic_v2.1")
	req.Header.Set("x-timestamp", "1591583608346")
	req.Header.Set("X-J3PopQvX-b", "rtmy5d")
	req.Header.Set("X-J3PopQvX-a", "4m-IJQ0qnc8=Q6V2Mdw8oBtAnvUWNsz6CI7JyNLPCWjMZAUbYu_zMJi0sAi-uLvpo2eDjDvW5lcBu723pQN1E=iYLjoOKR5Ktyra9NdHRAf_3yIQNdPSOy9ktmfiNckE_AMQQnqjW571EF3ycmHQLLdSf7yiVJi4IMGcWIZ1_0zjqMV12oiv=IWHBvdieuiv7nC3G=cOawTu_WnCLLb0eBip9T8zF18Cf8R1tcS-gqGmq9Z23pwgJCUcFRZqJ7OVVvvQvOsAYMbVP1o1JuUW3l4zOE-rbMQfg7gVRzhYuHCWBTNM6LMRxx=bHiZSaNtcUn8BU3fpsNnVAMAs5RzEWEH-MZkawK6ErTx52o52xaQDK6EA3HKJJhWyDvJcum-vmVukk6=s3reMm9NDUI0OhZcLxqscL8Hklq9_PD9fjcd_JUPPE3wLrYJxIKOOcaChUtmm48F9gUOxDjAK6uji=B5fRlefUfzyrYQY8ua9A73ShJ9cqO-SE4RWfwz4SnV2L5sIclHzCJiqkLNVvEG6w=qMMYtSgiJbrgcPjcVZB4-Ej0FurPkMqdlNN6I=36UfI45TgrR6KtmYZUdwd8zQ_ZdJijPiedmmGg-yGe1=DNd8LsQoqETlSoCAJM-eS8Be16kR_tz1mUJ2Y8Uhmw8IU_UFAM1tODw2ct=a4=RZz3I7MmH=5HbWGS-btxoPecxmBOdB8hwtbKdx-5sPwazDuAAKxSpsI=EsNvPsWziVcP9UWK8k2KYW=A4auNTuyCYIN1VxKq4OvSUMQzVepSYtN=J7vZNuYLgA1z7pG06L-_dPTE-uoq0cv6S3=k4-5PVK9qkfWu7q1kr1TLlSWZ0O6WmRP9z1lVVIG7EqVygC9qWHOTpcpOAZz=ZF4bPrMpNE6thuAuI2F8ivFv3wEORNBggL=RMDnR9zBGmMg2_bZ9W4pFEJqqnKhCSPHEj=wrrTyHDhtP2ycvh_hlUa=Pei82Vm=PjurO1SQEL_J0Ez3LDWDdlEC1-3zxbmr61En6gmskJmHMjhcUN_P=0hPfA4YQiUybyhF0_rZ10ejjxVNC_foOhbZVw4CJcd9lDYaFzALvO6hdjEeUmb_IITl3HrYx-gvIH3qguWPtkL9iLCsfE8-tohFTVt00f=dCYsdpsb=CHTAu74=1_IZYqMeUnK5Zhzqp_M8j5hmlad6cfWDk-jvAk72kj1cU=smwQ6Wrlgn9H_7zI27L1PJxd5gmV87-AOy-mK=Hrs4GnW6PkZs5-J8-I9oSc-DKnOaueIxQzlnwehRIZ0w_l=lf=bnZjiqWtIQWPism-bph9jptVJU5_a9DTu9gQmGmATTpqpB8abHCbsOaIjnLvB3SuLZK1hNawOooTS5tw5R7H9e9-emfhonde3AuKKgpeQBi4B6YrZ63HV0r=ZvTZYVEwKKda-LLqlCRx8opT-SP0scgAyzUPxyMBbDK0UQMt4z2oahZrlCQCRACnkVrHKmq_5CZSIGQ-WDGZnuMwooIwxodculNsaxW2GRa8avJp3fsItbLv_nwCk4rqVyC")
	req.Header.Set("User-Agent", "Sahibinden-IOS/4.7.0(2020/05/12,iPhone 6,iOS12.4.5)v3")
	req.Header.Set("x-api-hash", "E4E9C3C9AC90A23CA562B2C10870B68843C1BE97")
	req.Header.Set("X-Device-Id", "8CD86B9902A14495B9880120076FE32C")
	req.Header.Set("X-J3PopQvX-z", "p")
	req.Header.Set("x-bundle-id", "com.sahibinden.sahibinden")
	req.Header.Set("x-app-version", "4.7.0")
	req.Header.Set("x-language", "tr")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("X-J3PopQvX-f", "AyHWm5FyAQAA4Gu2p_tGO-sJuIioJO6XMYg5DjRdCdU-ZdCBwg8Sk0Ub-U-OAXUeSnmfQDXk7xUO8YLqosIO8Q==")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-J3PopQvX-e", "a;TqM2AKiYTV-m1BjiKKCmTMMYDe-tgt3qxohMDvQAvi-GVscd9E3yZTSuGEk6JOhNqP5IDEhm8TE7Tseoretuv9Z8JuJuAxwzz1URyO2FSFqXJofMj23VF7WHZ0eJ4fNuHwxKN_4Qm9wsARznpzTn6kAxgIOqRukD5-_3ASQaBkU6glk34PAidqUkyjWRAfdNJpF3OebFpT7OZlhbnccInkF4eKx_vn5xMG8uY1rNQPilt4XPUzK85AyILthF8sAaI3G1F9gfrMDzIItfu-ZipkYG7eVeG0x7jef6Dp1369PL9h3UxL_zB1QArLLHeaCOOs3qMUv6jSBKJ50S5Vd9pVFDP2GUExsuQ7qQnxXuGycGvWvkkvskX7vt88H0VPWzqH-i5XcKwgFBzKrE9MOeuietsLpW1-u44A0JNU589RIgxzISSr1V6LazTluEk3grPBtnQ7aHx7wbg5KpV9B9LlB5j4DeOP99rW5v26nqBIHdXmhUaRlLZkX8HGxn_Ryq73JSjIDSvWfIUvMxfH0OLVg6d9JP9sQDG0jrVKVwCkEQyS5QLSbVYJId9EMDcNtnyRAVrajnd4TCZ6m3i6Xj30-IhBUicxlgn5uz8z2W4_1MY0L4KfKlEWY1NGBcxWLYlB8dvBqETIomv6R62H2DFjl7EV34gp2eaESSfWcW1YAq8zeGtfSe2mvx6j6OSoxhpjMusNBZtqCEDBsV6h7JiYjeOBd-jvTIdCKwjv4oI50Aqo8UEPuQ88u4jDDv83sogF842VO8vL49CTq9vUTIQVE7rIj6Z_Ym;h96PhBA1aVcmtRGD2P5UogR9O8jad0jiFaYwKNY_HC8=")
	req.Header.Set("x-api-key", "f51c6d9290ea9522aceac00f3b87a0002695d38a")
	req.Header.Set("idfa", "66642314-815E-4E43-BCE8-D1D8652315B6")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	res := models.ResponseRes{}
	CheckError(err)
	err = json.Unmarshal([]byte(bodyText), &res)
	CheckError(err)

	for _, value := range res.Response.Classifieds {
		id, _ := strconv.Atoi(value.ID)
		evBase := models.Satikik{
			EvId:                   id,
			ShortName:              value.Shortname,
			CategoryID:             value.CategoryID,
			Title:                  value.Title,
			Status:                 value.Status,
			Live:                   value.Live,
			Price:                  value.Price,
			Currency:               value.Currency,
			FormattedPrice:         value.FormattedPrice,
			ClassifiedDate:         value.ClassifiedDate,
			ExpiryDate:             value.ExpiryDate,
			ImageURL:               value.ImageURL,
			ImageURLLargeThumbnail: value.ImageURLLargeThumbnail,
			ImageMainURL:           value.ImageMainURL,
			BanyoSayisi:            value.Attributes.A22,
			Fullm2:                 value.Attributes.A24,
			Netm2:                  value.Attributes.A107889,
			OdaSayi:                value.Attributes.A20,
			BinaYasi:               value.Attributes.A812,
			BuluduguKat:            value.Attributes.A811,
			KatSayi:                value.Attributes.A810,
			Isitma:                 value.Attributes.A23,
			Balkon:                 value.Attributes.A106960,
			Aidat:                  value.Attributes.A104239,
			CreatedAt:              time.Now(),
		}
		db := drivers.GetDB()
		db.NewRecord(evBase)
		db.Create(&evBase)

		fiyat := models.FiyatForSatlik{
			EvId:      id,
			Fiyat:     value.Price,
			CreatedAt: time.Now(),
		}
		db.NewRecord(fiyat)
		db.Create(&fiyat)
	}
}

func Kiralik() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.sahibinden.com/sahibinden-ral/rest/classifieds/search?a811=40594&a812=40603&price_max=3500&category=16624&price_min=1800&price_currency=1&a812=40728&a812=40602&a812=40604&a812=40605&a812=40606&a812=40607&a811=40599&a811=62364&a811=40597&a811=40600&a811=40595&a811=62365&a811=40598&a811=40593&a811=40601&a811=40596&address_quarter=51395&address_quarter=23114&address_quarter=52036&address_quarter=51400&address_quarter=23097&address_quarter=51399&m%3AsearchMeta=true&m%3AincludeCategoryTree=false&m%3AincludeMatchingFavoriteClassifieds=true&sorting=price_desc&m%3AsmartTextSearch=true&pagingOffset=0&pagingSize=200&language=tr", nil)
	CheckError(err)
	req.Header.Set("Host", "api.sahibinden.com")
	req.Header.Set("X-J3PopQvX-d", "o_0")
	req.Header.Set("X-J3PopQvX-c", "A2Fkm5FyAQAAo2gFqhedkwyFgGanWr_NMNRwqY5NYEGvXK70fo7IVDYUTedaAbw5PFKfQDXkMVIO8YLqosIO8Q==")
	req.Header.Set("X-NewRelic-ID", "VQACUF9bCBAGXVNbAQQOUVc=")
	req.Header.Set("x-client-profile", "Generic_v2.1")
	req.Header.Set("x-timestamp", "1591584701568")
	req.Header.Set("X-J3PopQvX-b", "rtmy5d")
	req.Header.Set("X-J3PopQvX-a", "4m-IJQ0qnc8=Q6V2Mdw8oBtAnvUWNsz6CI7JyNLPCWjMZAUbYu_zMJi0sAi-uLvpo2eDjDvW5lcBu723pQN1E=iYLjoOKR5Ktyra9NdHRAf_3yIQNdPSOy9ktmfiNckE_AMQQnqjW571EF3ycmHQLLdSf7yiVJi4IMGcWIZ1_0zjqMV12oiv=IWHBvdieuiv7nC3G=cOawTu_WnCLLb0eBip9T8zF18Cf8R1tcS-gqGmq9Z23pwgJCUcFRZqJ7OVVvvQvOsAYMbVP1o1JuUW3l4zOE-rbMQfg7gVRzhYuHCWBTNM6LMRxx=bHiZSaNtcUn8BU3fpsNnVAMAs5RzEWEH-MZkawK6ErTx52o52xaQDK6EA3HKJJhWyDvJcum-vmVukk6=s3reMm9NDUI0OhZcLxqscL8Hklq9_PD9fjcd_JUPPE3wLrYJxIKOOcaChUtmm48F9gUOxDjAK6uji=B5fRlefUfzyrYQY8ua9A73ShJ9cqO-SE4RWfwz4SnV2L5sIclHzCJiqkLNVvEG6w=qMMYtSgiJbrgcPjcVZB4-Ej0FurPkMqdlNN6I=36UfI45TgrR6KtmYZUdwd8zQ_ZdJijPiedmmGg-yGe1=DNd8LsQoqETlSoCAJM-eS8Be16kR_tz1mUJ2Y8Uhmw8IU_UFAM1tODw2ct=a4=RZz3I7MmH=5HbWGS-btxoPecxmBOdB8hwtbKdx-5sPwazDuAAKxSpsI=EsNvPsWziVcP9UWK8k2KYW=A4auNTuyCYIN1VxKq4OvSUMQzVepSYtN=J7vZNuYLgA1z7pG06L-_dPTE-uoq0cv6S3=k4-5PVK9qkfWu7q1kr1TLlSWZ0O6WmRP9z1lVVIG7EqVygC9qWHOTpcpOAZz=ZF4bPrMpNE6thuAuI2F8ivFv3wEORNBggL=RMDnR9zBGmMg2_bZ9W4pFEJqqnKhCSPHEj=wrrTyHDhtP2ycvh_hlUa=Pei82Vm=PjurO1SQEL_J0Ez3LDWDdlEC1-3zxbmr61En6gmskJmHMjhcUN_P=0hPfA4YQiUybyhF0_rZ10ejjxVNC_foOhbZVw4CJcd9lDYaFzALvO6hdjEeUmb_IITl3HrYx-gvIH3qguWPtkL9iLCsfE8-tohFTVt00f=dCYsdpsb=CHTAu74=1_IZYqMeUnK5Zhzqp_M8j5hmlad6cfWDk-jvAk72kj1cU=smwQ6Wrlgn9H_7zI27L1PJxd5gmV87-AOy-mK=Hrs4GnW6PkZs5-J8-I9oSc-DKnOaueIxQzlnwehRIZ0w_l=lf=bnZjiqWtIQWPism-bph9jptVJU5_a9DTu9gQmGmATTpqpB8abHCbsOaIjnLvB3SuLZK1hNawOooTS5tw5R7H9e9-emfhonde3AuKKgpeQBi4B6YrZ63HV0r=ZvTZYVEwKKda-LLqlCRx8opT-SP0scgAyzUPxyMBbDK0UQMt4z2oahZrlCQCRACnkVrHKmq_5CZSIGQ-WDGZnuMwooIwxodculNsaxW2GRa8avJp3fsItbLv_nwCk4rqVyC")
	req.Header.Set("User-Agent", "Sahibinden-IOS/4.7.0(2020/05/12,iPhone 6,iOS12.4.5)v3")
	req.Header.Set("x-api-hash", "E785A8C2D52A3E4FEBB7BDDE3DA285291A1A6E08")
	req.Header.Set("X-Device-Id", "8CD86B9902A14495B9880120076FE32C")
	req.Header.Set("X-J3PopQvX-z", "p")
	req.Header.Set("x-bundle-id", "com.sahibinden.sahibinden")
	req.Header.Set("x-app-version", "4.7.0")
	req.Header.Set("x-language", "tr")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("X-J3PopQvX-f", "AyHWm5FyAQAA4Gu2p_tGO-sJuIioJO6XMYg5DjRdCdU-ZdCBwg8Sk0Ub-U-OAXUeSnmfQDXk7xUO8YLqosIO8Q==")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-J3PopQvX-e", "a;kvFzjuO-Nc-_6sRpnh_H5Y11e8zw4GthQL78Pc-2XGFxK-jE90vJdTWNtWSYdS0Eo9KtPqy2UKnRhC6LrrVqAEZvnz6UmlaZzo96T4UpdSDTi_4y6piWmknJn0XGjfUFslk5RcaDmKl9dKN64o0jsyKdhePawjw67ZF3JAfygEJkYTsFeV90adxY6oKoTP2zrq_fWvbZv9163BFvj_HZv3KjHyH2zPaMebdrc--6UDKyzxUSiePRUFqmP38XvHZRxFQz1H1FfSfVhb0wMTjav4SYBItEijBj9ymrFJYSQy4WlWBN4fCLtkL1KdNYSGzH_B0hdVhXL-UsWMXbPW56_vFC0XrUOp2J6o4vRle2SUm9E7tMmFwaiusQz1vBmZP2hl_aifEwmkiF8I3FVyaelWyWS91aGOUrxDRpY-0CAFpMA7BYSsAee4Eutm4O59rH4I54z2Naeh058LNOMOhGgvAS8TaDXLMU41tX2JV7Et06RNxQigDQ8qmmjfdoEyTnhNd183nDD0J4n7hHasrc3mZuF3lFw3GMt0XAtqkzOtRROIEIWQO8D5Vk90864IC73jEtbgJVi3SW-mAwT4Ol6PRziONCVDVkSp4moTJ9KAxp87LImoTOp8yu8_yuJnL85W2-nvDuZngE-lH9syg1MJ_1n6E6qBFyLjSf1mL-9Z_fgneHZz1pnqF0VN2ehC0KhK0ZSfGHzhXKXT8BK0765ow0Va8QH09AdRj_SFEDC3OLkJSsGVetuZGnC8FZnwp6TAE6ElvrzWxTv4KLBHS9mdyy4rigjA==;yLAWG3VGTCuztVW4FTB1Vfyj469NXqunoPCt6aZTjy8=")
	req.Header.Set("x-api-key", "f51c6d9290ea9522aceac00f3b87a0002695d38a")
	req.Header.Set("idfa", "66642314-815E-4E43-BCE8-D1D8652315B6")
	resp, err := client.Do(req)
	CheckError(err)
	bodyText, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	fmt.Printf("%s\n", bodyText)

	res := models.ResponseRes{}
	CheckError(err)
	err = json.Unmarshal([]byte(bodyText), &res)
	CheckError(err)

	for _, value := range res.Response.Classifieds {
		id, _ := strconv.Atoi(value.ID)
		evBase := models.Kiralik{
			EvId:                   id,
			ShortName:              value.Shortname,
			CategoryID:             value.CategoryID,
			Title:                  value.Title,
			Status:                 value.Status,
			Live:                   value.Live,
			Price:                  value.Price,
			Currency:               value.Currency,
			FormattedPrice:         value.FormattedPrice,
			ClassifiedDate:         value.ClassifiedDate,
			ExpiryDate:             value.ExpiryDate,
			ImageURL:               value.ImageURL,
			ImageURLLargeThumbnail: value.ImageURLLargeThumbnail,
			ImageMainURL:           value.ImageMainURL,
			BanyoSayisi:            value.Attributes.A22,
			Fullm2:                 value.Attributes.A24,
			Netm2:                  value.Attributes.A107889,
			OdaSayi:                value.Attributes.A20,
			BinaYasi:               value.Attributes.A812,
			BuluduguKat:            value.Attributes.A811,
			KatSayi:                value.Attributes.A810,
			Isitma:                 value.Attributes.A23,
			Balkon:                 value.Attributes.A106960,
			Aidat:                  value.Attributes.A104239,
			CreatedAt:              time.Now(),
		}
		db := drivers.GetDB()
		db.NewRecord(evBase)
		db.Create(&evBase)

		fiyat := models.FiyatForKiralik{
			EvId:      id,
			Fiyat:     value.Price,
			CreatedAt: time.Now(),
		}
		db.NewRecord(fiyat)
		db.Create(&fiyat)

	}

}
