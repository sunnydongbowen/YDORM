package main

import (
	"fmt"
	"ydorm/entity"
	"ydorm/ydorm"
)

func main() {
	stu := entity.Student{
		ID:   "Y001",
		Name: "jack",
		Age:  25,
		Sex:  "男",
	}
	if isOk, _ := ydorm.Save(&stu); isOk {
		fmt.Println("新增成功")
	}
	if isOk, _ := ydorm.Update(&stu); isOk {
		fmt.Println("更新成功")
	}
	if isOk, _ := ydorm.Delete(&stu); isOk {
		fmt.Println("删除成功")
	}
}
