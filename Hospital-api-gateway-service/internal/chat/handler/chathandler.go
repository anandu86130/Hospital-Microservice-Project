package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/chat/pbC"
	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	userpb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader variable specifies the parmeters of upgrading HTTP request
// var Upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// // HandleWebSocketConnection handles the weboscket connection and bidirectional streaming
// func HandleWebSocketConnection(c *gin.Context, client pb.ChatServiceClient, userClient userpb.UserServiceClient) {
// 	ctx := c.Request.Context()

// 	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println("Error upgrading to WebSocket:", err)
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()
// 	log.Println("WebSocket connection established")

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// Context canceled, stop processing messages
// 			log.Println("WebSocket connection closed")
// 			return
// 		default:
// 			_, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Println("Error reading message:", err)
// 				return
// 			}

// 			var message model.Message
// 			err = json.Unmarshal(msg, &message)
// 			if err != nil {
// 				log.Println("Error decoding JSON:", err)
// 				continue
// 			}

// 			// Checking the user and receiver IDs
// 			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.UserID)})
// 			if err != nil {
// 				log.Println("Error fetching user or invalid userID:", err)
// 				continue
// 			}

// 			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.ReceiverID)})
// 			if err != nil {
// 				log.Println("Error fetching receiver or invalid receiverID:", err)
// 				continue
// 			}

// 			stream, err := client.Connect(ctx)
// 			if err != nil {
// 				log.Println("Error calling chat service:", err)
// 				continue
// 			}
// 			ch := &clientHandle{
// 				stream:     stream,
// 				userID:     uint32(message.UserID),
// 				receiverID: uint32(message.ReceiverID),
// 			}

// 			err = conn.WriteMessage(websocket.TextMessage, msg)
// 			if err != nil {
// 				log.Println("Error writing message:", err)
// 				return
// 			}

// 			go ch.sentMessage(message.Message)
// 			go ch.receiveMessage(conn, uint32(message.UserID), uint32(message.ReceiverID))
// 		}
// 	}
// }

// // ChatPage loads the chat page.
// func ChatPage(c *gin.Context, client pb.ChatServiceClient) {
// 	timeout := time.Second * 1000
// 	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
// 	defer cancel()

// 	id := c.Query("id")
// 	userID, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println("Error converting id to int:", err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting id to int", "Error": err.Error()})
// 		return
// 	}

// 	receiverIDStr := c.Query("receiverId")
// 	receiverID, err := strconv.Atoi(receiverIDStr)
// 	if err != nil {
// 		log.Println("Error converting receiverId to int:", err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting receiverId to int", "Error": err.Error()})
// 		return
// 	}

// 	response, err := client.FetchHistory(ctx, &pb.ChatID{User_ID: uint32(userID), Receiver_ID: uint32(receiverID)})
// 	if err != nil {
// 		log.Println("Error calling chat client:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "Error calling chat client", "Error": err.Error()})
// 		return
// 	}

// 	c.HTML(http.StatusOK, "chat.html", gin.H{"response": response.Chats, "id": userID})
// }

// type clientHandle struct {
// 	userID     uint32
// 	receiverID uint32
// 	stream     pb.ChatService_ConnectClient
// }

// func (c *clientHandle) sentMessage(msg string) {
// 	message := &pb.Message{
// 		User_ID:     c.userID,
// 		Receiver_ID: c.receiverID,
// 		Content:     msg,
// 	}

// 	err := c.stream.Send(message)
// 	if err != nil {
// 		log.Printf("Error while sending message to server: %v", err)
// 	}
// }

// func (c *clientHandle) receiveMessage(conn *websocket.Conn, userID, receiverID uint32) {
// 	for {
// 		mssg, err := c.stream.Recv()
// 		if err != nil {
// 			log.Printf("Error receiving message from server: %v", err)
// 			return
// 		}

// 		if userID == mssg.Receiver_ID && receiverID == mssg.User_ID {
// 			dom := &model.Message{
// 				UserID:     uint(mssg.User_ID),
// 				ReceiverID: uint(mssg.Receiver_ID),
// 				Message:    mssg.Content,
// 			}
// 			msg, err := json.Marshal(dom)
// 			if err != nil {
// 				log.Println("Error encoding JSON:", err)
// 				return
// 			}

// 			err = conn.WriteMessage(websocket.TextMessage, msg)
// 			if err != nil {
// 				log.Println("Error writing message:", err)
// 				return
// 			}
// 		}
// 	}
// }

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = struct {
	sync.RWMutex
	m map[uint32]*websocket.Conn
}{m: make(map[uint32]*websocket.Conn)}

func HandleWebSocketConnection(c *gin.Context, client pb.ChatServiceClient, userClient userpb.UserServiceClient) {
	ctx := c.Request.Context()

	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Get userID from query params
	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting userID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid userID"})
		return
	}

	// Store the WebSocket connection in the map
	connections.Lock()
	connections.m[uint32(userID)] = conn
	connections.Unlock()

	// Ensure the connection is removed when the function exits
	defer func() {
		connections.Lock()
		delete(connections.m, uint32(userID))
		connections.Unlock()
	}()

	// Start a goroutine to handle incoming WebSocket messages
	go handleIncomingMessages(client, ctx)

	// Handle sending outgoing messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		var message model.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Check if the receiver exists
		_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.ReceiverID)})
		if err != nil {
			log.Println("Receiver not found:", err)
			continue
		}

		// Send the message to the receiver WebSocket
		err = sendToReceiver(uint32(message.ReceiverID), msg)
		if err != nil {
			log.Println("Error sending message to receiver:", err)
			continue
		}

		// Also send the message back to the sender
		err = sendToReceiver(uint32(message.UserID), msg)
		if err != nil {
			log.Println("Error sending message to sender:", err)
		}
	}
}

// handleIncomingMessages handles incoming messages from the streaming service.
func handleIncomingMessages(client pb.ChatServiceClient, ctx context.Context) {
	// Start the streaming to receive messages from the service
	stream, err := client.Connect(ctx)
	if err != nil {
		log.Println("Error calling chat service:", err)
		return
	}

	// Handle incoming messages from the stream
	for {
		mssg, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message from stream: %v", err)
			return
		}

		// Forward the message to the correct receiver via WebSocket
		receiverConn, err := getReceiverConnection(uint32(mssg.Receiver_ID))
		if err != nil {
			log.Printf("Error retrieving receiver WebSocket connection: %v", err)
			continue
		}

		// Prepare and send the message
		message := model.Message{
			UserID:     uint(mssg.User_ID),
			ReceiverID: uint(mssg.Receiver_ID),
			Message:    mssg.Content,
		}

		msg, err := json.Marshal(message)
		if err != nil {
			log.Println("Error marshalling message:", err)
			continue
		}

		// Send the message to the receiver's WebSocket
		if err := receiverConn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error sending message to receiver WebSocket:", err)
		}

		// Also send the message back to the sender
		senderConn, err := getReceiverConnection(uint32(mssg.User_ID))
		if err == nil {
			if err := senderConn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Error sending message to sender WebSocket:", err)
			}
		}
	}
}

func sendToReceiver(receiverID uint32, msg []byte) error {
	receiverConn, err := getReceiverConnection(receiverID)
	if err != nil {
		return fmt.Errorf("receiver connection not found: %v", err)
	}

	// Send the message to the receiver's WebSocket connection
	return receiverConn.WriteMessage(websocket.TextMessage, msg)
}

func getReceiverConnection(receiverID uint32) (*websocket.Conn, error) {
	connections.RLock()
	defer connections.RUnlock()

	conn, exists := connections.m[receiverID]
	if !exists {
		return nil, fmt.Errorf("no WebSocket connection found for receiver ID %d", receiverID)
	}
	return conn, nil
}

func ChatPage(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting userID to int:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting userID", "Error": err.Error()})
		return
	}

	receiverIDStr := c.Query("receiverId")
	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		log.Println("Error converting receiverID to int:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting receiverID", "Error": err.Error()})
		return
	}

	// Fetch chat history
	response, err := client.FetchHistory(ctx, &pb.ChatID{User_ID: uint32(userID), Receiver_ID: uint32(receiverID)})
	if err != nil {
		log.Println("Error calling chat client:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "Error fetching chat history", "Error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "chat.html", gin.H{"response": response.Chats, "id": userID})
}

func StartVideoCall(c *gin.Context, client pb.ChatServiceClient) {
	var req struct {
		UserID     uint32 `json:"user_id" binding:"required"`
		ReceiverID uint32 `json:"receiver_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	response, err := client.StartVideoCall(ctx, &pb.VideoCallRequest{
		User_ID:     req.UserID,
		Receiver_ID: req.ReceiverID,
	})
	if err != nil {
		log.Printf("Failed to start video call: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start video call", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"room_url": response.RoomUrl})
}
