package pkg

import (
	"encoding/json"
	"net/http"
	"time"
)

type ApiResp struct {
	StatusCode int         `json:"-"`
	OK         bool        `json:"ok"`
	Data       interface{} `json:"data,omitempty"`
	ErrCode    string      `json:"err,omitempty"`
	Message    string      `json:"msg,omitempty"`
	Timestamp  int64       `json:"ts"`
}

func NewSuccessResp(data interface{}) ApiResp {
	return ApiResp{
		StatusCode: http.StatusOK,
		OK:         true,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

func WriteAPIResp(w http.ResponseWriter, resp ApiResp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	b, _ := json.Marshal(resp)
	w.Write(b)
}
