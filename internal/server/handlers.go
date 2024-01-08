package server

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/atos-digital/10.10.0-template/ui/pages"
)

func (s *Server) HandleFavicon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := assets.ReadFile("assets/img/favicon.ico")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(b)
	})
}

func (s *Server) handlePageIndex() http.Handler {
	return templ.Handler(pages.DefaultHome, templ.WithContentType("text/html"))
}

func (s *Server) handlePageOutcomes() http.Handler {
	return templ.Handler(pages.DefaultOutcomes, templ.WithContentType("text/html"))
}
