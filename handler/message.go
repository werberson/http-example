package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"http-example/crypt"
	"http-example/session"
	"net/http"
)

func MessageHandler(writer http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	if name == "" {
		http.Error(writer, "Name is required on QueryString. e.g /new/john", http.StatusBadRequest)
		return
	}

	sessionData, present := session.Get(name)

	if !present {
		http.Error(writer, "Call /new/{name} first.", http.StatusBadRequest)
		return
	}
	var requestBody map[string]string
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(writer, "I can't parse this fucking body.", http.StatusInternalServerError)
		return
	}
	message := requestBody["message"]
	if message == "" {
		http.Error(writer, `Message is required on body. e.g {"message" : "b531087ba0cd3f0832b8dde56a7892f394c73241bf5d6757a4f36eb75276cafa514fabecc1bd"}`, http.StatusBadRequest)
		return
	}

	fmt.Println(message)

	key := md5.Sum([]byte(sessionData.ServerKey))
	decrypted, err := crypt.Decrypt([]byte(message), key[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x => %s\n", decrypted, message)

	key = md5.Sum([]byte(sessionData.ClientKey))
	encrypted, err := crypt.Encrypt(decrypted, key[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x => %s\n", encrypted, message)

	json.NewEncoder(writer).Encode(map[string]string{"message": message})
}
