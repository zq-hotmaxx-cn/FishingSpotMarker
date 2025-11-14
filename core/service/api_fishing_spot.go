package service

import (
	"FishingSpotMarker/core/data"
	"FishingSpotMarker/core/repository"
	"FishingSpotMarker/pkg/http/request"
	"FishingSpotMarker/pkg/http/response"
	"FishingSpotMarker/pkg/model"
	"FishingSpotMarker/pkg/utils/user"

	"gorm.io/gorm"
)

func Post_Api_FishingSpot_Mark_Service(
	req *request.Post_Api_FishingSpot_Mark_Request,
	auth_user user.User,
) (*response.Post_Api_FishingSpot_Mark_Response, error) {
	var resp *response.Post_Api_FishingSpot_Mark_Response

	// 检查用户是否存在
	exist, err := repository.ExistsUser(auth_user.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	// 根据分类ID列表获取已存在的分类列表
	categories, err := repository.FindCategoryListByIDs(req.CategoryIDs)
	if err != nil {
		return nil, err
	}

	// 创建钓点标记
	fishing_spot_marker := &model.FishingSpotMarker{
		CreatorID:    auth_user.ID,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		CarAccess:    req.CarAccess,
		LocationName: req.LocationName,
		Location:     req.Location,
		Description:  req.Description,
		Categories:   categories,
	}

	// 开启事务
	if err := data.GetDataBase().Transaction(func(tx *gorm.DB) error {
		// 创建钓点标记
		if err := repository.CreateFishingSpotMarkerWithTransaction(tx, fishing_spot_marker); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// 响应数据
	resp = &response.Post_Api_FishingSpot_Mark_Response{
		FishingSpotMarkerID: fishing_spot_marker.ID,
	}

	return resp, nil
}

func Get_Api_FishingSpot_RectList_Service(
	req *request.Get_Api_FishingSpot_RectList_Request,
	auth_user user.User,
) ([]response.Get_Api_FishingSpot_RectList_Response, error) {
	var resp []response.Get_Api_FishingSpot_RectList_Response

	// 检查用户是否存在
	exist, err := repository.ExistsUser(auth_user.ID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	// 通过经纬度矩形获取范围内所有的钓点标记
	fishing_spot_markers, err := repository.FindFishingSpotMarkerListByRect(
		req.MinLatitude,
		req.MinLongitude,
		req.MaxLatitude,
		req.MaxLongitude,
	)

	// 响应数据
	resp = make([]response.Get_Api_FishingSpot_RectList_Response, len(fishing_spot_markers))
	for index, fishing_spot_marker := range fishing_spot_markers {
		// 获取分类ID列表
		category_ids := make([]uint, len(fishing_spot_marker.Categories))
		for index, category := range fishing_spot_marker.Categories {
			category_ids[index] = category.ID
		}

		// 响应数据
		resp[index] = response.Get_Api_FishingSpot_RectList_Response{
			FishingSpotMarkerID: fishing_spot_marker.ID,
			Latitude:            fishing_spot_marker.Latitude,
			Longitude:           fishing_spot_marker.Longitude,
			CarAccess:           fishing_spot_marker.CarAccess,
			LocationName:        fishing_spot_marker.LocationName,
			CreatorID:           fishing_spot_marker.CreatorID,
			CreatorNickName:     fishing_spot_marker.User.NickName,
			CategoryIDs:         category_ids,
		}
	}

	return resp, nil
}
