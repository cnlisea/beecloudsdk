package beecloudsdk

import (
	"bytes"
	"fmt"
	"strconv"
)

func VerifySignNotify(appId string, transactionId string, transactionType string, channelType string, transactionFee int64, masterSecret string) string {
	var b bytes.Buffer
	b.WriteString(appId)
	b.WriteString(transactionId)
	b.WriteString(transactionType)
	b.WriteString(channelType)
	b.WriteString(strconv.FormatInt(transactionFee, 10))
	b.WriteString(masterSecret)
	return fmt.Sprintf("%x", string(EncryptMD5(b.Bytes())))
}
