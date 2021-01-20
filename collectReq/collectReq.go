package collectReq

import (
	"encoding/json"
	"fmt"
	"juejinCollections/dal"
	"juejinCollections/httpRequest"
	"juejinCollections/model"

	"github.com/buger/jsonparser"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

var dataMock Mock

func init() {
	dataMock.Init()
	GetArticle("6844904034181070861")
}

func setError(err error) error {
	wrapError := errors.NewWithDepth(1, err.Error())
	fmt.Printf("%+v \n", wrapError)
	return wrapError
}

func setNewError(msg string) error {
	wrapError := errors.NewWithDepth(1, msg)
	fmt.Printf("%+v \n", wrapError)
	return wrapError
}

func getReqUrl(url string) string {
	return "https://api.juejin.cn" + url
}

// 获取收藏列表
func GetTagList() error {
	// https://api.juejin.cn/interact_api/v1/collectionSet/list
	// 1116759544852221
	// 2664871913078168

	httpReq := httpRequest.Request(&httpRequest.HttpRequest{
		Url:    "https://api.juejin.cn/interact_api/v1/collectionSet/list",
		Method: "GET",
		Params: &gin.H{
			"user_id": 1116759544852221,
			"cursor":  0,
			"limit":   200,
		},
	})
	// result, err := httpReq.DoRequest()
	// if err != nil {
	// 	setError(err)
	// 	return err
	// }
	// reqCollectList := &CollectListStruct{}
	// json.Unmarshal(*result.Data, reqCollectList)

	fmt.Println(httpReq.Method)
	b := []byte(`{"err_no":0,"err_msg":"success","data":[{"id":3521869,"tag_id":"6906660674944909325","tag_name":"环境配置","color":"","icon":"","back_ground":"","ctime":1608082524,"mtime":1608082524,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3514054,"tag_id":"6902573960622243847","tag_name":"浏览器","color":"","icon":"","back_ground":"","ctime":1607130915,"mtime":1608689662,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":2,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3513229,"tag_id":"6902199513092784142","tag_name":"项目架构","color":"","icon":"","back_ground":"","ctime":1607043560,"mtime":1607043560,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3506057,"tag_id":"6899234387330760712","tag_name":"设计模式","color":"","icon":"","back_ground":"","ctime":1606353240,"mtime":1606353240,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3485313,"tag_id":"6889585424071655431","tag_name":"React","color":"","icon":"","back_ground":"","ctime":1604106712,"mtime":1609204112,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":4,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3481502,"tag_id":"6888099531817222152","tag_name":"自动构建","color":"","icon":"","back_ground":"","ctime":1603760649,"mtime":1608944273,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":3,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3436952,"tag_id":"6867313903076900878","tag_name":"题目","color":"","icon":"","back_ground":"","ctime":1598921135,"mtime":1598921135,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3407523,"tag_id":"6855129021274030093","tag_name":"Electron","color":"","icon":"","back_ground":"","ctime":1595983684,"mtime":1600220774,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3402305,"tag_id":"6854573259401363464","tag_name":"Go","color":"","icon":"","back_ground":"","ctime":1595482691,"mtime":1596130529,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3396145,"tag_id":"6850418478235697159","tag_name":"视频","color":"","icon":"","back_ground":"","ctime":1594948667,"mtime":1596130521,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3385227,"tag_id":"6847903905662107661","tag_name":"ES6","color":"","icon":"","back_ground":"","ctime":1594083741,"mtime":1603241351,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":7,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3363923,"tag_id":"6845244365892386829","tag_name":"css","color":"","icon":"","back_ground":"","ctime":1592008829,"mtime":1607908395,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":10,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3358919,"tag_id":"6845244350209736711","tag_name":"git","color":"","icon":"","back_ground":"","ctime":1591664434,"mtime":1596129116,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":1,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3357437,"tag_id":"6845244346153828365","tag_name":"计算机基础","color":"","icon":"","back_ground":"","ctime":1591580354,"mtime":1596129114,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":2,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3392128,"tag_id":"6845244335156527118","tag_name":"网络","color":"","icon":"","back_ground":"","ctime":1591145574,"mtime":1606790665,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":6,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3347894,"tag_id":"6845244313798967304","tag_name":"typescript","color":"","icon":"","back_ground":"","ctime":1590541388,"mtime":1607391754,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":7,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3369655,"tag_id":"6845244383915147272","tag_name":"监控","color":"","icon":"","back_ground":"","ctime":1590022614,"mtime":1607735415,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":7,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3380757,"tag_id":"6845243956976943111","tag_name":"node","color":"","icon":"","back_ground":"","ctime":1588814378,"mtime":1599266463,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":10,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3360177,"tag_id":"6845244354240446472","tag_name":"webpack","color":"","icon":"","back_ground":"","ctime":1588120979,"mtime":1607068910,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":9,"concern_user_count":0,"isfollowed":false,"is_has_in":false},{"id":3385719,"tag_id":"6845244203853676557","tag_name":"算法","color":"","icon":"","back_ground":"","ctime":1578101494,"mtime":1606699418,"status":0,"creator_id":1116759544852221,"user_name":"starmooms","post_article_count":19,"concern_user_count":0,"isfollowed":false,"is_has_in":false}],"cursor":"20","count":0,"has_more":true}`)
	reqCollectList := &CollectListStruct{}
	err := json.Unmarshal(b, reqCollectList)
	if err != nil {
		return setError(err)
	}

	if _, err := dal.AddTags(&reqCollectList.Data); err != nil {
		return setError(err)
	}

	return nil
}

func GetArticle(id string) error {
	httpReq := httpRequest.Request(&httpRequest.HttpRequest{
		Url:    getReqUrl("/content_api/v1/article/detail"),
		Method: "POST",
		Params: &gin.H{
			"article_id": id,
		},
	})

	// result, err := httpReq.DoRequest()
	// if err != nil {
	// 	return setError(err)
	// }
	// res := &ArticleRes{}
	// if err := json.Unmarshal(*result.Data, res); err != nil {
	// 	return setError(err)
	// }

	fmt.Println(httpReq.Url)
	var err error
	res := &ArticleRes{}
	if err = json.Unmarshal(*dataMock.Article, res); err != nil {
		return setError(err)
	}
	if res.Err_no != 0 {
		return setNewError("Request Back Error:" + res.Err_msg)
	}
	hitsjson, _, _, err := jsonparser.Get(*dataMock.Article, "data", "article_info")
	if err != nil {
		return setError(err)
	}
	article := &model.ArticleModel{}
	if err = json.Unmarshal(hitsjson, article); err != nil {
		return setError(err)
	}
	fmt.Println(article)

	fmt.Println(res)
	return nil
}
