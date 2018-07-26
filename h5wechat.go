package beecloudsdk

import (
	"encoding/json"
	"strconv"
	"time"

	"fmt"

	"github.com/pkg/errors"
)

type H5WechatPayResponse struct {
	BaseResponse
	Id  string `json:"id"`
	Url string `json:"url"` // 确认支付页面url
}

// 微信H5支付
func H5WechatPay(cfg *Config) (string, error) {
	timestamp := time.Now().Unix()
	param := map[string]interface{}{
		"app_id":     cfg.AppId,
		"timestamp":  timestamp,
		"app_sign":   fmt.Sprintf("%x", string(EncryptMD5([]byte(cfg.AppId+strconv.FormatInt(timestamp, 10)+cfg.AppSecret)))),
		"channel":    "BC_WX_WAP",
		"total_fee":  cfg.Amount,
		"bill_no":    cfg.BillNo,
		"title":      cfg.Title,
		"return_url": cfg.ReturnUrl,
		"notify_url": cfg.NotifyUrl,
		"analysis":   cfg.analysis,
	}

	reqBody, err := json.Marshal(&param)
	if err != nil {
		return "", err
	}

	res, err := HttpSendPost("https://api.beecloud.cn/2/rest/bill", reqBody)
	if err != nil {
		return "", err
	}

	var response H5WechatPayResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return "", err
	}

	if response.ResultCode != 0 {
		fmt.Println(response.ErrDetail)
		return "", errors.New(response.ResultMsg)
	}

	return response.Url, nil
}
