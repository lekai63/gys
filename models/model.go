package models

import (
	"time"

	"gorm.io/gorm"
)

// Merchant 商家表
type Merchant struct {
	gorm.Model
	Name              string // 商家名称
	Mcode             string // 商家编码
	Mtype             string // 商家类型
	JoinRelation      string // 加盟关系（初始加盟商与否）
	CooperateRelation string // 是否合作(合作，不合作，待确定)
	Remarks           string // 备注

	IntroductionDate      time.Time // 引入日期
	BeUndertakerDate      time.Time // 转承接商日期
	BeServiceproviderDate time.Time // 转服务商日期
	ContractDate          time.Time // 签约日期
	ExpirationDate        time.Time // 签约失效日期

	// 非业务相关的商家一般性字段
	GeneralInfo

	// 关联关系
	// 迁移数据表时，先将以下belongs to 模型的关联关系注释，用以先创建many to many表
	// 之后再放开 belongs to 模型关联（注册地城市）

	// many2many 模型
	Areas []Area `gorm:"many2many:merchant_areas;"` // 覆盖的城市
	// 注册地所在城市
	RegisteredCityID int
	RegisteredCity   RegisteredCity
}

// GeneralInfo 非业务相关的商家一般性字段
type GeneralInfo struct {
	MainCategory      string    // 主要品类
	LegalPerson       string    // 法人
	EnterpriseType    string    // 企业类型
	RegisteredCapital uint      // 注册资本，精确到元
	EstablishDate     time.Time // 成立日期
	OperateStart      time.Time // 营业执照-营业起始日
	OperateEnd        time.Time // 营业执照-营业终止日，长期记为 9999-12-31
	BusinessScope     string    // 经营范围
	RegisteredAddress string    //注册地址
	OperateStatus     string    // 经营状态
	BillLicense       string    //代账许可
	BillLicenseStart  time.Time // 代账许可有效期-起始日
	BillLicenseEnd    time.Time // 代账许可有效期-终止日，长期或无固定期限记为 9999-12-31
	BankName          string    //开户银行
	BankAccount       string    // 银行账号
	TaxNumber         string    //税号
	ActualAddress     string    //实际地址
	Contact           string    //联系人
	Email             string
}

// RegisteredCity 注册地所属城市
type RegisteredCity struct {
	gorm.Model
	Region             string // 大区
	Province           string // 省
	CityName           string // 市
	CityClassification string // 城市分类
	IsKeyCity          bool   // 是否重点城市：true --> 是 ，false --> 否
	IsOnPlatform       bool   // 是否平台上线城市：true --> 是 ，false --> 否
	Bd                 string // 区域拓展人名字
}

// Area 服务覆盖的区域或城市
type Area struct {
	gorm.Model
	City      string
	Merchants []*Merchant `gorm:"many2many:merchant_areas;"`
}

// 派单情况
// type Dispatch struct {

// }
