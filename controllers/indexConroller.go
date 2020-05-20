package controllers

import (
	"../drivers"
	"../logic"
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
