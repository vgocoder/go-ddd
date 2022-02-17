package interfaces

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/application"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/hlog"
	"kingsmith.com.cn/ks-horizon/interfaces/restapi"
)

func (impl *RestApiImpl) GetDidsCount(w http.ResponseWriter, r *http.Request) {
	count, err := application.NewFacades().DidApplication.Count()
	if err != nil {
		ErrInvalidRequest(err)
		return
	}
	writeJsonResponse(w, count)
}

func (impl *RestApiImpl) GetDids(w http.ResponseWriter, r *http.Request, p restapi.GetDidsParams) {
	limit, _ := strconv.Atoi(p.Limit)
	page, _ := strconv.Atoi(p.Page)
	items, err := application.NewFacades().DidApplication.GetAll(qm.Limit(limit), qm.Offset((page-1)*limit))
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("Error occurred.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, items)
}
