package app

import (
	"encoding/json"
	"github.com/mitriygor/usersProject/service"
	"net/http"
)

type UserHandlers struct {
	service service.UserService
}

func (ch *UserHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	users, err := ch.service.GetAllUsers(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
