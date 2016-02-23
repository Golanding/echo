// +build !appengine

package fasthttp

import "github.com/admpub/fasthttp"

type (
	URL struct {
		url *fasthttp.URI
	}
)

func (u *URL) SetPath(path string) {
	u.url.SetPath(path)
}

func (u *URL) Path() string {
	return string(u.url.Path())
}

func (u *URL) QueryValue(name string) string {
	return string(u.url.QueryArgs().Peek(name))
}

func (u *URL) RawQuery() string {
	return string(u.url.QueryString())
}

func (u *URL) Object() interface{} {
	return u.url
}

func (u *URL) reset(url *fasthttp.URI) {
	u.url = url
}
