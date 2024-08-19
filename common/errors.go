package common

import "fmt"

type APIError struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("<APIError>\ncode = %d \nmessage = %s \n", e.ErrorCode, e.Description)
}

func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
