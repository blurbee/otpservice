package api

type StatusCode int

const (
	OK StatusCode = iota
	CONFIG_LOAD_FAILED
	INVALID_INPUT
	CONN_FAILED
	CONFIG_ERROR
	STORE_ERROR
	USER_NOT_FOUND
	VALUE_NOT_FOUND
	UNKNOWN_ERROR
	SEND_ERROR
)

var error_string = [...]string{
	"OK",
	"CONFIG_LOAD_FAILED",
	"INVALID_INPUT",
	"CONN_FAILED",
	"CONFIG_ERROR",
	"STORE_ERROR",
	"USER_NOT_FOUND",
	"VALUE_NOT_FOUND",
	"UNKNOWN_ERROR",
	"SEND_ERROR",
}

func (e StatusCode) Error() (s string) {
	return error_string[e]
}
