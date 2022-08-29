package ydorm

import (
	"fmt"
)

func Save(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("err:", err)
		}
	}()
	strSql, p := genInsertSQL(entity)
	fmt.Println(strSql)
	fmt.Println(p)
	// 执行sQL
	isOk = true
	return
}

func Update(entity interface{}) (isOK bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOK = false
			fmt.Println("err:", err)
		}
	}()
	strSql, p := genUpdateSQL(entity)
	fmt.Println(strSql)
	fmt.Println(p)
	// 执行sql
	isOK = true
	panic("模拟更新失败")
	return
}

func Delete(enetity interface{}) (isOK bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOK = false
			fmt.Println("err:", err)
		}
	}()
	strSql, p := genDeleteSQL(enetity)
	fmt.Println(strSql)
	fmt.Println(p)
	// 执行SQL
	isOK = true
	return

}
