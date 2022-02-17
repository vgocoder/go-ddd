package interfaces

import (
	"github.com/friendsofgo/errors"
	"github.com/go-chi/render"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
	"kingsmith.com.cn/ks-horizon/interfaces/restapi"
	"net/http"
)

type GetUsersRequest struct {
	*restapi.GetUsersJSONRequestBody
}

func (g *GetUsersRequest) Bind(r *http.Request) error {
	if g == nil {
		return errors.New("bad request params")
	}
	if g.Limit == 0 {
		return errors.New("bad request params:limit")
	}
	if g.Page == 0 {
		return errors.New("bad request params:page")
	}
	return nil
}

type UserResponse struct {
	*persistence.XJUser
}

func (u *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserResponse(user *persistence.XJUser) *UserResponse {
	resp := &UserResponse{
		XJUser: user,
	}
	return resp
}

func NewUserListResponse(users []*persistence.XJUser) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, NewUserResponse(user))
	}
	return list
}
