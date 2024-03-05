package ws

import (
	"GemiApp/app/middleware"
	"GemiApp/services/auth"
	"GemiApp/types"
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type WsHandler struct {
	AuthSrv      *auth.AuthService
	LobbyManager *LobbyManager
}

func New(as *auth.AuthService) *WsHandler {
	r := &WsHandler{
		AuthSrv:      as,
		LobbyManager: NewLobbyManager(),
	}

	go r.LobbyListener()
	return r
}

func (h *WsHandler) Route(r *mux.Router) {
	r.Handle("/ws/ludo/lobby", websocket.Handler(h.HandleLobby))
}

func (h *WsHandler) useAuth(r *http.Request) (*types.User, error) {
	usr, err := middleware.AuthMiddleware(r)
	if err != nil {
		return nil, err
	}
	var data *types.User
	if data, err = h.AuthSrv.UserStatus(usr.UserID); err != nil {
		return nil, err
	}
	return data, nil
}

func (h *WsHandler) HandleLobby(ws *websocket.Conn) {
	user, err := h.useAuth(ws.Request())
	if err != nil {
		ws.Close()
	}
	lusr := LobbyUser{
		ID:       user.ID,
		Username: user.Username,
		ImageUri: user.ImgUri,
		Ws:       ws,
	}

	h.LobbyManager.Connect(lusr)

	buf := make([]byte, 1024)
	for {
		buf, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				// CLIENT DISCONNECT
				// h.LobbyManager.Disconnect(lusr)
				return
			}
			log.Println(err)
			return
		}
		_ = buf
	}
}

func (h *WsHandler) LobbyListener() {
	for {
		msg := <-h.LobbyManager.NotifQueue
		log.Println("msg sent!")
		var buf bytes.Buffer
		tmpl := template.Must(template.New("lobby_ws.tmpl").Funcs(sprig.FuncMap()).ParseFiles("./app/web/views/ludo/lobby_ws.tmpl"))
		tmpl.Execute(&buf, msg.Users)

		for _, v := range msg.Users {
			v.Ws.Write(buf.Bytes())
		}
	}
}
