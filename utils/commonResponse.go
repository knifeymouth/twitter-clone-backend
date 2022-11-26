package utils

import (
	"encoding/json"
	"net/http"
	"twitter-clone/structs"
)

func CommonResponse(w http.ResponseWriter, code int, msg string, data interface{}) {
	json.NewEncoder(w).Encode(structs.CommonResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
