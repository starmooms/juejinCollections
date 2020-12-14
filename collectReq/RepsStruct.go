package collectReq

type CollectListStruct struct {
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
	Data    []struct {
		Id                 int    `json:"id"`
		Tag_id             string `json:"tag_id"`
		Tag_name           string `json:"tag_name"`
		Color              string `json:"color"`
		Icon               string `json:"icon"`
		Back_ground        string `json:"back_ground"`
		Ctime              int    `json:"ctime"`
		Mtime              int    `json:"mtime"`
		Status             int    `json:"status"`
		Creator_id         int    `json:"creator_id"`
		User_name          string `json:"user_name"`
		Post_article_count int    `json:"post_article_count"`
		Concern_user_count int    `json:"concern_user_count"`
		Isfollowed         bool   `json:"isfollowed"`
		Is_has_in          bool   `json:"is_has_in"`
	} `json:"data"`
	Cursor   string `json:"cursor"`
	Count    int    `json:"count"`
	Has_more bool   `json:"has_more"`
}
