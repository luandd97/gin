package helpers

import "strings"

type Respone struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

type UserID struct {
	UserID uint64 `json:"user_id"`
}

func BuildResponse(status bool, message string, data interface{}) Respone {
	res := Respone{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, error string, data interface{}) Respone {
	splittedError := strings.Split(error, "\n")
	res := Respone{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
