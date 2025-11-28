# API

## 通用请求头
* X-Request-Id：请求ID

## 健康检查

### 1. API状态
* 路由：`/api/health`
* 请求方式：`GET`

## 分类

### 1. 创建分类
* 路由：`/api/category`
* 请求方式：`POST`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * JSON请求体：
    * `string` name: 分类名称
    * `string` description: 分类描述
* 返回参数：
  * JSON响应体：
    * `uint` category_id: 分类ID

### 2. 获取所有分类列表
* 路由：`/api/category/list`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 返回参数：
  * JSON响应体：
    * `uint` category_id: 分类ID
    * `string` name: 分类名称

### 3. 获取分类详情
* 路由：`/api/category/detail`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * `uint` category_id: 分类ID
* 返回参数：
  * JSON响应体：
    * `string` name: 分类名称
    * `string` description: 分类描述
    * `uint` creator_id: 创建者ID
    * `string` creator_nick_name: 创建者昵称

## 钓点

### 1. 标记钓点
* 路由：`/api/fishing_spot/mark`
* 请求方式：`POST`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * JSON请求体：
    * `float64` longitude: 经度
    * `float64` latitude: 纬度
    * `bool` car_access: 是否允许车辆
    * `string` location_name: 位置名称
    * `string` location: 位置
    * `string` description: 描述
    * `[]uint` category_ids: 分类ID列表
* 返回参数：
  * JSON响应体：
    * `uint` fishing_spot_marker_id: 钓鱼点标记ID

### 2. 获取指定矩形范围钓点列表
* 路由：`/api/fishing_spot/rect-list`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * `float64` min_latitude: 最小纬度
    * `float64` min_longitude: 最小经度
    * `float64` max_latitude: 最大纬度
    * `float64` max_longitude: 最大经度
    * `bool` car_access: 是否允许车辆
* 返回参数：
  * JSON响应体-列表：
    * `uint` fishing_spot_marker_id: 钓鱼点标记ID
    * `float64` longitude: 经度
    * `float64` latitude: 纬度
    * `bool` car_access: 是否允许车辆通行
    * `string` location_name: 位置名
    * `uint` creator_id: 创建者ID
    * `string` creator_nick_name: 创建者昵称
    * `[]uint` category_ids: 分类ID列表

### 3. 获取指定钓点详情
* 路由：`/api/fishing_spot/detail`
* 请求方式：`GET`
* 请求头：
  * Authorization: `Bearer <token>`
* 请求参数：
  * Query参数：
    * `uint` fishing_spot_marker_id: 钓鱼点标记ID
* 返回参数：
  * JSON响应体：
    * `float64` longitude: 经度
    * `float64` latitude: 纬度
    * `bool` car_access: 是否允许车辆
    * `string` location_name: 位置名称
    * `string` location: 位置
    * `string` description: 描述
    * `uint` creator_id: 创建者ID
    * `string` creator_nick_name: 创建者昵称
    * `[]uint` category_ids: 分类ID列表
