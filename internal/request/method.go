package request

/*
3.1. Method

The method token indicates the request method to be performed on the target resource.
The request method is case-sensitive.

	method         = token
- * RFC_9100: GET, HEAD, POST, PUT, DELETE, OPTIONS, TRACE
- * RFC_5789: PATCH
*/

const (
	MethodGet     string = "GET"
	MethodHead    string = "HEAD"
	MethodPost    string = "POST"
	MethodPut     string = "PUT"
	MethodPatch   string = "PATCH"
	MethodDelete  string = "DELETE"
	MethodOptions string = "OPTIONS"
	MethodTrace   string = "TRACE"
)

// Validates a http method
func IsMethod(m string) bool {
	switch m {
	case MethodGet,
		MethodHead,
		MethodPost,
		MethodPut,
		MethodPatch,
		MethodDelete,
		MethodOptions,
		MethodTrace:
		return true
	default:
		return false
	}
}
