package apiserver

import (
	"io"
	"net/http"
)

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/getUser", s.handleGetUser())
	// http.Handle("/", s.router)
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *APIServer) handleGetUser(guid string) http.HandlerFunc {

	user, err := s.store.User.Get(guid)

	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, user.RefreshToken)

	}
}
