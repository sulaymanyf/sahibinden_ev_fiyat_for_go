package controllers

import (
	"../drivers"
	"../logic"
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/kkdai/youtube"
	"log"
	"net/http"
	"os/user"
)

func IndexHome(ctx *gin.Context) {

	//// query string
	//queryVal1 := ctx.Query("val1")
	//queryVal2 := ctx.DefaultQuery("val2", "val2_default")
	//
	//// post form data
	//formVal3 := ctx.PostForm("val3")
	//formVal4 := ctx.DefaultPostForm("val4", "val4_default")
	//
	//// path info
	//pathVal5 := ctx.Param("val5")
	dair := models.Daire{}
	sql := "select * from istanbul  group by ev_id order by price desc"
	rows, err := drivers.MysqlDb.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	evler := make([]models.Daire, 0)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dair.ID, &dair.EvID, &dair.ClassifiedDate, &dair.Comparable, &dair.Currency, &dair.ExpiryDate, &dair.FormattedPrice, &dair.ImageMainURL, &dair.Price, &dair.Shortname, &dair.Title, &dair.CreatTime)
		if err != nil {
			log.Fatal(err)
		}
		evler = append(evler, dair)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	data["tatol"] = len(evler)
	fmt.Println(data["tatol"])
	data["evler"] = evler
	//ctx.String(http.StatusOK, "hello %s %s %s %s %s", queryVal1, queryVal2, formVal3, formVal4, pathVal5)
	ctx.HTML(http.StatusOK, "index/index.html", data)
}

func GetPriceById(ctx *gin.Context) {
	id := ctx.Query("ev_id")
	filatlar := make([]models.Daire, 0)

	dair := models.Daire{}
	sql := "select price ,creat_time from istanbul where ev_id = ?"
	rows, err := drivers.MysqlDb.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&dair.Price, &dair.CreatTime)
		if err != nil {
			log.Fatal(err)
		}
		filatlar = append(filatlar, dair)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	ctx.HTML(http.StatusOK, "index/fiyat.html", filatlar)

}

func Sahibinden(ctx *gin.Context) {

	logic.GetEvinFiyat()

	IndexHome(ctx)

	//ctx.HTML(http.StatusOK, "index/index.html", gin.H{
	//	"msg": "easy gin",
	//})
}

func YoutubeDownLoad(ctx *gin.Context) {
	log.Println(ctx.ClientIP())
	log.Println(ctx.Request.Method == "GET")
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "index/youtube.html", gin.H{
			"msg": "easy gin",
		})
		return
	}
	if ctx.Request.Method == "POST" {
		url := ctx.Request.FormValue("url")
		usr, _ := user.Current()
		currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
		log.Println("download to dir=", currentDir)
		y := NewYoutube(true)
		log.Println("https://www.youtube.com/watch?v=zaakp81L9wU")
		log.Println(url)
		if err := y.DecodeURL(url); err != nil {
			fmt.Println("err:", err)
		}
		log.Println(y.StreamList)
		ctx.HTML(http.StatusOK, "index/youtube.html", y.StreamList[0])

	}

}

func Kiralik(ctx *gin.Context) {

	logic.Kiralik()

}

func Yeni(ctx *gin.Context) {

	ids := logic.GetEvIds("kiralik")
	evler := logic.GetEvinfo(ids)

	ctx.HTML(http.StatusOK, "index/evler.html", evler)

}
