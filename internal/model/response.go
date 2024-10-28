package model

type Response struct {
	Code uint64      `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
