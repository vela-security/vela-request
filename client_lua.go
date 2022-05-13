package request

import (
	"fmt"
	"github.com/vela-security/vela-public/lua"
	"strings"
)

func (c *Client) String() string                         { return fmt.Sprintf("http.client %p", c) }
func (c *Client) Type() lua.LValueType                   { return lua.LTObject }
func (c *Client) AssertFloat64() (float64, bool)         { return 0, false }
func (c *Client) AssertString() (string, bool)           { return "", false }
func (c *Client) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (c *Client) Peek() lua.LValue                       { return c }

func (c *Client) H(L *lua.LState) int {
	data := L.CheckString(1)
	kv := strings.Split(data, ":")

	if len(kv) == 2 {
		c.SetHeader(kv[0], kv[1])
	}
	L.Push(c)
	return 1
}

func after(v lua.LValue) ResponseMiddleware {
	if v.Type() != lua.LTFunction {
		return nil
	}

	return func(_ *Client, response *Response) error {
		co := xEnv.Coroutine()
		defer xEnv.Free(co)
		cp := xEnv.P(v.(*lua.LFunction))
		return co.CallByParam(cp, response)
	}

}

func before(v lua.LValue) RequestMiddleware {
	if v.Type() != lua.LTFunction {
		return nil
	}

	return func(_ *Client, r *Request) error {
		co := xEnv.Coroutine()
		defer xEnv.Free(co)
		cp := xEnv.P(v.(*lua.LFunction))
		return co.CallByParam(cp, r)
	}
}

func (c *Client) NewIndex(L *lua.LState, key string, val lua.LValue) {

	switch key {
	case "after":
		c.OnAfterResponse(after(val))

	case "before":
		c.OnBeforeRequest(before(val))

	case "auth":
		c.SetAuthToken(val.String())

	case "proxy":
		c.SetProxy(val.String())
	}

}

func (c *Client) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "R":
		return c.R()

	case "GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE", "TRACE":
		r := c.R()
		r.Method = key

		return L.NewFunction(r.exec)

	case "H":
		return L.NewFunction(c.H)

	default:

	}

	return lua.LNil
}
