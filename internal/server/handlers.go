package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/atos-digital/DHCW-clinic-outcomes/ui"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui/forms"
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

func (s *Server) handlePageOutcomes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b := session.Values["outcomes-form-data"]
		fmt.Println(session.Values["outcomes-form-data"])
		w.Header().Set("Content-Type", "text/html")
		var data map[string]string
		if b != nil {
			json.Unmarshal([]byte(b.(string)), &data)
		}
		ui.Index(pages.Outcomes(data)).Render(r.Context(), w)
	}
}

func (s *Server) handleOutcomesForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(session.Values["outcomes-form-data"])
		session.Values["outcomes-form-data"] = string(b)
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var outcomesForm map[string]string
		json.Unmarshal(b, &outcomesForm)
		pages.Outcomes(outcomesForm).Render(r.Context(), w)
	}
}

func (s *Server) handleAddFollowupTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		forms.FollowupTest().Render(r.Context(), w)
	}
}
