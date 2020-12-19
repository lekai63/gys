package models

import (
	"time"

	"gorm.io/gorm"
)

// Merchant 商家表
type Merchant struct {
	gorm.Model
	Name                  string    // 商家名称
	Mcode                 string    // 商家编码
	Mtype                 string    // 商家类型
	JoinRelation          string    // 加盟关系（初始加盟商与否）
	CooperateRelation     string    // 是否合作(合作，不合作，待确定)
	Remarks               string    // 备注
	IntroductionDate      time.Time // 引入日期
	BeUndertakerDate      time.Time // 转承接商日期
	BeServiceproviderDate time.Time // 转服务商日期

	// many2many 模型
	Areas []Area `gorm:"many2many:merchant_areas;"` // 覆盖的城市

	ContractDate   time.Time // 签约日期
	ExpirationDate time.Time // 签约失效日期
}

// Area 服务区域表
type Area struct {
	gorm.Model
	Region             string // 大区
	Province           string // 省
	City               string // 市
	CityClassification string // 城市分类
	IsKeyCity          bool   // 是否重点城市：true --> 是 ，false --> 否
	IsOnPlatform       bool   // 是否平台上线城市：true --> 是 ，false --> 否
	Bd                 string // 区域拓展人名字
	// 反向引用
	Merchants []*Merchant `gorm:"many2many:merchant_areas;"`
}

// 派单情况
// type Dispatch struct {

// }
