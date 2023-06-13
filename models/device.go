/*
   @TIME: 2023/6/12 11:07
   @Author: Nightz
   @Site:
   @File: device.go
   @SoftWare: GoLand
*/

package model

import (
	"database/sql"

	"yishi-ai.com/go-sm/constant"
)

type Algorithm struct {
	BaseModel
	AiName           string            `json:"ai_name" gorm:"type:varchar(8);not null;comment:模型名称"`
	AiStatus         constant.AiStatus `json:"ai_status" gorm:"size:2;default:0;comment:模型状态"`
	AiIsUsed         constant.AiIsUsed `json:"ai_is_used" gorm:"size:2;default:0;comment:是否使用"`
	AiSource         constant.AiSource `json:"ai_source" gorm:"size:2;default:0;comment:模型来源"`
	AiDescribe       string            `json:"ai_describe" gorm:"type:longtext;comment:模型描述"`
	AiFile           string            `json:"ai_file" gorm:"type:varchar(64);not null;comment:模型文件"`
	AiSourceId       string            `json:"ai_source_id" gorm:"comment:来源id"`
	Algorithm2Device []Device          `gorm:"many2many:device_algorithm"`
}

func (a *Algorithm) TableNme() string {
	return constant.TableNameAlgorithm
}

type Device struct {
	BaseModel
	DName            string           `json:"d_name" gorm:"type:varchar(8);not null;comment:设备名称"`
	DIp              string           `json:"d_ip" gorm:"type:varchar(16);not null;comment:设备ip"`
	DAddress         string           `json:"d_address" gorm:"type:varchar(128);comment:设备地址"`
	DGeo             string           `json:"d_geo" gorm:"type:varchar(64);comment:设备经纬度"`
	DSource          constant.DSource `json:"d_source" gorm:"size:2;default:0;comment:设备来源"`
	DType            constant.DType   `json:"d_type" gorm:"size:2;default:0;comment:设备类型"`
	DBrand           constant.DBrand  `json:"d_brand" gorm:"size:2;default:0;comment:设备品牌"`
	DStatus          constant.DStatus `json:"d_status" gorm:"size:2;default:0;comment:设备状态"`
	DFormat          constant.DFormat `json:"d_format" gorm:"size:2;default:0;comment:视频流类型"`
	DUserName        string           `json:"d_username" gorm:"type:varchar(8);not null;comment:设备用户名"`
	DPassWord        string           `json:"d_password" gorm:"type:varchar(16);not null;comment:设备密码"`
	DRtsp            string           `json:"d_rtsp" gorm:"type:varchar(64);comment:设备rtsp地址"`
	DChannel         uint8            `json:"d_channel" gorm:"comment:通道号"`
	DLastLogin       *sql.NullTime
	DLastLogout      *sql.NullTime
	DSourceId        string      `json:"d_source_id" gorm:"type:varchar(32);comment:来源id"`
	Device2Algorithm []Algorithm `gorm:"many2many:device_algorithm"`
}

func (d *Device) TableName() string {
	return constant.TableNameDevice
}

/*

 */
