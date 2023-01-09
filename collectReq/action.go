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
	wg              *tool.WaitGroup
	UserId          string
	DbArticleId     []string
	requestCount    int
	newArticleCount int
}

func NewAction(userId string) *Action {
	return &Action{
		UserId:          userId,
		wg:              tool.NewWaitGroup(10),
		requestCount:    0,
		newArticleCount: 0,
	}
}

func (ac *Action) infof(format string, args ...interface{}) {
	logger.Logger.Infof(format, args...)
	ac.callLog(format, args...)
}

func (ac *Action) Errorf(err error) {
	err = tool.ShowErr(err)
	ac.callLog("%+v\n", err)
}

func (ac *Action) callLog(format string, args ...interface{}) {
	fmt.Print(fmt.Sprintf(format, args...))
}

func (ac *Action) Run() {
	startTime := time.Now()

	err := dal.DbDal.OpenWal()
	if err != nil {
		ac.Errorf(err)
	}

	ac.start()
	ac.wg.Wait()

	err = dal.DbDal.CloseWal()
	if err != nil {
		ac.Errorf(err)
	}

	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)

	tFormat := "2006-01-02 15:04:05"
	sTime := startTime.Format(tFormat)
	eTime := endTime.Format(tFormat)
	ac.infof(`{
		"start": "%s",
		"end": "%s",
		"run": "%v",
		"requestCount": "%d",
		"newArticleCount": "%d",
		"taskTotal": %d
	}`,
		sTime,
		eTime,
		latencyTime,
		ac.requestCount,
		ac.newArticleCount,
		ac.wg.TaskRunLen,
	)
}

func (ac *Action) start() {
	if len(ac.DbArticleId) > 0 {
		ac.refreshDbArticleImg()
	} else {
		ac.getAllCollect()
	}
}

// 更新请求数量
func (ac *Action) addRequestCount(count int) {
	ac.wg.GetLock(func() {
		ac.requestCount += count
		logger.Logger.Infof(`requestCount: %d`, ac.requestCount)
	})
}

// 更新新增文章数量
func (ac *Action) addNewArticleCount(count int) {
	ac.wg.GetLock(func() {
		ac.newArticleCount += count
	})
}

// 根据文章Id在本地数据库查找文章，更新图片
func (ac *Action) refreshDbArticleImg() {
	ac.wg.Add(func() {
		for _, articleId := range ac.DbArticleId {
			article := &model.Article{
				ArticleId: articleId,
			}
			ac.wg.Add(func() {
				has, err := dal.Get(article)
				if err != nil {
					ac.Errorf(err)
					return
				} else if !has {
					ac.Errorf(errors.New(article.ArticleId + "Not Found In Db"))
					return
				}
				ac.saveArticleImg(article)
			})
		}
	})
}

// 请求收藏列表，获取全部文章
func (ac *Action) getAllCollect() {
	ac.wg.Add(func() {
		tagList, err := GetTagList(ac.UserId)
		ac.addRequestCount(1)
		if err != nil {
			ac.Errorf(err)
			return
		}

		if _, err := dal.AddTags(tagList); err != nil {
			ac.Errorf(err)
			return
		}

		for _, tagItem := range *tagList {
			ac.saveCollectData(tagItem.CollectionId)
		}
	})
}

// 保存收藏夹文章
func (ac *Action) saveCollectData(tagId string) {
	ac.wg.Add(func() {
		var collectData = &CollectArticle{}
		collectData.Has_more = true
		cursor := 0

		hasError := func(err error) bool {
			if err != nil {
				ac.Errorf(errors.Wrapf(err, "SaveCollectData Err At Cursor:%d", cursor))
				collectData = nil
				return true
			}
			return false
		}

		for collectData != nil && collectData.HasMore() {
			newData, allArticleList, tagArticle, err := GetCollectData(tagId, cursor)
			ac.addRequestCount(1)

			if hasError(err) {
				return
			}

			// 只添加不存在的文章
			newArticleList := []*model.Article{}
			for _, article := range *allArticleList {
				has, err := dal.HasArticel(article)
				if hasError(err) {
					return
				}
				if !has {
					newArticleList = append(newArticleList, article)
				}
			}

			if _, err = dal.AddArticle(&newArticleList); hasError(err) {
				return
			}

			if _, err = dal.AddTagArticle(tagArticle); hasError(err) {
				return
			}

			ac.addNewArticleCount(len(newArticleList))

			collectData = newData
			cursor += len(*allArticleList)
			for _, article := range newArticleList {
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

	hasError := func(err error) bool {
		if err != nil {
			ac.Errorf(errors.Wrap(err, "Get image Reg Error"))
			return true
		}
		return false
	}

	imageResult := [][]string{}
	if m.MarkContent != "" {
		reg, err := regexp.Compile("!\\[.*?\\]\\((http.+?)(\\s.*?)*?\\)")
		if hasError(err) {
			return
		}
		imageResult = reg.FindAllStringSubmatch(m.MarkContent, -1)
	} else if m.Content != "" {
		reg, err := regexp.Compile("<img.*?src=\"(http.+?)\".*?>")
		if hasError(err) {
			return
		}
		imageResult = reg.FindAllStringSubmatch(m.Content, -1)
	}

	// for _, rItem := range imageResult {
	// 	if len(rItem) >= 2 {
	// 		logger.Logger.Warn(rItem[1], m.ArticleId)
	// 		// ac.wg.Add(func() {
	// 		// 	err := GetImageData(rItem[1], m.ArticleId)
	// 		// 	if err != nil {
	// 		// 		tool.ShowErr(errors.Wrap(err, "Get image Request Error"))
	// 		// 	}
	// 		// })
	// 	}
	// }

	// if len(imageResult) > 2 {
	// 	fmt.Println("..")
	// }
	// for i := 0; i < 100; i++ {
	// 	c := i
	// 	ac.wg.Add(func() {
	// 		url := fmt.Sprintf("http://localhost:8012/%s/%d", m.ArticleId, c)
	// 		logger.GetLog().Warn(url)
	// 		err := GetImageData(url, m.ArticleId)
	// 		if err != nil {
	// 			tool.ShowErr(errors.Wrap(err, "Get image Request Error"))
	// 		}
	// 	})
	// }

	for _, rItem := range imageResult {
		if len(rItem) >= 2 {
			imgUrl := rItem[1]
			ac.wg.Add(func() {
				// log.Warn(imgUrl)
				has, err := dal.HasImage(imgUrl, m.ArticleId)
				if hasError(err) || has {
					return
				}

				image, err := GetImageData(imgUrl, m.ArticleId)
				ac.addRequestCount(1)
				if hasError(err) {
					return
				}

				_, err = dal.AddImage(image)
				hasError(err)
			})
		}
	}
}
