package request

import (
	"fmt"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

type httpLuaApi struct{}

func (hapi *httpLuaApi) String() string                         { return fmt.Sprintf("request.client %p", hapi) }
func (hapi *httpLuaApi) Type() lua.LValueType                   { return lua.LTObject }
func (hapi *httpLuaApi) AssertFloat64() (float64, bool)         { return 0, false }
func (hapi *httpLuaApi) AssertString() (string, bool)           { return "", false }
func (hapi *httpLuaApi) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (hapi *httpLuaApi) Peek() lua.LValue                       { return hapi }

func (hapi *httpLuaApi) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	case "client":
		return New()

	case "GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE", "TRACE":
		r := New().R()
		r.Method = key
		return L.NewFunction(r.exec)

	case "TLS":
		return L.NewFunction(newLuaTlsInfo)

	case "save":
		r := New().R()
		return L.NewFunction(r.save)

	}

	return lua.LNil
}

func WithEnv(env assert.Environment) {
	xEnv = env
	cli := &httpLuaApi{}
	env.Global("http", cli)
}
