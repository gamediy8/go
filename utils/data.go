package utils

type Data struct {
	Result interface{} `json:"result"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
}
