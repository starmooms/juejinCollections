package main

import (
	"fmt"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func getEng() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("dal Error", err)
		return nil
	}

	engine.ShowSQL(true)                  // 打印sql语句
	engine.SetMapper(names.GonicMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
	return engine
}

type rowItem struct {
	Id   int
	Name string
}

func (m *rowItem) TableName() string {
	return "test_sql"
}

func init() {
	engine := getEng()

	s := "\"计算机基础'\""
	row := []rowItem{}
	err2 := engine.Where("name= ?", s).Find(&row)
	fmt.Println("get row", s, err2, &row)

	// row2, err1 := engine.Query("SELECT * FROM `test_sql` WHERE name= ?", s)
	// fmt.Println(string(row2), err1)

}
