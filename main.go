package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type errorCode struct {
	Code  int
	Error string
}

func IsPanic(w http.ResponseWriter, r *http.Request) {

	example := r.URL.Query().Get("example")

	if example == "err" {
		log.Panic("panic")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result" :"success"})
	return
}

func recover_(w http.ResponseWriter, r *http.Request) {
	re := recover()
	fmt.Println(re)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"error" :"error msg"})
	return
		
	
}

func IsRecover(w http.ResponseWriter, r *http.Request) {

	example := r.URL.Query().Get("example")

	defer recover_(w,r)

	if example == "err" {
		log.Panic("panic")
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result" :"success"})
	return
}

func IsErrorHandler(w http.ResponseWriter, r *http.Request) {
	example := r.URL.Query().Get("example")

	if example == "err" {
		ErrorHandler(w,r,0)
		return 
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result" :"success"})
	return
}


func ErrorHandler(w http.ResponseWriter, r *http.Request, error_num int){

	error_json, err := os.Open("./error_code.json")
	if err != nil {
		fmt.Println(err)
	}
	defer error_json.Close()
	var code []errorCode
	getValue, _ := ioutil.ReadAll(error_json)
	json.Unmarshal(getValue, &code)
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{fmt.Sprintf("%s", "error") : fmt.Sprintf("%s", code[error_num].Error)})
		return
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/panic", IsPanic).Methods("GET")
	r.HandleFunc("/recover", IsRecover).Methods("GET")
	r.HandleFunc("/errorhandler", IsErrorHandler).Methods("GET")

	http.ListenAndServe(":8081", r)
}