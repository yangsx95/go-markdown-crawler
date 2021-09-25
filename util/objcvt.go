package util

import (
	"fmt"
	"reflect"
)

// CopyStructFields 拷贝结构体对象字段到另一个结构体中，通过标签copy并配合反射来完成
// 注意，需要进行拷贝的字段一定要是exports的
func CopyStructFields(originPointer interface{}, targetPointer interface{}) error {
	if originPointer == nil ||
		targetPointer == nil ||
		reflect.TypeOf(originPointer).Kind() != reflect.Ptr ||
		reflect.TypeOf(targetPointer).Kind() != reflect.Ptr {
		return fmt.Errorf("参数类型不能为nil且必须都是指针类型")
	}
	oT := reflect.TypeOf(originPointer).Elem()
	oV := reflect.ValueOf(originPointer).Elem()
	tT := reflect.TypeOf(targetPointer).Elem()
	tV := reflect.ValueOf(targetPointer).Elem()

	for i := 0; i < oT.NumField(); i++ {
		targetFieldName := oT.Field(i).Tag.Get("copy")
		// 判断t是否有该字段
		_, hasField := tT.FieldByName(targetFieldName)
		if !hasField {
			continue
		}
		// 如果有该字段
		// 将o的该字段的值赋值给t的该字段的值
		tV.FieldByName(targetFieldName).Set(oV.Field(i))
	}
	return nil
}
