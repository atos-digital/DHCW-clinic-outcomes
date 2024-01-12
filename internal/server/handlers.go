package server

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"

	"github.com/atos-digital/DHCW-clinic-outcomes/ui/pages"
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

func (s *Server) handleRadio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}
		selected := strings.Join(r.Form["outcomes-option"], " ")

		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["outcomes-option"] = selected
		session.Save(r, w)

		w.Header().Set("Content-Type", "text/html")
		pages.OutcomesOptions().Render(r.Context(), w)
	}
}
