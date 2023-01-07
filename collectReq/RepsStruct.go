package collectReq

import (
	"juejinCollections/model"

	"github.com/cockroachdb/errors"
)

type ResBase struct {
	Err_no  *int   `json:"err_no"`
	Err_msg string `json:"err_msg"`
}

func (r *ResBase) CheckErr() error {
	errNo := r.Err_no
	if errNo == nil || *errNo != 0 {
		return errors.Newf("juejing RequestErr: Code %d\n %s\n", *errNo, r.Err_msg)
	}
	return nil
}

// 分页列表
type ResCursor struct {
	Cursor string `json:"cursor"`
	// Count    int    `json:"count"`
	Has_more bool `json:"has_more"`
}

func (r *ResCursor) HasMore() bool {
	return r.Has_more
}

// 请求收藏列表返回
type CollectListStruct struct {
	ResBase
	ResCursor
	Data []model.Tag `json:"data"`
}

// 请求文章返回
type ArticleRes struct {
	ResBase
	// Data []byte `json:"data"`
}

// 收藏夹中的文章列表
type CollectArticle struct {
	ResBase
	ResCursor
}
