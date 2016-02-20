package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/gommon/log"
	"github.com/webx-top/echo/engine"
	"github.com/webx-top/echo/engine/standard"
)

type (
	ResponseRecorder struct {
		engine.Response
		Body *bytes.Buffer
	}
)

func NewRequest(method, url string, body io.Reader) engine.Request {
	r, _ := http.NewRequest(method, url, body)
	return standard.NewRequest(r)
}

func NewResponseRecorder() *ResponseRecorder {
	r := httptest.NewRecorder()
	return &ResponseRecorder{
		Response: standard.NewResponse(r, log.New("test")),
		Body:     r.Body,
	}
}
