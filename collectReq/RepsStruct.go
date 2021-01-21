package collectReq

import (
	"juejinCollections/model"

	"github.com/cockroachdb/errors"
)

type ResBase struct {
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
}

func (r *ResBase) GetBase() (int, string) {
	return r.Err_no, r.Err_msg
}

// 请求收藏列表返回
type CollectListStruct struct {
	ResBase
	Data     []model.TagModel `json:"data"`
	Cursor   string           `json:"cursor"`
	Count    int              `json:"count"`
	Has_more bool             `json:"has_more"`
}

// 请求文章返回
type ArticleRes struct {
	ResBase
	// Data []byte `json:"data"`
}

type ResBaseData interface {
	GetBase() (int, string)
}

func CheckErr(res ResBaseData) error {
	Err_no, Err_msg := res.GetBase()
	if Err_no != 0 {
		return errors.NewWithDepth(1, "JJ RequestErr:"+Err_msg)
	}
	return nil
}
