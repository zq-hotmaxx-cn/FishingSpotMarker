package service

import (
	"FishingSpotMarker/core/repository"
	"FishingSpotMarker/pkg/http/request"
	"FishingSpotMarker/pkg/http/response"
	"FishingSpotMarker/pkg/model"
	"FishingSpotMarker/pkg/utils/user"

	"gorm.io/gorm"
)

func Post_Api_Category_Service(
	request *request.Post_Api_Category_Request,
	auth_user user.User,
) (*response.Post_Api_Category_Response, error) {
	var resp *response.Post_Api_Category_Response

	// 检查用户是否存在
	exist, err := repository.ExistsUser(auth_user.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	category := &model.Category{
		CreatorID:   auth_user.ID,
		Name:        request.Name,
		Description: request.Description,
	}

	if err := repository.CreateCategory(category); err != nil {
		return nil, err
	}

	resp = &response.Post_Api_Category_Response{
		CategoryID: category.ID,
	}

	return resp, nil
}

func Get_Api_Category_List_Service(auth_user user.User) ([]response.Get_Api_Category_List_Response, error) {
	var resp []response.Get_Api_Category_List_Response

	// 检查用户是否存在
	exist, err := repository.ExistsUser(auth_user.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	categories, err := repository.FindAllCategoryList()
	if err != nil {
		return nil, err
	}

	resp = make([]response.Get_Api_Category_List_Response, len(categories))
	for index, category := range categories {
		resp[index] = response.Get_Api_Category_List_Response{
			CategoryID: category.ID,
			Name:       category.Name,
		}
	}

	return resp, nil
}

func Get_Api_Category_Detail_Service(
	req *request.Get_Api_Category_Detail_Request,
	auth_user user.User,
) (*response.Get_Api_Category_Detail_Response, error) {
	var resp *response.Get_Api_Category_Detail_Response

	// 检查用户是否存在
	exist, err := repository.ExistsUser(auth_user.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	category, err := repository.FirstCategoryWithID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	resp = &response.Get_Api_Category_Detail_Response{
		Name:            category.Name,
		Description:     category.Description,
		CreatorID:       category.CreatorID,
		CreatorNickName: category.Creator.NickName,
	}

	return resp, nil
}
