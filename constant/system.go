/*
   @TIME: 2023/6/12 14:32
   @Author: Nightz
   @Site:
   @File: system.go
   @SoftWare: GoLand
*/

package constant

type (
	MStatus uint8 // mqtt状态
	HStatus uint8 // http状态
	DConfig uint8 // 抓拍图片上传配置
)

// 自定义表名
const (
	TableNameUser          = "users"
	TableNameMqttConfig    = "mqtt_configs"
	TableNameHttpConfig    = "http_configs"
	TableNameNetWorkConfig = "network_configs"
	TableNameSystemInfo    = "system_infos"
)

// mqtt状态
const (
	MIsUp   MStatus = iota // 启用
	MIsDown                // 停用
)

// http状态
const (
	HIsUp   HStatus = iota // 启用
	HIsDown                // 停用
)

// 抓拍图片上传配置
const (
	FaceAndPlate     DConfig = iota // 人脸/车辆
	PeopleAndVehicle                // 行人/机动车/非机动车
	BackPhoto                       // 全景
)

/*

 */
