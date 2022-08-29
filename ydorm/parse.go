package ydorm

import (
	"reflect"
	"strings"
)

func parseEntity(entity interface{}) (tInfo *TableInfo, err error) {
	tInfo = &TableInfo{} // 返回值
	rt := reflect.TypeOf(entity)
	rv := reflect.ValueOf(entity)
	// 如果是指针，需要用Elem()
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	// 字段解析
	for i := 0; i < rt.NumField(); i++ {
		rtf := rt.Field(i) // struct field
		rvf := rv.Field(i) // struct value
		var f FieldInfo
		// 没有Tag,结构体字段和表字段名一致
		if rtf.Tag == "" {
			f = FieldInfo{rtf.Name, false, rvf}
		} else {
			strTag := string(rtf.Tag)
			if strings.Index(strTag, ":") == -1 {
				// Tag中没有:时，为表名字段
				tInfo.Name = strTag
				continue
			} else {
				// 解析tag为field键为表名字段
				field := rtf.Tag.Get("field")
				// 解析tag中的pk
				isKey := false
				strIskey := rtf.Tag.Get("iskey")
				if strIskey == "1" {
					isKey = true
				}
				f = FieldInfo{Name: field, IsPrimaryKey: isKey, refValue: rvf}
			}
		}
		tInfo.Fields = append(tInfo.Fields, f)
	}
	return
}
