package main

import (
	"encoding/csv"
	"fmt"
	"gys/models"
	"io/ioutil"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func initImport() {
	// 迁移 schema 创建Merchant表的同时还会创建Area表
	db = models.DB
	db.AutoMigrate(&models.Merchant{}, &models.Area{})
	importCSV()

}

func importCSV() {
	f, err := ioutil.ReadFile("datasheet.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	r2 := csv.NewReader(strings.NewReader(string(f)))
	rows, _ := r2.ReadAll()
	for i := 0; i < len(rows); i++ {
		row := rows[i]

		// row 取得了每一行所有列的值，row是[]string
		merchant := models.Merchant{
			Name:                  row[4],
			Mcode:                 row[1],
			Mtype:                 row[7],
			JoinRelation:          row[8],
			CooperateRelation:     row[9],
			Remarks:               row[17],
			IntroductionDate:      str2time(row[18]),
			BeUndertakerDate:      str2time(row[20]),
			BeServiceproviderDate: str2time(row[19]),

			ContractDate:   str2time(row[23]),
			ExpirationDate: str2time(row[24]),
		}

		area := models.Area{
			Region:             row[10],
			Province:           row[11],
			City:               row[12],
			CityClassification: row[13],
			IsKeyCity:          chi2bool(row[14]),
			IsOnPlatform:       chi2bool(row[15]),
			Bd:                 row[16],
		}

		db.FirstOrCreate(&area, models.Area{
			City: row[12],
		})

		db.FirstOrCreate(&merchant, models.Merchant{
			Name: row[4],
		})

		db.Model(&merchant).Association("Areas").Append(&area)
	}

}

// 日期转时间
func str2time(str string) time.Time {

	sSlice := strings.Split(str, "-")
	sFormat := ""
	if len(sSlice) == 3 {
		sFormat = sSlice[0] + "-" + addzero(sSlice[1]) + "-" + addzero(sSlice[2])
	}
	// 为免时区混乱，始终以UTC时间记录，数据库中以带时区的timestamptz记录
	// 或参考此处 https://gyaan.github.io/Working-With-Different-Timezone-In-Golang-And-PostgreSQL/
	theTime, _ := time.Parse(time.RFC3339, sFormat+"T00:00:00+00:00")
	return theTime
}

func addzero(str string) string {
	if len(str) <= 1 {
		return "0" + str
	}
	return str
}

func chi2bool(str string) bool {
	if str == "是" {
		return true
	}
	return false

}
