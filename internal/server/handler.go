package server

import (
	"io"

	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
	"github.com/LewisHendy2605/HttpServerGolang/internal/response"
)

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}
type Handler func(w io.Writer, req *request.Request) *HandlerError
