package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"` // is tag se yeh hoga ki jb json me convert hoga to small letter me aayga
	Error  string `json:"error"`
}

const (
	status      = "OK"
	statusError = "Error"
)

func WrtieJson(w http.ResponseWriter, status int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: statusError,
		Error:  err.Error(),
	}
}
