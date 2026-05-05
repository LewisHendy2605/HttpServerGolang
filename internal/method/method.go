package method

/*
3.1. Method

The method token indicates the request method to be performed on the target resource.
The request method is case-sensitive.

	method         = token
- * RFC_9100: GET, HEAD, POST, PUT, DELETE, OPTIONS, TRACE
- * RFC_5789: PATCH
*/

const (
	Get     string = "GET"
	Head    string = "HEAD"
	Post    string = "POST"
	Put     string = "PUT"
	Patch   string = "PATCH"
	Delete  string = "DELETE"
	Options string = "OPTIONS"
	Trace   string = "TRACE"
)

// Validates a http method
func IsMethod(m string) bool {
	switch m {
	case Get,
		Head,
		Post,
		Put,
		Patch,
		Delete,
		Options,
		Trace:
		return true
	default:
		return false
	}
}
