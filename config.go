package beecloudsdk

import "io"

type Config struct {
	AppId        string                 // beeCloud 平台应用唯一标识
	AppSecret    string                 // beeCloud 平台应用密钥
	MasterSecret string                 // beeCloud平台密钥
	Amount       int64                  // 金额，单位分
	BillNo       string                 // 商户订单号 length: [8-20]
	Title        string                 // 订单标题 32 byte
	ReturnUrl    string                 // 同步返回地址
	NotifyUrl    string                 // 异步通知地址
	Analysis     map[string]interface{} // 附加数据

	NotifyBody io.ReadCloser // 异步通知数据
}

const (
	ConfigAnalysisSceneInfo = "scene_info"
	ConfigAnalysisIp        = "ip"
)
