package httpresponse

import "gobanking/internal/core-layer/domain"

type Response struct {
	Data interface{} `json:"data"`
	// client will check is and if its true it will continue else it will stop
	Status  bool `json:"status"`
	ErrCode domain.ErrCode `json:"error_code"`
	Msg domain.ErrMsg `json:"error_message"`
}