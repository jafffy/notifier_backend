package main

import (
	"net/http"
	"io"
	"log"
)

func getToday(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "운동화 들고 나가기!\n")
}

func main() {
	http.HandleFunc("/today", getToday)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
