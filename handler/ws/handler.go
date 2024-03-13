package ws

import (
	"GemiApp/services/auth"
	"GemiApp/services/game"
	"bytes"
	"html/template"
	"io"

	"github.com/Masterminds/sprig/v3"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type WsHandler struct {
	AuthSrv      *auth.AuthService
	GameSrv      *game.GameService
	LobbyManager *LobbyManager
}

func New(as *auth.AuthService, gs *game.GameService) *WsHandler {
	r := &WsHandler{
		AuthSrv:      as,
		GameSrv:      gs,
		LobbyManager: NewLobbyManager(),
	}

	go r.LobbyListener()
	return r
}

func (h *WsHandler) Route(r *mux.Router) {
	r.Handle("/ws/ludo/lobby", websocket.Handler(h.HandleLobby))
}

func (h *WsHandler) HandleLobby(ws *websocket.Conn) {
	user, err := h.useAuth(ws.Request())
	if err != nil {
		ws.Close()
		return
	}
	game, mode, err := h.validateGameAndMode(ws.Request())
	if err != nil {
		ws.Close()
		return
	}

	lusr := LobbyUser{
		ID:        user.ID,
		Username:  user.Username,
		ImageUri:  user.ImgUri,
		GameID:    game.ID,
		ModeID:    mode.ID,
		MaxPlayer: mode.NbPlayers,
		Ws:        ws,
	}

	h.LobbyManager.Connect(lusr)

	buf := make([]byte, 1024)
	for {
		buf, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				// CLIENT DISCONNECT
				h.LobbyManager.Disconnect(lusr)
				return
			}
			return
		}
		_ = buf
	}
}

func calEmptySpace(count int) []int {
	items := []int{}
	for i := 0; i < (count); i++ {
		items = append(items, i)
	}
	return items
}
func (h *WsHandler) LobbyListener() {
	for {
		room := <-h.LobbyManager.NotifQueue
		var buf bytes.Buffer
		tmpl := template.Must(
			template.New("lobby_ws.tmpl").Funcs(sprig.FuncMap()).ParseFiles("./web/views/ludo/lobby_ws.tmpl"),
		)
		tmpl.Execute(&buf, map[string]any{
			"room":       room,
			"emptySpace": calEmptySpace(room.MaxPlayers - len(room.Users)),
		})

		for _, v := range room.Users {
			v.Ws.Write(buf.Bytes())
		}
	}
}
