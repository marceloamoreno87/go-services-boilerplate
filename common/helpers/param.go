package helpers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetParam(r *http.Request, value string) (ID int) {
	idStr := chi.URLParam(r, value)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return id
}

func GetFormString(r *http.Request, value string) (ID string) {
	idStr := r.FormValue(value)
	return idStr
}
