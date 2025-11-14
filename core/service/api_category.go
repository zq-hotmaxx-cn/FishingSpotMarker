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
