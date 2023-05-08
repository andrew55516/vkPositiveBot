package callback

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Encoding interface {
	URLEncoded() url.Values
}

type vkResponse struct {
	Response int `json:"response"`
	Error    struct {
		ErrorCode     int         `json:"error_code"`
		ErrorMsg      string      `json:"error_msg"`
		RequestParams interface{} `json:"request_params"`
	} `json:"error"`
}

type callback struct {
	GroupID int64  `json:"group_id"`
	Type    string `json:"type"`
	Object  object `json:"object"`
	EventID string `json:"event_id"`
}

type object struct {
	Message message `json:"message"`
	UserID  int64   `json:"user_id"`
	PeerID  int64   `json:"peer_id"`
	EventID string  `json:"event_id"`
	Payload payload `json:"payload"`
}

type payload struct {
	Button string `json:"button"`
}

type message struct {
	PeerID int64  `json:"peer_id"`
	FromID int64  `json:"from_id"`
	Text   string `json:"text"`
}

type sendMessage struct {
	UserID     int64  `json:"user_id"`
	RandomID   int32  `json:"random_id"`
	PeerID     int64  `json:"peer_id"`
	Message    string `json:"message"`
	Attachment string `json:"attachment"`
}

type callbackConfirmationResponse struct {
	Response response `json:"response"`
}

type response struct {
	Code string `json:"code"`
}

type keyboard struct {
	Inline  bool       `json:"inline"`
	OneTime bool       `json:"one_time"`
	Buttons [][]button `json:"buttons"`
}

type button struct {
	Action action `json:"action"`
	Color  string `json:"color,omitempty"`
}

type action struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Link    string `json:"link,omitempty"`
	Payload string `json:"payload"`
}

type eventAnswer struct {
	EventID   string    `json:"event_id"`
	UserID    int64     `json:"user_id"`
	PeerID    int64     `json:"peer_id"`
	EventData eventData `json:"event_data"`
}

type eventData struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (m sendMessage) URLEncoded() url.Values {
	data := url.Values{}
	if m.Message != "" {
		data.Add("message", m.Message)
	}

	if m.Attachment != "" {
		data.Add("attachment", m.Attachment)
	}

	if m.PeerID != 0 {
		data.Add("peer_id", fmt.Sprint(m.PeerID))
	}

	data.Add("random_id", fmt.Sprint(m.RandomID))

	if m.UserID != 0 {
		data.Add("user_id", fmt.Sprint(m.UserID))
	}

	return data
}

func (k keyboard) URLEncoded() url.Values {
	data := url.Values{}
	jsonKeyboard, _ := json.Marshal(k)
	data.Add("keyboard", string(jsonKeyboard))
	return data
}

func (e eventAnswer) URLEncoded() url.Values {
	data := url.Values{}
	data.Add("event_id", e.EventID)
	data.Add("user_id", fmt.Sprint(e.UserID))
	data.Add("peer_id", fmt.Sprint(e.PeerID))

	jsonEventData, _ := json.Marshal(e.EventData)
	data.Add("event_data", string(jsonEventData))
	return data
}
