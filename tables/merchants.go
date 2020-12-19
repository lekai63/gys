package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMerchantsTable(ctx *context.Context) table.Table {

	merchants := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := merchants.GetInfo()

	info.AddField("序号", "id", db.Int8).
		FieldSortable()
	info.AddField("Created_at", "created_at", db.Timestamptz).
		FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamptz).
		FieldHide()
	info.AddField("Deleted_at", "deleted_at", db.Timestamptz).
		FieldHide()
	info.AddField("引入日期", "introduction_date", db.Timestamptz).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("转承接商日期", "be_undertaker_date", db.Timestamptz).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("转服务商日期", "be_serviceprovider_date", db.Timestamptz).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("签约日期", "contract_date", db.Timestamptz).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("失效日期", "expiration_date", db.Timestamptz).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("备注", "remarks", db.Text).
		FieldEditAble()
	info.AddField("商家名称", "name", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("商家编码", "mcode", db.Text).
		FieldSortable().
		FieldEditAble()
	info.AddField("商家类型", "mtype", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("加盟关系", "join_relation", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("合作关系", "cooperate_relation", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()

	info.SetTable("merchants").SetTitle("商家信息表").SetDescription("Merchants")

	formList := merchants.GetForm()
	formList.AddField("序号", "id", db.Int8, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHideWhenUpdate()
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
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("引入日期", "introduction_date", db.Timestamptz, form.Date)
	formList.AddField("转承接商日期", "be_undertaker_date", db.Timestamptz, form.Date)
	formList.AddField("转服务商日期", "be_serviceprovider_date", db.Timestamptz, form.Date)
	formList.AddField("签约日期", "contract_date", db.Timestamptz, form.Date)
	formList.AddField("失效日期", "expiration_date", db.Timestamptz, form.Date)
	formList.AddField("备注", "remarks", db.Text, form.RichText)
	formList.AddField("商家名称", "name", db.Text, form.Text)
	formList.AddField("商家编码", "mcode", db.Text, form.Text)
	formList.AddField("商家类型", "mtype", db.Text, form.Text)
	formList.AddField("加盟关系", "join_relation", db.Text, form.Text)
	formList.AddField("合作关系", "cooperate_relation", db.Text, form.Text)

	formList.SetTable("merchants").SetTitle("商家信息详情").SetDescription("Merchants")

	return merchants
}
