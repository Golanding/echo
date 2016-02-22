// +build !appengine

package fasthttp

import "github.com/admpub/fasthttp"

type (
	RequestHeader struct {
		header *fasthttp.RequestHeader
	}

	ResponseHeader struct {
		header *fasthttp.ResponseHeader
	}
)

func (h *RequestHeader) Add(key, val string) {
	h.header.Set(key, val)
}

func (h *RequestHeader) Del(key string) {
	h.header.Del(key)
}

func (h *RequestHeader) Get(key string) string {
	return string(h.header.Peek(key))
}

func (h *RequestHeader) Set(key, val string) {
	h.header.Set(key, val)
}

func (h *RequestHeader) Object() interface{} {
	return h.header
}

func (h *ResponseHeader) Add(key, val string) {
	h.header.Set(key, val)
}

func (h *RequestHeader) reset(hdr *fasthttp.RequestHeader) {
	h.header = hdr
}

func (h *ResponseHeader) Del(key string) {
	h.header.Del(key)
}

func (h *ResponseHeader) Get(key string) string {
	return string(h.header.Peek(key))
}

func (h *ResponseHeader) Set(key, val string) {
	h.header.Set(key, val)
}

func (h *ResponseHeader) Object() interface{} {
	return h.header
}

func (h *ResponseHeader) reset(hdr *fasthttp.ResponseHeader) {
	h.header = hdr
}
