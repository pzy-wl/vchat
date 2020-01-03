package ypg

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/weihaoranW/vchat/common/ytime"
)

type BaseModel struct {
	CreatedTime ytime.Date
	UpdatedTime ytime.Date
	//DelTime     ytime.Date
}

// // 注册新建钩子在持久化之前
func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		//nowTime := time.Now().Unix()
		now := ytime.OfNow()
		if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(now)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedTime"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(now)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdatedTime", ytime.OfNow())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DelTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
