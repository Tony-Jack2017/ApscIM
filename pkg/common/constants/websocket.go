package constants

// websocket connection status

const (
	ConnOpenedSuc = 10001
	ConnClosedSuc = 10002
	ConnClosedFal = 10003
	ConnClosing   = 10004
)

const AuthError = 20001  // token invalided
const ParamsError = 2002 // params invalided
