package tables

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/template/types"
)

var dateDisplay = func(model types.FieldModel) interface{} {
	if model.Value == "" {
		return ""
	}
	display := model.Value[:9]
	if display[:4] == "0001" {
		display = ""
	}
	fmt.Println("display:", display)
	return display
}
