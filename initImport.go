package main

import (
	"encoding/csv"
	"fmt"
	"gys/models"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func initImport() {
	// 迁移 schema 创建Merchant表的同时还会创建Area表
	db = models.DB
	db.AutoMigrate(&models.Merchant{}, &models.RegisteredCity{}, &models.Area{})
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
	for i := 2; i < len(rows); i++ {
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

			GeneralInfo: models.GeneralInfo{
				MainCategory:      row[22],
				LegalPerson:       row[25],
				EnterpriseType:    row[26],
				RegisteredCapital: wanStr2YuanUint(row[27]),
				EstablishDate:     str2time(row[28]),
				OperateStart:      str2time(row[30]),
				OperateEnd:        str2time(row[31]), // 营业执照-营业终止日，长期记为 9999-12-31
				BusinessScope:     row[32],           // 经营范围
				RegisteredAddress: row[33],           //注册地址
				OperateStatus:     row[34],           // 经营状态
				BillLicense:       row[35],           //代账许可
				BillLicenseStart:  str2time(row[36]), // 代账许可有效期-起始日
				BillLicenseEnd:    str2time(row[37]), // 代账许可有效期-终止日，长期或无固定期限记为 9999-12-31
				BankName:          row[38],           //开户银行
				BankAccount:       row[39],           // 银行账号
				TaxNumber:         row[40],           //税号
				ActualAddress:     row[43],           //实际地址
				Contact:           row[44],           //联系人
				Email:             row[45],
			},
		}

		registeredCity := models.RegisteredCity{
			Region:             row[10],
			Province:           row[11],
			CityName:           row[12],
			CityClassification: row[13],
			IsKeyCity:          chi2bool(row[14]),
			IsOnPlatform:       chi2bool(row[15]),
			Bd:                 row[16],
		}

		area := models.Area{
			City: row[22],
		}

		db.FirstOrCreate(&registeredCity, models.RegisteredCity{
			CityName: row[12],
		})

		db.FirstOrCreate(&area, models.Area{
			City: row[22],
		})

		db.FirstOrCreate(&merchant, models.Merchant{
			Name: row[4],
		})

		// 添加 服务区域关联关系
		db.Model(&merchant).Association("Areas").Append(&area)
		// 服务区域添加后，再把下面这行去掉注释，以添加 注册地所在城市的关联关系
		db.Model(&merchant).Association("RegisteredCity").Append(&registeredCity)

	}

}

// 日期转时间
func str2time(str string) time.Time {

	if strings.ContainsRune(str, '长') || strings.ContainsRune(str, '固') {
		theTime, _ := time.Parse(time.RFC3339, "9999-12-31"+"T00:00:00+00:00")
		return theTime
	}

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

func wanStr2YuanUint(str string) uint {
	index := strings.IndexRune(str, '万')
	var i uint
	if index == -1 {
		temp, _ := strconv.Atoi(str)
		i = uint(temp)
	} else {
		str = string([]rune(str)[:index])
		temp, _ := strconv.Atoi(str)
		i = uint(temp) * 10000
	}

	return i
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
