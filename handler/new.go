package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/werberson/http-example/session"
	"net/http"
)

const serverName = "werberson"

func KeyHandler(writer http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	if name == "" {
		http.Error(writer, "Name is required on QueryString. e.g /message/john", http.StatusBadRequest)
		return
	}

	var requestBody map[string]string
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(writer, `I can't parse this fucking body. You must send a body like {"key" : "d2f45abf-12c6-4cfa-8c39-aad5efb4cbf0"}`, http.StatusInternalServerError)
		return
	}
	if requestBody["key"] == "" {
		http.Error(writer, `Key is required on body. e.g {"key" : "d2f45abf-12c6-4cfa-8c39-aad5efb4cbf0"}`, http.StatusBadRequest)
		return
	}

	createdSession := session.Create(name, requestBody["key"])
	json.NewEncoder(writer).Encode(map[string]string{"key": createdSession.ServerKey, "name": serverName})
}
