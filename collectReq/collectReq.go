package collectReq

import (
	"encoding/json"
	"fmt"
	"juejinCollections/httpRequest"

	"github.com/gin-gonic/gin"
)

func setError(err error) {
	fmt.Println(err)
}

// 获取收藏列表
func GetList() error {
	// https://api.juejin.cn/interact_api/v1/collectionSet/list
	// 1116759544852221
	// 2664871913078168

	httpReq := httpRequest.Request(&httpRequest.HttpRequest{
		Url:    "https://api.juejin.cn/interact_api/v1/collectionSet/list",
		Method: "GET",
		Params: &gin.H{
			"user_id": 1116759544852221,
			"cursor":  0,
			"limit":   20,
		},
	})
	result, err := httpReq.DoRequest()
	if err != nil {
		setError(err)
		return err
	}

	reqCollectList := &CollectListStruct{}
	json.Unmarshal(*result.Data, reqCollectList)

	return nil
}
