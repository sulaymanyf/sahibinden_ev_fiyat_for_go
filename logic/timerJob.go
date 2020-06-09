package logic

import (
	"../drivers"
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SatilikJob() {
	var start int
	var end int
	db := drivers.GetDB()
	db.Table("satikiks").Count(&start)
	Satilik()
	db.Table("satikiks").Count(&end)

	if start != end {
		SendDingMsg("yeni  " + string(end-start) + " Satilik evler gelmis")
	}

}

func GetEvIds(name string) (ids []int) {

	db := drivers.GetDB()
	format := time.Now().Format("2006-01-02")
	if name == "kiralik" {
		var Kiraliks []models.Kiralik

		db.Table("kiraliks").Where("created_at >= ?", format).Select("ev_id").Find(&Kiraliks)

		for _, v := range Kiraliks {
			ids = append(ids, v.EvId)
		}
		log.Print(ids)
		return
		//return find
	} else {
		var satliks []models.Satikik
		db.Table("satikiks").Where("created_at >= ?", format).Select("ev_id").Find(&satliks)

		for _, v := range satliks {
			ids = append(ids, v.EvId)
		}
		log.Print(ids)
		return

	}

}

func KiralikJob() {
	var start int
	var end int
	db := drivers.GetDB()
	db.Table("kiraliks").Count(&start)
	Kiralik()
	db.Table("kiraliks").Count(&end)

	if start != end {
		SendDingMsg("yeni " + string(end-start) + " Kiralik evler gelmis")
	}

}

func SendEvInfo(msg string, evler []models.TekEvRes) {
	SendDingMsg(msg)
	for _, ev := range evler {
		log.Print(ev.Response)
		//请求地址模板
		webHook := `https://oapi.dingtalk.com/robot/send?access_token=1174ae5d4bc3460aaa28d8afd80c1e492d6566929784950b76285242b53e0efb`
		content := `{"msgtype": "markdown",
		"markdown": {
				"title": "` + ev.Response.Title + `",
				"text": '` + ev.Response.Description + `',
				"picUrl": "` + ev.Response.Images[0].Normal + `",
				"messageUrl":	"` + ev.Response.URL + `"
            }
	}`

		log.Print(content)
		//创建一个请求
		req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
		if err != nil {
			// handle error
		}

		client := &http.Client{}
		//设置请求头
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		//发送请求
		resp, err := client.Do(req)
		//关闭请求
		defer resp.Body.Close()

		if err != nil {
			// handle error
		}
	}
}

func GetEvinfo(ids []int) (evler []models.Response) {
	client := &http.Client{}
	for _, id := range ids {
		id := strconv.Itoa(id)
		log.Print(id)
		req, err := http.NewRequest("GET", "https://api.sahibinden.com/sahibinden-ral/rest/classifieds/"+id+"?language=tr", nil)
		log.Print("===NewRequest==")
		CheckError(err)
		req.Header.Set("Host", "api.sahibinden.com")
		req.Header.Set("X-J3PopQvX-d", "o_0")
		req.Header.Set("X-J3PopQvX-c", "A-QFjJZyAQAAPE6mbdqyB-X0i7w-YXjI0yissi6xS21XJQSF8jd68D7lxQdEAV-t-KyfQDXkMVIO8YLqosIO8Q==")
		req.Header.Set("X-NewRelic-ID", "VQACUF9bCBAGXVNbAQQOUVc=")
		req.Header.Set("x-client-profile", "Generic_v2.1")
		req.Header.Set("x-timestamp", "1591676258471")
		req.Header.Set("X-J3PopQvX-b", "qwbw2x")
		req.Header.Set("X-J3PopQvX-a", "dsiWR5Z98ys_xoUgC3eWGH0dlyiNBDC6tS3l6CGCf1Y8q65PPV2ohL-uqmyoBVf44GuPWojrk_IY0J9MCf5TYN0oN-e_7G-YVJXL-wT1Q237Wl6amiOD_eIR-nYP1WrhnTsDaHfs657lBK_-IeUIAn_UffYNlIA73vR03IOJpLIF=SpBVP-kTJzQQk6vN-ce56qWXHL_vIOFv15qA0euZ6ZOcWEIpH46w=QwJ7sGX3mFc97uIiyorcMfF6P4SX05GQkXnUu=KRLEu=oDC9315wHAkDdXBReFTs4eL8t4SO6xYDLAGiySXSZZh1evPJfcvhYzEn4HLB6L6xXkKPWV7VY4IMsnWQiXxkGvZlDmzUlW_4OGjwYU4H_8Fwx0S8FuBlSYDs6qgkBXLqnUhza=GG4szCOLyp3BdElR52KE30LM9Mp_YsXYtCuawRH=VIQHgSPrMHKRZ1qWN6AUTwu2FuSLMdU_mcN7KFmLzyDyMWE8LBeiQJTYnWHXk9HYgBcTFgIKCY3OnjFKnaovQvzh3iMLXKvT05CJj9HA2o_rwwr_CXRLNeeBSmx5UCl4vsL1paQhw6thWhhXICYC4aWdmLzFx2NSe5ZqjG83B8YZ=2zpPTmrN9L9sdE4T-DUoIeYiUOMUrv-pOFlYzq6=rAPhB2l4yCrmcS1RcX-VzCji-P58qJFGRIcaHQhlYV3cJ8Q9jGFks3TD6BHyXtkX6kY3ny=TT-qJOdap0oLxmnT5anrhFtU2-ul8qkLw6Hc=EDvpKt4WQDRKHRBzZvEeX1N=pTx-LdOUr5j7GQ7Pt7vOgOUxXoEg7UDRAA9YxSQMoIP7ZRGl49qlQ==jA=XV-Q5uCpcmotxoW2hZAKX0lVUdRKmFEYlgovUTfCPpjI_uf7MmlCOZ7JLgg9UKfDEm5grNv40YnF1vAr__EPfGyF8RB7=KCOO_OFD0warU-efVUyRRNN=448Tr7_LZfL3yUWJBUKUaT3uiXvBSdyY2q_iMWCKa8uqY0QIjfXmW4umZ1XYhEHKC5y8h=MP-wnIH6GYZm2FdVWkmOYl5Fc6huYnf81R5KFif=lk6WUGOL0MWYqRQYgdaypLB8HO77QkgI2pv0o9L56jjQV4cPnJie=MsfmzwxmIxCRepDBZ0XXKaw2eD8EuQuvYQ2-l_ZQCr2KAOm3RWBvSF0cIH0hwid2Gj0UfX6OSitKhYOFVdqojlFez==TlfoNgO=JODf=ywyaZtupGFkOQGT9N0dMj02z1FNrHFPInF=H5ffz2L133i2iwSU3S6=OsacDEUcOTZvZXJlos7x1pkCxkPL5yAJJ6tUAkSWaYKIct-f4mvunpfOWxLLOV3GfxcgNEORx5LfpJ-2svxo5axCkVInRRuxGyRMcaX1=USHQ6EUxNo6w56-Js83oODMoynrynLBLo8QtKYR_UYS7OPx9pgsUSlILftiIELy45SUvTwCSD7YxDUqaMWn6E0q_hOO=7D-_uSfRGQ5ITpSGhgglqkgkeQSz-A0PCjcPkEiiJ39vB-efY6U29Lj2Nh3cvht")
		req.Header.Set("User-Agent", "Sahibinden-IOS/4.7.0(2020/05/12,iPhone 6,iOS12.4.5)v3")
		req.Header.Set("x-api-hash", "915A679DC9CD4208C0E1BB5862939A7D45A83EB8")
		req.Header.Set("X-Device-Id", "8CD86B9902A14495B9880120076FE32C")
		req.Header.Set("X-J3PopQvX-z", "p")
		req.Header.Set("x-bundle-id", "com.sahibinden.sahibinden")
		req.Header.Set("x-app-version", "4.7.0")
		req.Header.Set("x-language", "tr")
		req.Header.Set("Accept-Language", "zh-cn")
		req.Header.Set("X-J3PopQvX-f", "A2tFjJZyAQAATNxglDhxdc9uj7k0dcEF8K4wXrXjs9NqxF08kItIRug2pJE3AXUeSnmfQDXk7xUO8YLqosIO8Q==")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-J3PopQvX-e", "a;3UlMlU6oJEmOrcVwcIJyKIsapMPg8uo7gYSP8w0sXFYAYbP_XwxZZ_xphidW_DxANXNWVmgI8OH8TZQ26TMPO1Kl7sSa5mcOTcCkCousSeXglmQU8gJoolsLvv6I31v30yVxCKEvocAN3O2_yS4PHraCYDF9JJiwUWyo9N-Sur5tylDQnWt82Kzd9jJqtBOsPqE-a5H-rtUr4khZhVJXmI3FHdbsc5H2DgZGY8nLgSOTFY97iuiEAXYmIlrPwFSD4Gkp3uQvK4Fv5BydV_pqzAl-KxIMuPX-TcVYTHtP5YndJm5nV58UcqN24whrtoGDtVDwPfcylYB6iN75Ql5qbaBPsAtM4GbYKeZe-ivWIsmLhTP7BULzEtpjkuay3kNOf79TbWJ5rdvsB00__g4QPLrUfz1FWeuqZzXBkJfPp2qdce2DG-Hr-47kYmIQFCiiMcf27s61BSYrd-IhB18KjytdDmcARCwVySi8stJIp9KIVeFYQdh-xIW_FmgqctcQf073s3iTa2rDUrKGYJX_8SHaH0A1fs4T6In7g3h0dIYaUi1Be0ITRejObSEmn2UFtyhj5Oe61Dv939FyV2zedyyq2USwuXNzkTSyav3iR-0MqHhJLJlYUC8so_z_7yhKu43ozU04m8DfF8W551l0tcKHxQMxDaOV9EFx0s1urkx5WSglkOMcsGACulYwKBStIPNqapI4hSLjn2mnLpFHxc96WwsylYyBUJwOjAB9SXCW6ZUaPBvVAgd0eXojv-np9F25UJnqAZk19qfOG8B5H7TJ-PR5m_VJ;CqL-7kJC1W-7-H48sMycInIH1UVHaZFWjfTaVCbJog0=")
		req.Header.Set("x-api-key", "f51c6d9290ea9522aceac00f3b87a0002695d38a")
		req.Header.Set("idfa", "66642314-815E-4E43-BCE8-D1D8652315B6")
		resp, err := client.Do(req)
		CheckError(err)
		bodyText, err := ioutil.ReadAll(resp.Body)
		CheckError(err)
		fmt.Printf("%s\n", bodyText)
		res := models.TekEvRes{}
		CheckError(err)
		err = json.Unmarshal([]byte(bodyText), &res)
		CheckError(err)
		log.Print(res.Success == false)
		if res.Success == false {
			return
		}
		evler = append(evler, res.Response)
	}
	log.Print("---------evler-------------")
	log.Print(evler)
	return evler
}

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=1174ae5d4bc3460aaa28d8afd80c1e492d6566929784950b76285242b53e0efb`
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}

func SendMacdown() {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=1174ae5d4bc3460aaa28d8afd80c1e492d6566929784950b76285242b53e0efb`
	content := `{
     "msgtype": "markdown",
     "markdown": {
         "title":"杭州天气 ev",
         "text": "Iskeleye ve maramaraya 8-10 dakika yurume mesafesinde

Uskudar merkez Mimar sinan mahallesinde

Deprem sonrasi 2007 yilinda yapilmistir.&#xa0;

3 katli binanin 2. katindadir ve 2 cepheledir.&#xa0;

Kiraci 12 temmuzda evden ayrilacaktir bu sebeple bu tarihten itibaren kiralanabilir.&#xa0;

Binanaya dis cephe izolasyonu yapilmistir. Balkona Pimapen vasistas yapilmistir. \n"
     },
      "at": {
          "atMobiles": [
              "150XXXXXXXX"
          ],
          "isAtAll": false
      }
 }`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}
