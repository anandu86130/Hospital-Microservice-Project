package model

type Message struct {
	UserID     uint   `json:"userId"`
	ReceiverID uint   `json:"receiverId"`
	Message    string `json:"content"`
	CallType   string `json:"callType,omitempty"`
	RoomID     string `json:"roomId,omitempty"`
	Duration   int    `json:"duration,omitempty"`
}

type ChatRequest struct {
	FriendID string `query:"FriendID" validate:"required"`
	Offset   string `query:"Offset" validate:"required"`
	Limit    string `query:"Limit" validate:"required"`
}
type TempMessage struct {
	SenderID    string
	RecipientID string `json:"RecipientID" validate:"required"`
	Content     string `json:"Content" validate:"required"`
	Timestamp   string `json:"TimeStamp" validate:"required"`
}
