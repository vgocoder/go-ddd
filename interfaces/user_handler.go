package interfaces

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/application"
	"net/http"
)

func (impl *RestApiImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	// 参数验证
	data := &GetUsersRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	params := data.GetUsersJSONRequestBody
	limit := params.Limit
	page := params.Page
	// 业务处理
	users, err := application.NewFacades().UserApplication.GetAll(qm.Limit(limit), qm.Offset((page-1)*limit))
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	// 结果处理
	resp := NewJsonResponse(NewUserListResponse(users))
	if err := render.Render(w, r, resp); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (impl *RestApiImpl) GetUsersCount(w http.ResponseWriter, r *http.Request) {
	count, err := application.NewFacades().UserApplication.Count()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	fmt.Println(count)
	render.JSON(w, r, NewJsonResponse(count))
}
