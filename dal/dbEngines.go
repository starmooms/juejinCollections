package dal

import (
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DbDal *Dal
var logs = logger.GetLog()

const (
	MODEL_WAL    = "wal"
	MODEL_DELETE = "delete"
)

// 创建生成数据库
func NewDal() {
	DbDal = &Dal{}
	DbDal.init()
}

type Dal struct {
	Engine  *xorm.Engine
	walMode struct {
		walEngine *xorm.Engine
		delEngine *xorm.Engine
	}
}

func (d *Dal) init() {
	engine, err := d.newEngine()
	if err != nil {
		tool.PanicErr(err)
	}
	d.Engine = engine
	tool.PanicErr(engine.Sync2(new(model.TagModel)))
	tool.PanicErr(engine.Sync2(new(model.ArticleModel)))
	tool.PanicErr(engine.Sync2(new(model.TagArticleModel)))
	tool.PanicErr(engine.Sync2(new(model.Image)))
}

func (d *Dal) newEngine() (*xorm.Engine, error) {
	var err error

	engine, err := xorm.NewEngine("sqlite3", "./test.db?cache=shared&mode=rwc&_journal_mode=WAL") // ?cache=shared&mode=rwc&_journal_mode=WAL
	if err != nil {
		return nil, err
	}

	// err = engine.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	engine.SetLogger(DalLogNew())
	//engine.ShowSQL(true)                  // 打印sql语句
	engine.SetMapper(names.GonicMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
	engine.DatabaseTZ, err = time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}

	return engine, nil
}

// func (d *Dal) GetMode(engine *xorm.Engine) (string, error) {
// 	v, err := engine.QueryString("PRAGMA journal_mode;")
// 	if err != nil {
// 		return "", errors.Wrap(err, "get journal_mode error")
// 	}
// 	logs.Error(v)
// 	return v[0]["journal_mode"], nil
// }

// func (d *Dal) OpenWal() error {
// 	mode := MODEL_WAL
// 	nowMode, err := d.GetMode(d.Engine)
// 	if err != nil {
// 		return err
// 	}
// 	if nowMode != mode {
// 		walEngine, err := d.newEngine()
// 		if err != nil {
// 			return errors.Wrap(err, "Dal create wal Engine Error")
// 		}

// 		_, err = walEngine.Exec("PRAGMA journal_mode=" + mode + ";")
// 		if err != nil {
// 			return errors.Wrap(err, "Dal Open wal Error")
// 		}

// 		d.walMode.delEngine = d.Engine
// 		d.walMode.walEngine = walEngine
// 		d.Engine = walEngine
// 	}
// 	return nil
// }

// // 关闭wal模式
// func (d *Dal) CloseWal() error {
// 	nowMode, err := d.GetMode(d.Engine)
// 	if err != nil {
// 		return err
// 	}

// 	// 先关闭当前多个wal的db（engine中有多个db，只能先全部关闭）
// 	// 再生成一个新的db（修改mode为delete）
// 	if nowMode == MODEL_WAL {
// 		d.Engine.Close()

// 		closeEngine, err := d.newEngine()
// 		if err != nil {
// 			return errors.Wrap(err, "Dal create wal closeEngine Error")
// 		}
// 		_, err = closeEngine.Exec("PRAGMA journal_mode=" + MODEL_DELETE + ";")
// 		if err != nil {
// 			return errors.Wrap(err, "Dal Close wal change mode Error")
// 		}
// 		closeEngine.Close()

// 		delEngine := d.walMode.delEngine
// 		if delEngine == nil {
// 			delEngine = closeEngine
// 		} else {
// 			closeEngine.Close()
// 		}
// 		d.Engine = delEngine
// 	}

// 	return nil
// }
