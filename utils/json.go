package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReadJSON(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	defer r.Body.Close()
	return err
}

func ResponseJSON(w http.ResponseWriter, code int, message string, payload interface{}) error {
	response := Response{
		Code:    code,
		Message: message,
		Data:    payload,
	}

	result, err := json.Marshal(response)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(result)
	return err
}

func ResponseError(w http.ResponseWriter, code int, errorMessage string) error {
	response := map[string]interface{}{
		"Message": errorMessage,
	}

	result, err := json.Marshal(response)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(result)
	return err
}
