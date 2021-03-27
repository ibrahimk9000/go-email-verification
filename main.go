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
	r.HandleFunc("/mail/{mail}", emailfunc)
	r.HandleFunc("/verification/{verification}/", verifyfunc)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func emailfunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	mail := vars["mail"]

	hashmail := base64.StdEncoding.EncodeToString(hashfunc(mail))

	vs := false
	if exist(hashmail) {
		vs = checkvalue(hashmail)
	} else {
		addkey(hashmail)
		err := verifcationlink(mail, r.Host)
		if err != nil {
			log.Println(err)
		}
	}

	json.NewEncoder(w).Encode(map[string]bool{"verified": vs})

}

func verifyfunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	token := vars["verification"]

	mail, expirein, err := verifytoken(token)

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

		changestatus(hashmail)
		vs = checkvalue(hashmail)

	}

	json.NewEncoder(w).Encode(map[string]bool{"verified": vs})

}
