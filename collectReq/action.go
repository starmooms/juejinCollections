package collectReq

import (
	"juejinCollections/model"
	"juejinCollections/tool"
	"regexp"

	"github.com/cockroachdb/errors"
)

type Action struct {
	wg     *tool.WaitGroup
	UserId string
}

func NewAction(userId string) *Action {
	return &Action{
		UserId: userId,
		wg:     tool.NewWaitGroup(10),
	}
}

func (ac *Action) Start() {
	ac.wg.Add(func() {
		tagList, err := GetTagList(ac.UserId)
		if err != nil {
			tool.ShowErr(err)
			return
		}
		for _, tagItem := range *tagList {
			ac.SaveCollectData(tagItem.TagId)
		}
	})
	ac.wg.Wait()
}

// 保存收藏夹文章
func (ac *Action) SaveCollectData(tagId string) {
	ac.wg.Add(func() {
		var collectData = &CollectArticle{}
		collectData.Has_more = true
		cursor := 0

		for collectData != nil && collectData.HasMore() {
			newData, articleList, err := GetCollectData(tagId, cursor)
			if err != nil {
				tool.ShowErr(errors.Wrapf(err, "SaveCollectData Err At Cursor:%d", cursor))
				collectData = nil
				return
			}

			collectData = newData
			cursor += len(*articleList)
			for _, article := range *articleList {
				ac.SaveArticleImg(article)
			}
		}
	})
}

// 保存文章图片
func (ac *Action) SaveArticleImg(m *model.ArticleModel) {
	if m == nil {
		return
	}

	imageResult := [][]string{}
	if m.MarkContent != "" {
		reg, err := regexp.Compile("!\\[.*?\\]\\((http.+?)\\)")
		if err != nil {
			tool.ShowErr(errors.Wrap(err, "Get image Reg Error"))
			return
		}
		imageResult = reg.FindAllStringSubmatch(m.MarkContent, -1)
	} else if m.Content != "" {
		reg, err := regexp.Compile("<img.*?src=\"(http.+?)\".*?>")
		if err != nil {
			tool.ShowErr(errors.Wrap(err, "Get image Reg Error"))
			return
		}
		imageResult = reg.FindAllStringSubmatch(m.Content, -1)
	}

	for _, rItem := range imageResult {
		if len(rItem) == 2 {
			ac.wg.Add(func() {
				err := GetImageData(rItem[1], m.ArticleId)
				if err != nil {
					tool.ShowErr(err)
				}
			})
		}
	}
}
