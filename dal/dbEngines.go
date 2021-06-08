package dal

import (
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"
	"time"

	"github.com/cockroachdb/errors"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DbDal *Dal
var logs = logger.Logger

const (
	MODEL_WAL    = "wal"
	MODEL_DELETE = "delete"
)

// 创建生成数据库
func NewDal(dbFile string) {
	DbDal = &Dal{
		DbFile: dbFile,
	}
	DbDal.init()
}

type Dal struct {
	Engine *xorm.Engine
	DbFile string
}

func (d *Dal) init() {
	engine, err := d.newEngine()
	if err != nil {
		tool.PanicErr(err)
	}
	d.Engine = engine
	tool.PanicErr(engine.Sync2(new(model.Tag)))
	tool.PanicErr(engine.Sync2(new(model.Article)))
	tool.PanicErr(engine.Sync2(new(model.TagArticleId)))
	tool.PanicErr(engine.Sync2(new(model.Image)))
}

func (d *Dal) newEngine() (*xorm.Engine, error) {
	var err error

	engine, err := xorm.NewEngine("sqlite3", d.DbFile) // ?cache=shared&mode=rwc&_journal_mode=WAL
	if err != nil {
		return nil, err
	}

	// err = engine.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	engine.SetLogger(DalLogNew())
	// engine.ShowSQL(true)                  // 打印sql语句
	engine.SetMapper(names.GonicMapper{}) // 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名
	engine.DatabaseTZ, err = time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func (d *Dal) GetMode(engine *xorm.Engine) (string, error) {
	v, err := engine.QueryString("PRAGMA journal_mode;")
	if err != nil {
		return "", errors.Wrap(err, "get journal_mode error")
	}
	logs.Debug("get journal_mode is ", v)
	return v[0]["journal_mode"], nil
}

// 打开wal模式
func (d *Dal) OpenWal() error {
	mode := MODEL_WAL
	nowMode, err := d.GetMode(d.Engine)
	if err != nil {
		return err
	}
	if nowMode != mode {
		_, err = d.Engine.Exec("PRAGMA journal_mode = " + mode + ";")
		if err != nil {
			return errors.Wrap(err, "Dal Open wal Error")
		}
	}
	return nil
}

// 关闭wal模式
func (d *Dal) CloseWal() error {
	nowMode, err := d.GetMode(d.Engine)
	if err != nil {
		return err
	}

	if nowMode == MODEL_WAL {
		// 关闭到只剩一个连接，才能设置取消wal模式
		dbStatus := d.Engine.DB().Stats()
		maxConnect := dbStatus.MaxOpenConnections
		d.Engine.SetMaxOpenConns(1)
		defer func() {
			d.Engine.SetMaxOpenConns(maxConnect)
		}()

		_, err = d.Engine.Exec("PRAGMA journal_mode = " + MODEL_DELETE + ";")
		if err != nil {
			return errors.Wrap(err, "Dal Close wal change mode Error")
		}
	}

	return nil
}
