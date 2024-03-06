package ws

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type LobbyUser struct {
	ID        string
	GameID    string
	ModeID    string
	MaxPlayer int
	Username  string
	ImageUri  string
	Ws        *websocket.Conn
}
type LobbyRoom struct {
	ID         string
	GameID     string
	ModeID     string
	MaxPlayers int
	Users      map[string]LobbyUser
}

type MainQueueMessage struct {
	Op   string
	User LobbyUser
}

type LobbyManager struct {
	Rooms      map[string]LobbyRoom
	MainQueue  chan MainQueueMessage
	NotifQueue chan LobbyRoom
}

func NewLobbyManager() *LobbyManager {
	r := &LobbyManager{
		Rooms:      map[string]LobbyRoom{},
		MainQueue:  make(chan MainQueueMessage),
		NotifQueue: make(chan LobbyRoom),
	}
	go r.QueueListener()
	return r
}
func (m *LobbyManager) QueueListener() {
	for {
		msg := <-m.MainQueue
		GameID := msg.User.GameID
		MaxPlayers := msg.User.MaxPlayer
		ModeID := msg.User.ModeID

		if msg.Op == "CONNECT" {
			room_id := m.FindRoom(GameID, ModeID)
			if room_id == "" {
				room_id = m.CreateRoom(GameID, ModeID, MaxPlayers)
			}
			room := m.AddUserToRoom(room_id, msg.User)
			m.NotifQueue <- room
			continue
		}

		if msg.Op == "DISCONNECT" {
			room, ok := m.DropUserFromRoom(GameID, ModeID, msg.User.ID)
			if ok {
				m.NotifQueue <- room
				continue
			}
			log.Println("ROOM NOT FOUND")
		}
	}
}

func (m *LobbyManager) Connect(user LobbyUser) {
	log.Println("CONNECT", user.GameID, user.MaxPlayer)
	m.MainQueue <- MainQueueMessage{Op: "CONNECT", User: user}
}

func (m *LobbyManager) Disconnect(user LobbyUser) {
	m.MainQueue <- MainQueueMessage{Op: "DISCONNECT", User: user}
}

func (m *LobbyManager) FindRoom(game_id, mode_id string) string {
	for id, v := range m.Rooms {
		if v.GameID == game_id && v.ModeID == mode_id {
			return id
		}
	}
	return ""
}

func (m *LobbyManager) CreateRoom(game_id, mode_id string, nb int) string {
	id := uuid.New().String()
	m.Rooms[id] = LobbyRoom{
		ID:         id,
		GameID:     game_id,
		ModeID:     mode_id,
		MaxPlayers: nb,
		Users:      map[string]LobbyUser{},
	}
	return id
}

func (m *LobbyManager) AddUserToRoom(room_id string, usr LobbyUser) LobbyRoom {
	room := m.Rooms[room_id]
	m.Rooms[room_id].Users[usr.ID] = usr
	m.Rooms[room_id] = room
	return room
}

func (m *LobbyManager) DropUserFromRoom(game_id, mode_id, user_id string) (LobbyRoom, bool) {
	for id, v := range m.Rooms {
		if v.GameID == game_id && v.ModeID == mode_id {
			room := m.Rooms[id]
			delete(m.Rooms[id].Users, user_id)
			m.Rooms[id] = room
			return room, true
		}
	}
	return LobbyRoom{}, false
}
