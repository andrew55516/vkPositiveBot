package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const ()

func main() {
	http.HandleFunc("/callback", hadleCallback)
	log.Fatal(http.ListenAndServe(":5986", nil))
}

func hadleCallback(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var conf struct {
		Type    string `json:"type"`
		GroupId int    `json:"group_id"`
	}

	err = json.Unmarshal(body, &conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	if conf.Type == "confirmation" && conf.GroupId == 220370106 {
		_, err = w.Write([]byte("e048fbdb"))
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	return
}
