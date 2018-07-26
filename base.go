package beecloudsdk

type BaseResponse struct {
	ResultCode int    `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
	ErrDetail  string `json:"err_detail"`
}
