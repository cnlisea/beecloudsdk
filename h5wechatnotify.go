package beecloudsdk

import (
	"encoding/json"
	"errors"
)

type H5WeChatNotifyRequest struct {
	Signature       string `json:"signature"`        // 签名
	Timestamp       int64  `json:"timestamp"`        // 时间毫秒
	ChannelType     string `json:"channel_type"`     // 取道类型
	TransactionId   string `json:"transaction_id"`   // 交易单号
	TransactionType string `json:"transaction_type"` // 交易类型
	TransactionFee  int64  `json:"transaction_fee"`  // 实付金额
	BillFee         int64  `json:"bill_fee"`         // 订单金额
	TradeSuccess    bool   `json:"trade_success"`    // 交易是否成功
}

func H5WeChatNotifyAnalysis(cfg *Config) (*H5WeChatNotifyRequest, error) {
	var (
		req = new(H5WeChatNotifyRequest)
		err error
	)

	if err = json.NewDecoder(cfg.NotifyBody).Decode(req); err != nil {
		return nil, err
	}

	if req.Signature != VerifySignNotify(cfg.AppId, req.TransactionId, req.TransactionType, req.ChannelType, req.TransactionFee, cfg.MasterSecret) {
		return nil, errors.New("signature not match")
	}

	return req, nil
}
