package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vkPositiveBot/callback"
)

func main() {
	http.HandleFunc("/callback", callback.HandleCallback)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	APP_IP := os.Getenv("APP_IP")
	APP_PORT := os.Getenv("APP_PORT")

	fmt.Println(APP_IP + ":" + APP_PORT)
	log.Fatal(http.ListenAndServe(APP_IP+":"+APP_PORT, nil))
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
