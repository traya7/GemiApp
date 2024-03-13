package ws

import (
	"GemiApp/helpers"
	"GemiApp/types"
	"errors"
	"net/http"
)

func (h *WsHandler) useAuth(r *http.Request) (*types.User, error) {
	val, err := helpers.AuthMiddleware(r)
	if err != nil {
		return nil, err
	}
	var data *types.User
	if data, err = h.AuthSrv.UserStatus(val); err != nil {
		return nil, err
	}
	return data, nil
}

func (h *WsHandler) validateGameAndMode(r *http.Request) (*types.Game, *types.GameMode, error) {

	gid := r.URL.Query().Get("gid")
	mid := r.URL.Query().Get("id")
	game, err := h.GameSrv.GetGameDetails(gid)
	if err != nil || game.Status == false {
		return nil, nil, err
	}

	for _, v := range game.Modes {
		if v.ID == mid {
			return game, &v, nil
		}
	}
	return nil, nil, errors.New("not found mode")
}
