/*
   @TIME: 2023/6/12 12:21
   @Author: Nightz
   @Site:
   @File: device.go
   @SoftWare: GoLand
*/

package constant

type (
	AiStatus uint8 // 模型启用状态
	AiIsUsed uint8 // 模型使用状态
	AiSource uint8 // 模型来源

	DSource uint8 // 设备来源
	DType   uint8 // 设备类型
	DBrand  uint8 // 设备品牌
	DStatus uint8 // 设备状态
	DFormat uint8 // 设备流类型
)

// 自定义表名
const (
	TableNameAlgorithm = "algorithms"
	TableNameDevice    = "devices"
)

// 模型启用状态
const (
	AiIsUp   AiStatus = iota // 启用
	AiIsDown                 // 停用
)

// 模型使用状态
const (
	AiIsUse    AiIsUsed = iota // 使用
	AiIsNotUse                 // 未使用
)

// 模型来源
const (
	AiIsSourceByImport   AiSource = iota // 第三方导入
	AiIsSourceByPlatForm                 // 平台新增
)

// 设备来源
const (
	DeviceIsSourceByImport   DSource = iota // 第三方导入
	DeviceIsSourceByPlatForm                // 平台新增
)

// 设备类型
const (
	DeviceIsAi      DType = iota // 智能摄像机
	DeviceIsNetWork              // 网络摄像机
)

// 设备品牌
const (
	DeviceByHaiKang DBrand = iota // 海康
	DeviceByDaHua                 // 大华
	DeviceByHuaWei                // 华为
	DeviceByYuShi                 // 宇视
	DeviceByOther                 // 其它
)

// 设备状态
const (
	DeviceIsUp    DStatus = iota // 在线
	DeviceIsDown                 // 离线
	DeviceIsError                // 异常
)

// 设备流类型
const (
	DeviceIsRtsp DFormat = iota
	DeviceIsOnvif
	DeviceIsGb28181
	DeviceIsGat1400
)

/*

 */
