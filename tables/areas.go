package tables

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAreasTable(ctx *context.Context) table.Table {

	areas := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := areas.GetInfo()

	info.AddField("Id", "id", db.Int8).
		FieldHide()
	info.AddField("Created_at", "created_at", db.Timestamptz).
		FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamptz).
		FieldHide()
	info.AddField("Deleted_at", "deleted_at", db.Timestamptz).
		FieldHide()
	info.AddField("城市名称", "city", db.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}). // 模糊查询
		FieldSortable().
		FieldEditAble()

	info.AddField("商家类型", "mtype", db.Text).
		FieldJoin(types.Join{
			Table:     "merchant_areas",
			JoinField: "area_id",
			Field:     "id",
			BaseTable: "areas",
		}).FieldJoin(types.Join{
		Table:     "merchants",
		JoinField: "id",
		Field:     "merchant_id",
		BaseTable: "merchant_areas",
	}).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptionsFromTable("merchants", "mtype", "mtype", func(sql *db.SQL) *db.SQL {
			return sql.GroupBy("mtype")
		}).
		/* 	FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{
			{Value: "服务商", Text: "服务商"},
			{Value: "承接商", Text: "承接商"},
			{Value: "不合作", Text: "不合作"},
		}). */
		FieldDisplay(func(model types.FieldModel) interface{} {
			// 删除非中文字符
			chiReg := regexp.MustCompile("[^\u4e00-\u9fa5（）]")
			sliceMatch := chiReg.ReplaceAllString(model.Value, " ")

			// TODO:dirty code need to be modified
			var s1, s2, s3 string
			count1 := strings.Count(model.Value, "服务商")
			if count1 != 0 {
				s1 = "服务商x" + strconv.Itoa(count1)
			}
			count2 := strings.Count(model.Value, "承接商")
			if count2 != 0 {
				s2 = "承接商x" + strconv.Itoa(count2)
			}
			count3 := strings.Count(model.Value, "不合作")
			if count3 != 0 {
				s3 = "不合作" + strconv.Itoa(count3)
			}
			return sliceMatch + "\r\n汇总:"+s1+s2+s3
		}) 
	// FieldSortable().
	// FieldEditAble()

	info.AddField("商家名称", "name", db.Text).
		FieldJoin(types.Join{
			Table:     "merchant_areas",
			JoinField: "area_id",
			Field:     "id",
			BaseTable: "areas",
		}).FieldJoin(types.Join{
		Table:     "merchants",
		JoinField: "id",
		Field:     "merchant_id",
		BaseTable: "merchant_areas",
	}).FieldDisplay(func(model types.FieldModel) interface{} {
		// fmt.Printf("model:%+v", model)
		// 删除非中文字符
		chiReg := regexp.MustCompile("[^\u4e00-\u9fa5（）]")
		s := chiReg.ReplaceAllString(model.Value, "    ")
		return s
	})
	/* 		.FieldFilterOptions(types.FieldOptions{
	   			{Value: "服务商", Text: "服务商"},
	   		{Value: "承接商", Text: "承接商"},
	   		{Value: "所有合作商家", Text: "所有合作商家"},
	   		{Value: "不合作", Text: "不合作"},
	   		{Value: "所有商家（含不合作）", Text: "所有商家（含不合作）"},
	   		}).
	FieldDisplay(func(model types.FieldModel) interface{} {
		orm := models.DB
		var area models.Area
		var merchants []models.Merchant

		// fmt.Printf("info:%+v", info)

		area.ID = uint((model.Row["id"].(int64)))
		orm.Model(&area).Association("Merchants").Find(&merchants)
		// fmt.Println("merchants:", merchants)
		shows := make([]string, 0)
		for _, merchant := range merchants {
			shows = append(shows, merchant.Name)
		}
		return strings.Join(shows, ",")
	}) */

	info.SetTable("areas").SetTitle("城市维度").SetDescription("Areas")

	formList := areas.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
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
	formList.AddField("City", "city", db.Text, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("areas").SetTitle("城市维度详情页").SetDescription("Areas")

	return areas
}

