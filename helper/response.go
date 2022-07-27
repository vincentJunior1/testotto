package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
