package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	_ "modernc.org/sqlite"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/config"
	"github.com/atos-digital/DHCW-clinic-outcomes/internal/store/db"
)

type Server struct {
	r    *chi.Mux
	srv  *http.Server
	conf config.Config
	sess sessions.Store
	db   *db.DB
}

func New(conf config.Config) (*Server, error) {
	s := new(Server)
	s.conf = conf
	s.r = chi.NewRouter()
	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Handler: s.r,
	}
	s.sess = s.cookieStore()
	db, err := db.NewSqlite(conf.SqlitePath)
	if err != nil {
		return nil, err
	}
	s.db = db
	err = s.db.Migrate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.middleware()
	s.Routes()
	// address for use when testing cookies locally
	if s.conf.Host == "0.0.0.0" {
		log.Printf("server: listening on http://localhost:%s", s.conf.Port)
	} else {
		log.Printf("server: listening on http://%s", s.srv.Addr)
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
