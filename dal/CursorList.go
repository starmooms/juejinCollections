package dal

import (
	"github.com/cockroachdb/errors"

	"xorm.io/xorm"
)

type CursorBase struct {
	List   interface{} `form:"list" json:"list"`
	Limt   int         `form:"limt" json:"limt"`
	Cursor int         `form:"cursor" json:"cursor"`
}

func NewCursorList(list interface{}, limt int, cursor int) *CursorBase {
	if limt == 0 {
		limt = 10
	}

	return &CursorBase{
		List:   list,
		Limt:   limt,
		Cursor: cursor,
	}
}

func (c *CursorBase) GetCursorData(session *xorm.Session) error {
	err := session.Limit(c.Limt, c.Cursor).Find(c.List)
	if err != nil {
		return errors.Wrap(err, "db Err")
	}
	if c.List == nil {
		c.List = []interface{}{}
	}
	return nil
}
