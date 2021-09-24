package apiserver

import (
	"encoding/json"
	"fmt"
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

func (s *APIServer) handleGetUser() http.HandlerFunc {

	type request struct {
		Guid string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		fmt.Println("RRRXXX: ", req)

		user, err := s.store.User.Get(req.Guid)

		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		s.respond(w, r, 200, user)

	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
