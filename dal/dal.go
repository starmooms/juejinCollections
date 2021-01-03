package dal

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DbDal *Dal

func init() {
	DbDal = &Dal{}
	DbDal.Init()
}

// 添加收藏列表
func AddTags(list []TagModel) (int64, error) {
	model := &TagModel{}
	key := []string{"id", "tag_id", "tag_name", "color", "icon", "back_ground", "ctime", "mtime", "status", "creator_id", "user_name", "post_article_count", "concern_user_count", "isfollowed", "is_has_in", "create_time", "update_time"}
	sql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES`, model.TableName(), strings.Join(key, ","))

	fileKey := []string{}
	for _, val := range key {
		newstr := make([]rune, 0)
		name := strings.ToLower(val)
		parts := strings.Split(name, "_")
		for _, partItem := range parts {
			for i, r := range partItem {
				if i == 0 && 'a' <= r && r <= 'z' {
					r -= ('a' - 'A')
				}
				newstr = append(newstr, r)
			}
		}
		fileKey = append(fileKey, string(newstr))
	}
	fmt.Println(fileKey)

	// now := time.Now().Format("2006-01-02 15:04:05")
	for _, v := range list {
		val := reflect.ValueOf(v)
		var valArr = []string{}
		for _, keyItem := range fileKey {
			valArr = append(valArr, val.FieldByName(keyItem).String())
		}
		sql += "(" + strings.Join(valArr, ",") + ")"
		// key := reflect.TypeOf(v)
	}

	fmt.Println(sql)

	// typeInfo := reflect.ValueOf(model)
	// num := typeInfo.NumField()

	// DbDal.Engine.Exec(`INSERT INTO `)
	return DbDal.Engine.Insert(list)
}

type Dal struct {
	Engine *xorm.Engine
}

func (d *Dal) Init() error {

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("dal Error", err)
		return err
	}

	if err := engine.Ping(); err != nil {
		fmt.Println("dal Error", err)
		return err
	}

	engine.ShowSQL(true)                  // 打印sql语句
	engine.SetMapper(names.GonicMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
	engine.DatabaseTZ, _ = time.LoadLocation("Local")

	d.Engine = engine

	// has, err := engine.IsTableExist("collect")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if !has {
	// 	engine.Exec(`
	// 	CREATE TABLE "main"."collect" (
	// 		"id" INTEGER NOT NULL,
	// 		"tag_id" TEXT NOT NULL,
	// 		"tag_name" TEXT NOT NULL,
	// 		"color" TEXT,
	// 		"icon" TEXT,
	// 		"back_ground" TEXT,
	// 		"ctime" INTEGER,
	// 		"mtime" INTEGER NOT NULL,
	// 		"status" INTEGER,
	// 		"creator_id" INTEGER,
	// 		"user_name" TEXT,
	// 		"post_article_count" INTEGER,
	// 		"concern_user_count" INTEGER,
	// 		"isfollowed" INTEGER,
	// 		"is_has_in" INTEGER,
	// 		PRIMARY KEY ("id")
	// 	);
	// 	CREATE UNIQUE INDEX "main"."UQE_collect_id" ON "collect" ("id" ASC);
	// 	CREATE UNIQUE INDEX "main"."UQE_collect_tag_id" ON "collect" ("tag_id" ASC);
	// 	CREATE UNIQUE INDEX "main"."UQE_collect_tag_name" ON "collect" ("tag_name" ASC);
	// 	`)
	// }

	if err := engine.Sync2(new(TagModel)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
