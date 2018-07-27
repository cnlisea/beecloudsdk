package beecloudsdk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type OrderQueryResponse struct {
	BaseResponse
	Bills []*OrderQueryResponseBill `json:"bills"`
}

type OrderQueryResponseBill struct {
	SpayResult  bool   `json:"spay_result"`  // 订单是否成功
	TradeNo     string `json:"trade_no"`     // 渠道交易号
	TotalFee    int64  `json:"total_fee"`    // 金额
	SuccessTime int64  `json:"success_time"` // 订单支付成功时间
}

func OrderQuery(cfg *Config) (*OrderQueryResponseBill, error) {
	timestamp := time.Now().UnixNano() / 1e6
	param := map[string]interface{}{
		"app_id":    cfg.AppId,
		"timestamp": timestamp,
		"app_sign":  fmt.Sprintf("%x", string(EncryptMD5([]byte(cfg.AppId+strconv.FormatInt(timestamp, 10)+cfg.AppSecret)))),
		"bill_no":   cfg.BillNo, // 订单号
	}

	reqBody, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}

	res, err := HttpSendGet("https://api.beecloud.cn/2/rest/bills", reqBody)
	if err != nil {
		return nil, err
	}

	var response OrderQueryResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	if response.ResultCode != 0 {
		return nil, errors.New(response.ResultMsg)
	}

	if len(response.Bills) == 0 {
		return nil, nil
	}

	return response.Bills[0], nil
}
