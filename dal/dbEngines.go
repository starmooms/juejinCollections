package dal

import (
	"juejinCollections/model"
	"juejinCollections/tool"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DbDal *Dal

// 创建生成数据库
func NewDal() {
	DbDal = &Dal{}
	DbDal.Init()
}

type Dal struct {
	Engine *xorm.Engine
}

func (d *Dal) Init() error {

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	tool.PanicErr(err)

	engine.SetLogger(DalLogNew())

	err = engine.Ping()
	tool.PanicErr(err)

	//	engine.ShowSQL(true)                  // 打印sql语句
	engine.SetMapper(names.GonicMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
	engine.DatabaseTZ, _ = time.LoadLocation("Local")

	d.Engine = engine

	// // 手动生成表
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

	tool.PanicErr(engine.Sync2(new(model.TagModel)))
	tool.PanicErr(engine.Sync2(new(model.ArticleModel)))
	tool.PanicErr(engine.Sync2(new(model.TagArticleModel)))
	tool.PanicErr(engine.Sync2(new(model.Image)))

	return nil
}
