package ws

import (
	"GemiApp/app/middleware"
	"GemiApp/services/auth"
	"GemiApp/types"
	"bytes"
	"context"
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

type httpFunc func(w http.ResponseWriter, r *http.Request)

func (h *WsHandler) WithAuth(f httpFunc) httpFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr, err := middleware.AuthMiddleware(r)
		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		var data *types.User

		if data, err = h.AuthSrv.UserStatus(usr.UserID); err != nil {
			cookie := middleware.NewEmptyCookie()
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			w.Write(nil)
			return
		}
		ctx := context.WithValue(r.Context(), "data", data)
		f(w, r.WithContext(ctx))
	}
}

// //////

func (h *WsHandler) HandleLobby(ws *websocket.Conn) {
	// usr := r.Context().Value("data").(*types.User)
	// lusr := LobbyUser{
	// 	ID:       usr.ID,
	// 	Username: usr.Username,
	// 	ImageUri: usr.ImgUri,
	// 	Ws:       ws,
	// }
	//
	// h.LobbyManager.Connect(lusr)

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
