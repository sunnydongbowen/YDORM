package entity

// Student 用Tag描述字段信息
type Student struct {
	TableName string "t_student"
	ID        string `field:"id" iskey:"1"`
	Name      string `field:"canme"`
	Age       int    `field:"age"`
	Sex       string
}
