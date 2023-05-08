package callback

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
	"vkPositiveBot/utils"
)

const (
	confirmationCode               = "14d2213e"
	TOKEN                          = "vk1.a.wUDWptIgLrdDUB4gOC8Qlo1S4xCp5FdPyX8gEhIQlWJwAMV-ntyvlXA8pMdO5HBnQMNM3rziPJafUVv9qAB_CznfcSpbbXjEmT1e-UMwVZqJQcIBhkHzxVWsLWwvzHWbD2GXN7wrCg-cbuZvgbABFDIZy2-bO6cwAI2GqJHQxMpkIAMIKqGxoj-Q9ZB8RIdRP3Cyk94uJWmDDm2-fme8cw"
	sendMessageAPI                 = "https://api.vk.com/method/messages.send"
	getCallbackConfirmationCodeAPI = "https://api.vk.com/method/groups.getCallbackConfirmationCode"
	EventAnswerAPI                 = "https://api.vk.com/method/messages.sendMessageEventAnswer"
	API_VERSION                    = "5.131"
	startMessage                   = "Привет &#9995;, я бот для поднятия настроения!\nВыбери одну из нескольких комманд на клавиатуре &#128519;"
	logUserID                      = 159083295
)

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	var m message
	var callback callback

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrMsg(err)
		return
	}
	r.Body.Close()

	err = json.Unmarshal(body, &callback)
	if err != nil {
		sendErrMsg(err)
		return
	}

	switch callback.Type {
	case "confirmation":
		//log.Println("confirmation", callback.GroupID)
		//confCode, err := callbackConfirmation(callback.GroupID)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		_, err = fmt.Fprintf(w, confirmationCode)
		if err != nil {
			sendErrMsg(err)
		}
		return

	case "message_new":
		_, err = fmt.Fprintf(w, "ok")
		if err != nil {
			sendErrMsg(err)
			return
		}

		m = callback.Object.Message

	case "group_join":
		_, err = fmt.Fprintf(w, "ok")
		if err != nil {
			sendErrMsg(err)
			return
		}

		m = message{
			PeerID: callback.Object.PeerID,
			FromID: callback.Object.UserID,
			Text:   "Начать",
		}

	case "message_event":
		_, err = fmt.Fprintf(w, "ok")
		if err != nil {
			sendErrMsg(err)
			return
		}

		ans := eventAnswer{
			EventID: callback.Object.EventID,
			UserID:  callback.Object.UserID,
			PeerID:  callback.Object.PeerID,
			EventData: eventData{
				Type: "show_snackbar",
				Text: eventAnswerText[rand.Intn(len(eventAnswerText))],
			},
		}

		ansData := setupURLEncoded(ans)

		resp, err := http.Post(APIRequest(EventAnswerAPI),
			"application/x-www-form-urlencoded", strings.NewReader(ansData))
		defer resp.Body.Close()

		if err != nil {
			sendErrMsg(err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			sendErrMsg(err)
			return
		}

		var vkResp vkResponse
		if err = json.Unmarshal(body, &vkResp); err != nil {
			sendErrMsg(err)
			return
		}

		if vkResp.Error.ErrorCode != 0 {
			sendErrMsg(fmt.Errorf("eventAnswer error: %d: %s\n %v",
				vkResp.Error.ErrorCode,
				vkResp.Error.ErrorMsg,
				vkResp.Error.RequestParams))
			return
		}

		m = message{
			PeerID: callback.Object.PeerID,
			FromID: callback.Object.UserID,
			Text:   callback.Object.Payload.Button,
		}
	}

	err = newMessage(m)
	if err != nil {
		sendErrMsg(err)
	}

}

func callbackConfirmation(groupID int64) (confCode string, err error) {
	callbackURL := callbackConfirmationReq(groupID)
	r, err := http.Get(callbackURL)
	defer r.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var response callbackConfirmationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	confCode = response.Response.Code
	return
}

func newMessage(m message) error {
	rand.Seed(time.Now().UnixNano())

	var data, keyboardKey string
	responseMessage := sendMessage{
		UserID:   m.FromID,
		RandomID: rand.Int31(),
		PeerID:   m.PeerID,
	}

	switch m.Text {
	case "ping":
		responseMessage.Message = "pong"

	case "Начать":
		responseMessage.Message = startMessage
		keyboardKey = "start"

	case "cat", "dog":
		imgURL, err := utils.GetRandomPicture(fmt.Sprintf("cute %s", m.Text))
		if err != nil {
			return err
		}
		responseMessage.Attachment = imgURL

	case "quote":
		quote, err := utils.GetRandomQuote()
		if err != nil {
			return err
		}
		responseMessage.Message = quote

	case "birthday", "new year":
		greeting, err := utils.GetRandomGreeting(m.Text)
		if err != nil {
			return err
		}
		responseMessage.Message = greeting

	default:
		if _, ok := messagesText[m.Text]; ok {
			responseMessage.Message = messagesText[m.Text]

			if _, ok := keyboards[m.Text]; ok {
				keyboardKey = m.Text
			}

		} else {
			responseMessage.Message = messagesText["help"]
			keyboardKey = "start"
		}
	}

	if keyboardKey != "" {
		data = setupURLEncoded(responseMessage, keyboards[keyboardKey])
	} else {
		data = setupURLEncoded(responseMessage)
	}

	sendMessageURL := APIRequest(sendMessageAPI)

	//log.Println(sendMessageURL)

	resp, err := http.Post(sendMessageURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var vkResp vkResponse
	if err = json.Unmarshal(body, &vkResp); err != nil {
		return err
	}

	if vkResp.Error.ErrorCode != 0 {
		return fmt.Errorf("sendMessage error: %d: %s\n %v",
			vkResp.Error.ErrorCode, vkResp.Error.ErrorMsg, vkResp.Error.RequestParams)
	}

	return err
}

func APIRequest(startURL string) string {
	return fmt.Sprintf("%s?access_token=%s&v=%s",
		startURL, TOKEN, API_VERSION)
}

func callbackConfirmationReq(groupID int64) string {
	return fmt.Sprintf("%s?access_token=%s&group_id=%d&v=%s",
		getCallbackConfirmationCodeAPI, TOKEN, groupID, API_VERSION)
}

func setupURLEncoded(encodings ...Encoding) string {
	data := url.Values{}
	for _, e := range encodings {
		temp := e.URLEncoded()
		for k, v := range temp {
			for _, vv := range v {
				data.Add(k, vv)
			}
		}
	}

	return data.Encode()
}

func sendErrMsg(err error) {
	errMessage := sendMessage{
		UserID:     logUserID,
		RandomID:   rand.Int31(),
		Message:    err.Error(),
		Attachment: "",
	}

	data := setupURLEncoded(errMessage)

	sendMessageURL := APIRequest(sendMessageAPI)

	_, err = http.Post(sendMessageURL, "application/x-www-form-urlencoded", strings.NewReader(data))

	if err != nil {
		log.Fatal(err)
	}
}
