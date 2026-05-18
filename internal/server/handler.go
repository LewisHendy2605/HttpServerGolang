package server

import (
	"github.com/LewisHendy2605/HttpServerGolang/internal/request"
	"github.com/LewisHendy2605/HttpServerGolang/internal/response"
)

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}
type Handler func(res *response.Response, req *request.Request) *HandlerError
