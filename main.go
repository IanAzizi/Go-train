package main

import (
    "bace-resp/helper"
    "net/http"
)

func main() {
    server := http.Server{}
    mux := http.NewServeMux()

    server.Handler = mux
    mux.HandleFunc("/GetUser", GetUser)
    server.Addr = ":8080"
    server.ListenAndServe()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    helper.SuccessBaceResponce(w, "working", true, http.StatusOK)
}