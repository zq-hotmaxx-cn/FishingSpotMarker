# API

## 通用请求头
* X-Request-Id：请求ID

## 分类

### 1. 创建分类
* 路由：`/api/category`
* 请求方式：`POST`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * JSON请求体：
    * name: 分类名称
    * description: 分类描述
* 返回参数：
  * JSON响应体：
    * category_id: 分类ID

### 2. 获取所有分类列表
* 路由：`/api/category/list`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 返回参数：
  * JSON响应体：
    * category_id: 分类ID
    * name: 分类名称

### 3. 获取分类详情
* 路由：`/api/category/detail`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * category_id: 分类ID
* 返回参数：
  * JSON响应体：
    * name: 分类名称
    * description: 分类描述
    * creator_id: 创建者ID
    * creator_nick_name: 创建者昵称

## 钓点

### 1. 标记钓点
* 路由：`/api/fishing_spot/mark`
* 请求方式：`POST`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * JSON请求体：
    * longitude: 经度
    * latitude: 纬度
    * car_access: 是否允许车辆
    * location_name: 位置名称
    * location: 位置
    * description: 描述
    * category_ids: 分类ID列表
* 返回参数：
  * JSON响应体：
    * fishing_spot_marker_id: 钓鱼点标记ID

### 2. 获取指定矩形范围钓点列表
* 路由：`/api/fishing_spot/rect-list`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * min_latitude: 最小纬度
    * min_longitude: 最小经度
    * max_latitude: 最大纬度
    * max_longitude: 最大经度
    * car_access: 是否允许车辆
* 返回参数：
  * JSON响应体-列表：
    * fishing_spot_marker_id: 钓鱼点标记ID
    * longitude: 经度
    * latitude: 纬度
    * car_access: 是否允许车辆通行
    * location_name: 位置名
    * creator_id: 创建者ID
    * creator_nick_name: 创建者昵称
    * category_ids: 分类ID列表

### 3. 获取指定钓点详情
* 路由：`/api/fishing_spot/detail`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * fishing_spot_marker_id: 钓鱼点标记ID
* 返回参数：
  * JSON响应体：
    * longitude: 经度
    * latitude: 纬度
    * car_access: 是否允许车辆
    * location_name: 位置名称
    * location: 位置
    * description: 描述
    * creator_id: 创建者ID
    * creator_nick_name: 创建者昵称
    * category_ids: 分类ID列表
