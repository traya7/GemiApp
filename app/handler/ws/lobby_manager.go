package ws

import "github.com/gorilla/websocket"

type LobbyUser struct {
	ID       string
	Username string
	ImageUri string
	Ws       *websocket.Conn
}

type MainQueueMessage struct {
	Op   string
	User LobbyUser
}

type NotifQueueMessage struct {
	Users map[string]LobbyUser
}

type LobbyManager struct {
	Clients    map[string]LobbyUser
	MainQueue  chan MainQueueMessage
	NotifQueue chan NotifQueueMessage
}

func NewLobbyManager() *LobbyManager {
	r := &LobbyManager{
		Clients:    map[string]LobbyUser{},
		MainQueue:  make(chan MainQueueMessage),
		NotifQueue: make(chan NotifQueueMessage),
	}
	go r.QueueListener()
	return r
}
func (m *LobbyManager) QueueListener() {
	for {
		msg := <-m.MainQueue
		if msg.Op == "CONNECT" {
			m.Clients[msg.User.ID] = msg.User
			m.NotifQueue <- NotifQueueMessage{Users: m.Clients}
			continue
		}

		if msg.Op == "DISCONNECT" {
			delete(m.Clients, msg.User.ID)
			m.NotifQueue <- NotifQueueMessage{Users: m.Clients}
			continue
		}

		return
	}
}

func (m *LobbyManager) Connect(user LobbyUser) {
	m.MainQueue <- MainQueueMessage{Op: "CONNECT", User: user}
}

func (m *LobbyManager) Disconnect(user LobbyUser) {
	m.MainQueue <- MainQueueMessage{Op: "DISCONNECT", User: user}
}
