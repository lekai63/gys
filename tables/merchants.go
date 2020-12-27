package tables

import (
	"gys/models"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMerchantsTable(ctx *context.Context) table.Table {

	merchants := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))
	info := merchants.GetInfo()
	info.HideFilterArea() //默认隐藏筛选框

	info.SetFilterFormLayout(form.LayoutFourCol)
	info.AddField("序号", "id", db.Int8).
		FieldSortable()
	info.AddField("商家名称", "name", db.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}). // 模糊查询
		FieldSortable().
		FieldEditAble()
		/* 	info.AddField("商家编码", "mcode", db.Text).
		FieldSortable().
		FieldEditAble() */
	info.AddField("商家类型", "mtype", db.Text).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "服务商", Text: "服务商"},
			{Value: "承接商", Text: "承接商"},
			{Value: "不合作", Text: "不合作"},
		}).
		FieldSortable().
		FieldEditAble()
	info.AddField("加盟关系", "join_relation", db.Text).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "新引入", Text: "新引入"},
			{Value: "意向加盟商", Text: "意向加盟商"},
			{Value: "加盟商", Text: "加盟商"},
		}).
		FieldSortable().
		FieldEditAble()
	info.AddField("合作关系", "cooperate_relation", db.Text).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "合作", Text: "合作"},
			{Value: "不合作", Text: "不合作"},
			{Value: "待确定", Text: "待确定"},
		}).
		FieldSortable().
		FieldEditAble()
	info.AddField("引入日期", "introduction_date", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("转承接商日期", "be_undertaker_date", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("转服务商日期", "be_serviceprovider_date", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("签约日期", "contract_date", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("失效日期", "expiration_date", db.Timestamptz).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("备注", "remarks", db.Text).
		FieldEditAble()
	info.AddField("服务区域", "custom_serve_areas", db.Text).
		FieldDisplay(func(model types.FieldModel) interface{} {
			orm := models.DB
			var merchant models.Merchant
			var areas []models.Area
			merchant.ID = uint((model.Row["id"].(int64)))
			orm.Model(&merchant).Association("Areas").Find(&areas)
			// fmt.Println("areas:", areas)
			shows := make([]string, 0)
			for _, area := range areas {
				shows = append(shows, area.City)
			}
			return strings.Join(shows, ",")
		})

	info.AddField("签约失效日期", "expiration_date", db.Date).
		FieldFilterable().
		FieldSortable().
		FieldEditAble().
		FieldDisplay(dateDisplay)
	info.AddField("注册资本:元", "registered_capital", db.Int8).
		FieldEditAble().
		FieldHide()
	info.AddField("成立日期", "establish_date", db.Date).
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("营业执照起始日", "operate_start", db.Date).
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("营业执照终止日", "operate_end", db.Date).
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("代账许可起始日", "bill_license_start", db.Date).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("代账许可终止日", "bill_license_end", db.Date).
		FieldFilterable(types.FilterType{FormType: form.DateRange}).
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("registered_city_id", "registered_city_id", db.Int8).
		FieldHide()
	info.AddField("法人", "legal_person", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("企业类型", "enterprise_type", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("开户行", "bank_name", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("银行账号", "bank_account", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("税号", "tax_number", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("实际地址", "actual_address", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("经营范围", "business_scope", db.Text).
		FieldEditAble().
		FieldHide()

	info.AddField("注册地址", "registered_address", db.Text).
		FieldHide()
	info.AddField("合作状态", "operate_status", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("代账许可", "bill_license", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble().
		FieldHide()
	info.AddField("联系方式", "contact", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("Email", "email", db.Text).
		FieldEditAble().
		FieldHide()
	info.AddField("主营类目", "main_category", db.Text).
		FieldEditAble().
		FieldHide()

	info.SetTable("merchants").SetTitle("商家信息表").SetDescription("Merchants")

	formList := merchants.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Timestamptz, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Deleted_at", "deleted_at", db.Timestamptz, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.AddField("商家名称", "name", db.Text, form.Text)
	formList.AddField("商家编码", "mcode", db.Text, form.Text)
	formList.AddField("商家类型", "mtype", db.Text, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Value: "服务商", Text: "服务商"},
			{Value: "承接商", Text: "承接商"},
			{Value: "不合作", Text: "不合作"},
		})
	formList.AddField("加盟关系", "join_relation", db.Text, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Value: "新引入", Text: "新引入"},
			{Value: "意向加盟商", Text: "意向加盟商"},
			{Value: "加盟商", Text: "加盟商"},
		})
	formList.AddField("合作关系", "cooperate_relation", db.Text, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Value: "不合作", Text: "不合作"},
			{Value: "待确定", Text: "待确定"},
			{Value: "合作", Text: "合作"},
		})
	formList.AddField("备注", "remarks", db.Text, form.TextArea)

	formList.AddField("引入日期", "introduction_date", db.Timestamptz, form.Date)
	formList.AddField("转承接商日期", "be_undertaker_date", db.Timestamptz, form.Date)
	formList.AddField("转服务商日期", "be_serviceprovider_date", db.Timestamptz, form.Date)
	formList.AddField("签约日期", "contract_date", db.Timestamptz, form.Date)
	formList.AddField("签约失效日期", "expiration_date", db.Timestamptz, form.Date)

	formList.AddField("注册地所在城市", "registered_city_id", db.Int8, form.Text).
		FieldDisplay(func(model types.FieldModel) interface{} {
			val, _ := strconv.Atoi(model.Value)
			db := models.DB
			var rCity models.RegisteredCity
			rCity.ID = uint(val)
			db.First(&rCity)
			// fmt.Println("rCity:%+v", rCity)

			return rCity.CityName
		})
	formList.AddField("法人", "legal_person", db.Text, form.Text)
	formList.AddField("企业类型", "enterprise_type", db.Text, form.Text)
	formList.AddField("开户行", "bank_name", db.Text, form.Text)
	formList.AddField("银行账号", "bank_account", db.Text, form.Text)
	formList.AddField("税号", "tax_number", db.Text, form.Text)
	formList.AddField("实际办公地址", "actual_address", db.Text, form.Text)

	formList.AddField("注册地址", "registered_address", db.Text, form.Text)
	formList.AddField("营业状态", "operate_status", db.Text, form.SelectSingle)
	formList.AddField("代账许可", "bill_license", db.Text, form.SelectSingle)
	formList.AddField("联系方式", "contact", db.Text, form.Text)
	formList.AddField("Email", "email", db.Text, form.Email)
	formList.AddField("主营品类", "main_category", db.Text, form.Text)

	formList.AddField("注册资本", "registered_capital", db.Int8, form.Number).FieldHelpMsg("单位：元")
	formList.AddField("成立日期", "establish_date", db.Timestamptz, form.Date)
	formList.AddField("营业执照起始日", "operate_start", db.Timestamptz, form.Date)
	formList.AddField("营业执照终止日", "operate_end", db.Timestamptz, form.Date)
	formList.AddField("营业执照经营范围", "business_scope", db.Text, form.TextArea)

	formList.AddField("代账许可起始日", "bill_license_start", db.Timestamptz, form.Date)
	formList.AddField("代账许可终止日", "bill_license_end", db.Timestamptz, form.Date)

	formList.SetTable("merchants").SetTitle("商家信息详情表").SetDescription("Merchants")

	return merchants
}
