package yqiniu

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/vhaoran/vchat/common/reflectUtils"
)

func Key2UrlOfQiNiu(obj interface{}, fields ...string) error {
	if !reflectUtils.IsStruct(obj) {
		return errors.New("不是结构")
	}
	if !reflectUtils.IsPointer(obj) {
		return errors.New("不是指针，无法继续")
	}

	v := reflect.Indirect(reflect.ValueOf(obj))
	//
	for _, fdName := range fields {
		if _, ok := v.Type().FieldByName(fdName); ok {
			//i := fd.Index
			key := v.FieldByName(fdName).String()
			fmt.Println(" --- fdValue:", key)
			url := GetVisitURL(key)
			//
			fmt.Println(" --- url:", url)
			//
			v.FieldByName(fdName).SetString(url)
			continue
		}
	}
	return nil
}