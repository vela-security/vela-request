package request

import "github.com/vela-security/vela-public/lua"

type state struct {
	version   uint16
	handshake bool
	IsCA      bool
	host      string
	after     int64
	subject   string
}

func (st *state) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	case "version":
		return lua.LInt(st.version)

	case "handshake":
		return lua.LBool(st.handshake)

	case "is_ca":
		return lua.LBool(st.IsCA)

	case "host":
		return lua.S2L(st.host)

	case "after":
		return lua.LNumber(st.after)

	case "subject":
		return lua.S2L(st.subject)

	default:
		return lua.LNil
	}
}
