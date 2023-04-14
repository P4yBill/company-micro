package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

type GeneralErrResponse struct {
	HTTPStatusCode int `json:"-"` // http response status code

	Success bool   `json:"success"`
	Message string `json:"message"` // user-level status message
}

func NewGeneralErrResponse(status int, message string) render.Renderer {
	return &GeneralErrResponse{
		HTTPStatusCode: status,
		Message:        message,
	}
}

func NewErrResponse(status int, err error, message string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     message,
		ErrorText:      err.Error(),
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (e *GeneralErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
