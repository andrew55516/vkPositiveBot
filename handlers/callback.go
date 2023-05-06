package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	TOKEN                          = "vk1.a.wUDWptIgLrdDUB4gOC8Qlo1S4xCp5FdPyX8gEhIQlWJwAMV-ntyvlXA8pMdO5HBnQMNM3rziPJafUVv9qAB_CznfcSpbbXjEmT1e-UMwVZqJQcIBhkHzxVWsLWwvzHWbD2GXN7wrCg-cbuZvgbABFDIZy2-bO6cwAI2GqJHQxMpkIAMIKqGxoj-Q9ZB8RIdRP3Cyk94uJWmDDm2-fme8cw"
	sendMessageURL                 = "https://api.vk.com/method/messages.send"
	getCallbackConfirmationCodeURL = "https://api.vk.com/method/groups.getCallbackConfirmationCode"
	API_VERSION                    = "5.131"
	ConfirmationCode               = "39e324a8"
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Body.Close()

	var callback callback
	err = json.Unmarshal(body, &callback)
	if err != nil {
		log.Fatal(err)
		return
	}

	switch callback.Type {
	case "confirmation":
		fmt.Println("confirmation", callback.GroupID)
		confCode, err := callbackConfirmation(callback.GroupID)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = fmt.Fprintf(w, confCode)
		if err != nil {
			log.Fatal(err)
			return
		}

	case "message_new":
		_, err = fmt.Fprintf(w, "ok")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = newMessage(callback.Object.Message)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

}

func callbackConfirmation(groupID int64) (confCode string, err error) {
	//c := callbackConfirmationRequest{
	//	AccessToken: TOKEN,
	//	GroupID:       groupID,
	//	V:             API_VERSION,
	//}

	//data, err := json.Marshal(c)
	//if err != nil {
	//	return
	//}

	//r, err := http.Post(getCallbackConfirmationCodeURL, "application/json", bytes.NewBuffer(data))
	url := fmt.Sprintf("%s?access_token=%s&group_id=%d&v=%s",
		getCallbackConfirmationCodeURL, TOKEN, groupID, API_VERSION)
	fmt.Println(url)
	r, err := http.Get(url)
	defer r.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))

	var response callbackConfirmationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	log.Println(response)

	confCode = response.Response.Code
	return
}

func newMessage(m message) error {
	rand.Seed(time.Now().UnixNano())

	var response sendMessage
	switch m.Text {
	case "ping":
		response = sendMessage{
			//AccessToken: TOKEN,
			UserID:   m.FromId,
			RandomID: rand.Int31(),
			PeerId:   m.PeerId,
			Message:  "pong",
			//V:           API_VERSION,
		}

	default:
		response = sendMessage{
			//AccessToken: TOKEN,
			UserID:   m.FromId,
			RandomID: rand.Int31(),
			PeerId:   m.PeerId,
			Message:  "Hi!",
			//V:           API_VERSION,
		}
	}

	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?access_token=%s&random_id=%d&user_id=%d&peer_id=%d&message=%s&v=%s",
		sendMessageURL, TOKEN, rand.Int31(), m.FromId, m.PeerId, response.Message, API_VERSION)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return err
}
