package dal

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func init() {
	engine, err := xorm.NewEngine("sqlite3", "./db/test.db")
	if err != nil {
		fmt.Println("dal Error", err)
		return
	}
	engine.Ping()

	if err != nil {
		fmt.Println("dal Error", err)
	}

	engine.ShowSQL(true)                 // 打印sql语句
	engine.SetMapper(names.SameMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
}
