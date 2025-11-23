package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Rishabhsingh78/student-api/internal/types"
	"github.com/Rishabhsingh78/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WrtieJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WrtieJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if r.Body != nil {
			response.WrtieJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		w.Write([]byte("Welcome to student api"))
	}
}
