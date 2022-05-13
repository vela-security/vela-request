package request

import (
	"github.com/vela-security/vela-public/lua"
	"strings"
)

func (r *Response) Type() lua.LValueType                   { return lua.LTObject }
func (r *Response) AssertFloat64() (float64, bool)         { return 0, false }
func (r *Response) AssertString() (string, bool)           { return "", false }
func (r *Response) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (r *Response) Peek() lua.LValue                       { return r }

func (r *Response) catch(L *lua.LState) int {

	if r.Err != nil {
		xEnv.Errorf("%s request error %v", r.Request.URL, r.Err)
		return 0
	}

	if r.RawResponse == nil {
		xEnv.Errorf("%s request not found response", r.Request.URL)
		return 0
	}

	n := L.GetTop()
	code := r.StatusCode()
	for i := 1; i <= n; i++ {
		if L.CheckInt(i) == code {
			return 0
		}
	}

	xEnv.Errorf("%s request not found valid status code , got: %d body: %s", r.Request.URL, code, r.Body())
	return 0
}

func (r *Response) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "body":
		return lua.B2L(r.Body())
	case "size":
		return lua.LInt(r.size)
	case "code":
		return lua.LInt(r.StatusCode())
	case "url":
		return lua.S2L(r.Request.URL)

	case "catch":
		return L.NewFunction(r.catch)
	}

	if strings.HasPrefix(key, "http_") {
		return lua.S2L(r.Header().Get(U2H(key[5:])))
	}

	return lua.LNil
}
