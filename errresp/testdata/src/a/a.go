package a

import "errors"

type Server struct{}

func (s Server) ErrorResponse(err error) {}

func doSomething() {}

func bad(s Server) {
	s.ErrorResponse(errors.New("fail")) // want `ErrorResponse\(\) not immediately followed by return`
	doSomething()
}

func good(s Server) {
	s.ErrorResponse(errors.New("fail"))
	return
}
