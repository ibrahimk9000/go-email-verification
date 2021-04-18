package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mail/{mail}", emailHandler)
	r.HandleFunc("/verification/{verification}/", verifyHandler)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	mail := vars["mail"]

	hashmail := base64.StdEncoding.EncodeToString(hashfunc(mail))

	vs := false
	if exist(hashmail) {
		vs = checkValue(hashmail)
	} else {
		addKey(hashmail)
		err := verifcationLink(mail, r.Host)
		if err != nil {
			log.Println(err)
		}
	}

	json.NewEncoder(w).Encode(map[string]bool{"verified": vs})

}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	token := vars["verification"]

	mail, expirein, err := verifyToken(token)

	if err != nil {
		log.Println("error verify token", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "token invalid "})
		return
	}

	if expirein < time.Now().Unix() {
		json.NewEncoder(w).Encode(map[string]string{"error": "token expired "})
		return
	}

	hashmail := base64.StdEncoding.EncodeToString(hashfunc(mail))
	vs := false

	if exist(hashmail) {

		changeStatus(hashmail)
		vs = checkValue(hashmail)

	}

	json.NewEncoder(w).Encode(map[string]bool{"verified": vs})

}
