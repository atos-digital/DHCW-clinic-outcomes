package server

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed assets
var assets embed.FS

func (s *Server) Routes() error {
	// special case: handler assets content for Chi router with subroute, default go router in 1.22 will not require this step
	contentAssets, err := fs.Sub(fs.FS(assets), "assets")
	if err != nil {
		return err
	}
	s.r.Method(http.MethodGet, "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(contentAssets))))
	s.r.Method(http.MethodGet, "/favicon.ico", s.HandleFavicon())

	s.r.Method(http.MethodGet, "/", s.handlePageIndex())
	s.r.Method(http.MethodGet, "/form", s.handleNewForm())
	s.r.Method(http.MethodGet, "/form/{id}", s.handleDraftForm())
	s.r.Method(http.MethodGet, "/submissions/{id}", s.handleSubmission())

	s.r.Route("/hx", func(r chi.Router) {
		r.Method(http.MethodPost, "/autosave-form", s.handleAutosaveForm())
		r.Method(http.MethodPost, "/save-form", s.handleSaveForm())
		r.Method(http.MethodPost, "/submit-form", s.handleSubmitForm())
	})

	return nil
}
