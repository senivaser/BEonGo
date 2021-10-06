package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-errors/errors"

	"github.com/senivaser/BEonGo/internal/app/utils"
)

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/auth/guid/login", s.handleGUIDLogin())
	s.router.HandleFunc("/auth/refresh", s.handleRefresh())
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *APIServer) handleGUIDLogin() http.HandlerFunc {

	type request struct {
		Guid string
	}

	type response struct {
		AccessToken  string
		RefreshToken string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New(err))
			return
		}

		_, err := s.store.User.GetBy("guid", req.Guid)

		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		accessToken, errT := s.CreateAccessToken(req.Guid)
		refreshToken, errT := s.CreateRefreshToken(req.Guid)

		if errT != nil {
			s.error(w, r, 500, errors.New(errT))
			return
		}

		res := &response{
			AccessToken:  utils.ToString(accessToken),
			RefreshToken: utils.ToString(refreshToken),
		}

		s.respond(w, r, 200, res)

	}
}

func (s *APIServer) handleRefresh() http.HandlerFunc {

	type request struct {
		RefreshToken string
	}

	type response struct {
		AccessToken  string
		RefreshToken string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, errors.Wrap(err, 0))
			return
		}

		guid, err := s.ValidateRefreshToken(req.RefreshToken)
		fmt.Println("guid:" + utils.ToString(guid))

		if err != nil {
			s.error(w, r, 403, errors.New("Token is not valid"))
			return
		}

		user, errG := s.store.User.GetBy("guid", utils.ToString(guid))

		if errG != nil {
			s.error(w, r, 403, errG)
			return
		}

		checkResult := s.CheckPasswordHash(req.RefreshToken, user.RefreshToken)

		if checkResult == false {
			s.error(w, r, 403, errors.New("Token is not valid"))
			return
		}

		accessToken, err := s.CreateAccessToken(utils.ToString(guid))
		refreshToken, err := s.CreateRefreshToken(utils.ToString(guid))

		res := &response{
			AccessToken:  utils.ToString(accessToken),
			RefreshToken: utils.ToString(refreshToken),
		}

		s.respond(w, r, 200, res)

	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err *errors.Error) {
	s.respond(w, r, code, map[string]string{"error": err.Error(), "stack": err.ErrorStack()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
