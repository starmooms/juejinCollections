package collectReq

import (
	"fmt"
	"juejinCollections/dal"
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"
	"regexp"
	"time"

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

func (ac *Action) Run() {
	startTime := time.Now()

	err := dal.DbDal.OpenWal()
	if err != nil {
		tool.ShowErr(err)
	}

	ac.start()
	ac.wg.Wait()

	err = dal.DbDal.CloseWal()
	if err != nil {
		tool.ShowErr(err)
	}

	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)

	tFormat := "2006-01-02 15:04:05"
	sTime := startTime.Format(tFormat)
	eTime := endTime.Format(tFormat)
	logger.Logger.Infof(`{
		"start": "%s",
		"end": "%s",
		"run": "%v",
		"taskTotal": %d
	}`,
		sTime,
		eTime,
		latencyTime,
		ac.wg.TaskRunLen,
	)
}

func (ac *Action) start() {
	ac.wg.Add(func() {
		tagList, err := GetTagList(ac.UserId)
		if err != nil {
			tool.ShowErr(err)
			return
		}
		for _, tagItem := range *tagList {
			ac.saveCollectData(tagItem.TagId)
		}
	})
}

// 保存收藏夹文章
func (ac *Action) saveCollectData(tagId string) {
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
				ac.saveArticleImg(article)
			}
		}
	})
}

// 保存文章图片
func (ac *Action) saveArticleImg(m *model.Article) {
	if m == nil {
		return
	}

	imageResult := [][]string{}
	if m.MarkContent != "" {
		reg, err := regexp.Compile("!\\[.*?\\]\\((http.+?)(\\s.*?)*?\\)")
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

	if len(imageResult) > 2 {
		fmt.Println("..")
	}
	for i := 0; i < 100; i++ {
		c := i
		ac.wg.Add(func() {
			url := fmt.Sprintf("http://localhost:8012/%s/%d", m.ArticleId, c)
			logger.GetLog().Warn(url)
			err := GetImageData(url, m.ArticleId)
			if err != nil {
				tool.ShowErr(errors.Wrap(err, "Get image Request Error"))
			}
		})
	}

	// for _, rItem := range imageResult {
	// 	if len(rItem) >= 2 {
	// 		ac.wg.Add(func() {
	// 			err := GetImageData(rItem[1], m.ArticleId)
	// 			if err != nil {
	// 				tool.ShowErr(errors.Wrap(err, "Get image Request Error"))
	// 			}
	// 		})
	// 	}
	// }
}
