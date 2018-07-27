package beecloudsdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"unsafe"
)

func HttpSendPost(url string, body []byte) ([]byte, error) {
	res, err := http.Post(url, "application/json;charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func HttpSendGet(u string, param []byte) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(u)
	b.WriteString("?para=")
	b.WriteString(url.QueryEscape(*(*string)(unsafe.Pointer(&param))))

	res, err := http.Get(b.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
