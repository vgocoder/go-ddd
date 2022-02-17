package interfaces

import (
	"fmt"
	"github.com/go-chi/render"
	"kingsmith.com.cn/ks-horizon/application"
	"net/http"
	"strconv"
)

func (impl *RestApiImpl) PutRegionXjId(w http.ResponseWriter, r *http.Request, paramsXjId string) {
	// 接收参数
	xjIdInt, err := strconv.Atoi(paramsXjId)
	fmt.Println(xjIdInt)
	xjId := uint(xjIdInt)
	// 调用
	err = application.NewFacades().RegionApplication.UpsertRegion(xjId)
	if err != nil {
		ErrRender(err)
	}
	render.JSON(w, r, NewJsonResponse(true))
}
