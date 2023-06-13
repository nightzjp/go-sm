/*
   @TIME: 2023/6/12 11:07
   @Author: Nightz
   @Site:
   @File: system.go
   @SoftWare: GoLand
*/

package model

import (
	"time"
	"yishi-ai.com/go-sm/constant"
)

type User struct {
	BaseModel
	UserName string `json:"username" gorm:"type:varchar(8);not null;comment:用户名"`
	PassWord string `json:"password" gorm:"type:varchar(16);not null;comment:密码"`
}

func (u *User) TableName() string {
	return constant.TableNameUser
}

type MqttConfig struct {
	BaseModel
	MStatus   *constant.MStatus `json:"m_status" gorm:"size:2;default:0;comment:是否启用"`
	MClientId string            `json:"m_client_id" gorm:"type:varchar(32);not null;comment:客户端id"`
	MIp       *string           `json:"m_ip" gorm:"type:varchar(16);comment:ip地址"`
	MPort     *uint16           `json:"m_port" gorm:"comment:端口号"`
	MUserName *string           `json:"m_username" gorm:"type:varchar(32);comment:用户名"`
	MPassWord *string           `json:"m_password" gorm:"type:varchar(32);comment:密码"`
}

func (m *MqttConfig) TableName() string {
	return constant.TableNameMqttConfig
}

type HttpConfig struct {
	BaseModel
	HStatus   *constant.HStatus `json:"h_status" gorm:"size:2;default:0;comment:是否启用"`
	HIp       *string           `json:"h_ip" gorm:"type:varchar(16);comment:ip地址"`
	HPort     *uint16           `json:"h_port" gorm:"comment:端口号"`
	HUserName *string           `json:"h_username" gorm:"type:varchar(32);comment:用户名"`
	HPassWord *string           `json:"h_password" gorm:"type:varchar(32);comment:密码"`
	HPeriod   *time.Duration    `json:"h_period" gorm:"comment:同步周期"`
}

func (h *HttpConfig) TableName() string {
	return constant.TableNameHttpConfig
}

type NetWorkConfig struct {
	BaseModel
	NName      string `json:"n_name" gorm:"type:varchar(32);comment:网卡名称"`
	NIp        string `json:"n_ip" gorm:"type:varchar(16);comment:ip地址"`
	NMask      string `json:"n_mask" gorm:"type:varchar(16);comment:子网掩码"`
	NBroadCast string `json:"n_broadcast" gorm:"type:varchar(16);comment:广播地址"`
	NGateWay   string `json:"n_gateway" gorm:"type:varchar(16);comment:网关"`
	NDns       string `json:"n_dns" gorm:"type:varchar(16);comment:dns配置"`
}

func (n *NetWorkConfig) TableName() string {
	return constant.TableNameNetWorkConfig
}

type SystemInfo struct {
	BaseModel
	DType             string            `json:"d_type" gorm:"comment:设备型号"`
	DId               string            `json:"d_id" gorm:"comment:设备id"`
	DMac              string            `json:"d_mac" gorm:"comment:Mac地址"`
	DFirmWareVersion  *string           `json:"d_firmware_version" gorm:"comment:固件版本"`
	DMachineVersion   *string           `json:"d_machine_version" gorm:"comment:一体机版本"`
	DAlgorithmVersion *string           `json:"d_algorithm_version" gorm:"comment:算法软件版本"`
	DFfmpegVersion    *string           `json:"d_ffmpeg_version" gorm:"comment:拉流软件版本"`
	DConfig           *constant.DConfig `json:"d_config" gorm:"comment:抓拍图片上传配置"`
}

func (s *SystemInfo) TableName() string {
	return constant.TableNameSystemInfo
}

/*

 */
