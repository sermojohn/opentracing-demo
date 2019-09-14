package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func setResponseBody(w http.ResponseWriter, s string) {
	w.Write([]byte(fmt.Sprintf("replying from app1\n%s", s)))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	s, err := callUpstream()
	if err != nil {
		w.WriteHeader(500)
		setResponseBody(w, "error")
		return
	}

	w.WriteHeader(200)
	setResponseBody(w, s)
}

func callUpstream() (string, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://localhost:8082/")
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
