package middleware

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/test"
)

func TestLog(t *testing.T) {
	// Note: Just for the test coverage, not a real test.
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := echo.NewContext(req, rec, e)
	h := Log()(echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}))

	// Status 2xx
	h.Handle(c)

	// Status 3xx
	rec = test.NewResponseRecorder()
	c = echo.NewContext(req, rec, e)
	h = Log()(echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	}))
	h.Handle(c)

	// Status 4xx
	rec = test.NewResponseRecorder()
	c = echo.NewContext(req, rec, e)
	h = Log()(echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusNotFound, "test")
	}))
	h.Handle(c)

	// Status 5xx with empty path
	req = test.NewRequest(echo.GET, "", nil)
	rec = test.NewResponseRecorder()
	c = echo.NewContext(req, rec, e)
	h = Log()(echo.HandlerFunc(func(c echo.Context) error {
		return errors.New("error")
	}))
	h.Handle(c)
}

func TestLogIPAddress(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := echo.NewContext(req, rec, e)
	buf := new(bytes.Buffer)
	e.Logger().(*log.Logger).SetOutput(buf)
	ip := "127.0.0.1"
	h := Log()(echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}))

	// With X-Real-IP
	req.Header().Add(echo.XRealIP, ip)
	h.Handle(c)
	assert.Contains(t, buf.String(), ip)

	// With X-Forwarded-For
	buf.Reset()
	req.Header().Del(echo.XRealIP)
	req.Header().Add(echo.XForwardedFor, ip)
	h.Handle(c)
	assert.Contains(t, buf.String(), ip)

	// with req.RemoteAddr
	buf.Reset()
	h.Handle(c)
	assert.Contains(t, buf.String(), ip)
}
