package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

// 全局响应处理
func init() {
	render.Respond = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if err, ok := v.(error); ok {
			// 默认错误响应
			if _, ok := r.Context().Value(render.StatusCtxKey).(int); !ok {
				render.Render(w, r, ErrInvalidRequest(err))
			}

			// 错误日志输出
			fmt.Printf("Logging err: %s\n", err.Error())

			// 错误处理，输出用户友好的错误
			render.Render(w, r, ErrInvalidRequest(errors.New("网络错误，请稍后再试")))
			return
		}
		render.DefaultResponder(w, r, v)
	}
}

const (
	ResponseOkCode    = 0
	ResponseOkMessage = "ok"
)

type JsonResponse struct {
	Code    int         `json:"code"`    // application-specific error code
	Message string      `json:"message"` // application-level error message, for debugging
	Result  interface{} `json:"result,omitempty"`
}

func NewJsonResponse(v interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    ResponseOkCode,
		Message: ResponseOkMessage,
		Result:  v,
	}
}

func (j *JsonResponse) Render(w http.ResponseWriter, r *http.Request) error {
	var httpStatusCode int
	if j.Code == ResponseOkCode {
		httpStatusCode = http.StatusOK
	} else {
		httpStatusCode = j.Code
	}
	render.Status(r, httpStatusCode)
	return nil
}

type ErrResponse struct {
	Code    int    `json:"code"`    // application-specific error code
	Message string `json:"message"` // application-level error message, for debugging

	Err            error  `json:"-"`      // low-level runtime error
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Code:           http.StatusBadRequest,
		Message:        err.Error(),
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Code:           http.StatusUnprocessableEntity,
		Message:        err.Error(),
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Code:           http.StatusInternalServerError,
		Message:        err.Error(),
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error.",
	}
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Code:           http.StatusNotFound,
		Message:        err.Error(),
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Resource not found.",
	}
}

func writeJsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
