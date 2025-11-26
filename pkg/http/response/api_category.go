package response

type Post_Api_Category_Response struct {
	CategoryID uint `json:"category_id"`
}

type Get_Api_Category_List_Response struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type Get_Api_Category_Detail_Response struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	// 创建者ID
	CreatorID uint `json:"creator_id"`
	// 创建者昵称
	CreatorNickName string `json:"creator_nick_name"`
}
