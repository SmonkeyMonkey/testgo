package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"test/internal/app/models"
	"test/internal/app/store"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store store.Store) *server {
	srv := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	srv.initRoutes()
	return srv
}
func (s *server) initRoutes() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/users", s.handleGetAllUsers()).Methods("GET")
	s.router.HandleFunc("/games", s.handleGetAllGames()).Methods("GET")
	s.router.HandleFunc("/games", s.handleGamesCreate()).Methods("POST")
}
func (s *server) handleGetAllUsers() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users := s.store.User().GetAll()
		s.respond(writer, request, http.StatusOK, users)
	}
}
func (s *server) handleGetAllGames() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users := s.store.Game().GetAll()
		s.respond(writer, request, http.StatusOK, users)
	}
}
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email     string `json:"email"`
		LastName  string `json:"last_name"`
		Country   string `json:"country"`
		City      string `json:"city"`
		Gender    string `json:"gender"`
		BirthDate string `json:"birth_date"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:     req.Email,
			LastName:  req.LastName,
			Country:   req.Country,
			City:      req.City,
			Gender:    req.Gender,
			BirthDate: req.BirthDate,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}
func (s *server) handleGamesCreate() http.HandlerFunc {
	type request struct {
		PointsGained string `json:"points_gained"`
		WinStatus    string `json:"win_status"`
		GameType     string `json:"game_type"`
		Created      string `json:"created"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		g := &models.Game{
			PointsGained: req.PointsGained,
			WinStatus:    req.PointsGained,
			GameType:     req.GameType,
			Created:      req.Created,
		}
		s.store.Game().Create(g)
		s.respond(w, r, http.StatusCreated, g)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
