package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
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
	defer c.Close()

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

	// Message handling
	s.logger.Info("start listening for incoming messages")
	for {
		var msg Message

		if err := c.ReadJSON(&msg); err != nil {
			s.logger.Errorf("error handling message: %v", err)
			return
		}

		s.logger.Infof("got incoming message: %v", msg)

		msg.From = user.Email
		msg.DateTime = time.Now()
		if msg.To != "" {
			// Send message as private
			s.hmu.RLock()
			userTo, ok := s.hub[msg.To]
			s.hmu.RUnlock()

			if !ok {
				s.logger.Errorf("unnown user to send message: %s", msg.To)
				continue
			}

			// Write to sender
			if err := c.WriteJSON(msg); err != nil {
				s.logger.Errorf("error sending message: %v", err)
			}

			// Write to receiver
			if err := userTo.Conn.WriteJSON(msg); err != nil {
				s.logger.Errorf("error sending message: %v", err)
			}
		} else {
			// Send message as public
			s.hmu.RLock()

			for _, user := range s.hub {
				if err := user.Conn.WriteJSON(msg); err != nil {
					s.logger.Error("error sending message: %v", err)
					continue
				}
			}

			s.history = append(s.history, msg)
			s.hmu.RUnlock()
		}
	}
}
