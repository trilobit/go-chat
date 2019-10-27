package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/trilobit/go-chat/src/models"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s *Websocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Connecting
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error("error creating upgrader: %v", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()

	// Authentication
	token := r.URL.Query().Get("token")
	user, err := s.userRepo.FindByToken(token)
	if err != nil || user == nil {
		s.logger.Errorf("unable to find user by token '%s': %v", token, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	s.hmu.Lock()

	s.hub[user.Email] = User{
		Model: user,
		Conn:  c,
	}

	s.hmu.Unlock()

	// Sending history
	s.logger.Infof("Sending history to user %s", user.Email)
	for _, message := range s.history {
		if err := c.WriteJSON(message); err != nil {
			s.logger.Errorf("error sending history: %v", err)
		}
	}
	s.system(fmt.Sprintf("> User <b>%s</b> enter the room.", user.Email))

	// Message handling
	s.logger.Info("start listening for incoming messages")
	for {
		var msg models.Message

		if err := c.ReadJSON(&msg); err != nil {
			s.logger.Errorf("error handling message: %v", err)
			return
		}

		s.logger.Infof("got incoming message: %v", msg)

		msg.Sender = user.Email
		msg.CreatedAt = time.Now()
		if msg.Receiver != "" {
			s.direct(msg)
		} else {
			s.broadcast(msg)
		}
	}
}

// broadcast sends message for every connected user
func (s *Websocket) broadcast(msg models.Message) {
	s.hmu.RLock()

	for _, receiver := range s.hub {
		if err := receiver.Conn.WriteJSON(msg); err != nil {
			s.logger.Errorf("error sending message: %v", err)
			continue
		}
	}

	s.history = append(s.history, msg)
	if err := s.historyRepo.Add(msg); err != nil {
		s.logger.Errorf("error writing history to db: %v", err)
	}

	s.hmu.RUnlock()
}

// broadcast sends message for every connected user
func (s *Websocket) system(text string) {
	msg := models.Message{
		ID:        0,
		Sender:    "system",
		Receiver:  "",
		Text:      text,
		CreatedAt: time.Now(),
	}

	s.hmu.RLock()

	for _, receiver := range s.hub {
		if err := receiver.Conn.WriteJSON(msg); err != nil {
			s.logger.Errorf("error sending message: %v", err)
			continue
		}
	}

	s.hmu.RUnlock()
}

// direct sends private message from user to user
func (s *Websocket) direct(msg models.Message) {
	s.hmu.RLock()
	userFrom, okFrom := s.hub[msg.Sender]
	userTo, okTo := s.hub[msg.Receiver]
	s.hmu.RUnlock()

	if !okFrom {
		s.logger.Errorf("unnown user from send message: %s", msg.Sender)
		return
	}

	if !okTo {
		s.logger.Errorf("unnown user to send message: %s", msg.Receiver)
		return
	}

	// Write to sender
	if err := userFrom.Conn.WriteJSON(msg); err != nil {
		s.logger.Errorf("error sending message: %v", err)
	}

	// Write to receiver
	if err := userTo.Conn.WriteJSON(msg); err != nil {
		s.logger.Errorf("error sending message: %v", err)
	}
}
