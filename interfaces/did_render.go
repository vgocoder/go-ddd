package interfaces

import (
	"github.com/go-chi/render"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
	"net/http"
)

type UserDidResponse struct {
	*persistence.XJUserDid
}

func NewUserDidResponse(userDid *persistence.XJUserDid) *UserDidResponse {
	resp := &UserDidResponse{
		XJUserDid: userDid,
	}
	return resp
}

func (u *UserDidResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserDidListResponse(userDids []*persistence.XJUserDid) []render.Renderer {
	resp := []render.Renderer{}
	for _, userDid := range userDids {
		resp = append(resp, NewUserDidResponse(userDid))
	}
	return resp
}
