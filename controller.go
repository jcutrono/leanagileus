package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/init", initData).Methods("GET")
	router.HandleFunc("/test/{name}", getData).Methods("GET")
}

func initData(resp http.ResponseWriter, req *http.Request) {

	insert()
}

func getData(resp http.ResponseWriter, req *http.Request) {

	name, _ := mux.Vars(req)["name"]
	data := callDb(name)

	if data.Name == "" {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	val, _ := json.Marshal(data)
	resp.Write(val)
}
