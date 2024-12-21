package chat

import (
	"github.com/anandu86130/Hospital-api-gateway/internal/chat/handler"
	"github.com/gin-gonic/gin"
)

// Chat handles the websocket connection in the chat page.
func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

// VideoCall handles initiating a video call.
func (c *Chat) VideoCall(ctx *gin.Context) {
	handler.StartVideoCall(ctx, c.client)
}

// ChatPage handles to load the chat page
func (c *Chat) ChatPage(ctx *gin.Context) {
	handler.ChatPage(ctx, c.client)
}