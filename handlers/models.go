package handlers

type callback struct {
	GroupID int64  `json:"group_id"`
	Type    string `json:"type"`
	Object  object `json:"object"`
	EventID string `json:"event_id"`
}

type object struct {
	Message message `json:"message"`
}

type message struct {
	PeerId int64  `json:"peer_id"`
	FromId int64  `json:"from_id"`
	Text   string `json:"text"`
}

type sendMessage struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	RandomID    int32  `json:"random_id"`
	PeerId      int64  `json:"peer_id"`
	Message     string `json:"message"`
	V           string `json:"v"`
}

type callbackConfirmationRequest struct {
	AccessToken string `json:"access_token"`
	GroupID     int64  `json:"group_id"`
	V           string `json:"v"`
}

type callbackConfirmationResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	Code string `json:"code"`
}
