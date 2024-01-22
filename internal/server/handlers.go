package server

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"

	"github.com/atos-digital/DHCW-clinic-outcomes/ui"
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
		w.Header().Set("Content-Type", "text/html")
		if session.Values["outcomes-form-data"] != nil {
			ui.Index(pages.Outcomes(session.Values["outcomes-form-data"].(map[string]string))).Render(r.Context(), w)
		} else {
			ui.Index(pages.Outcomes(map[string]string{})).Render(r.Context(), w)
		}
	}
}

func (s *Server) handleOutcomesForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var outcomesForm map[string]string
		err := json.NewDecoder(r.Body).Decode(&outcomesForm)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["outcomes-form-data"] = outcomesForm
		session.Save(r, w)
		b, _ := json.MarshalIndent(outcomesForm, "", "  ")
		os.Stdout.Write(b)
		pages.Outcomes(outcomesForm).Render(r.Context(), w)
	}
}

func (s *Server) handleOutcomesOptionsRadio() http.HandlerFunc {
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
		session.Values["outcomes-option"] = []string{selected}
		session.Save(r, w)

		w.Header().Set("Content-Type", "text/html")
		pages.OutcomesOptions(nil).Render(r.Context(), w)
	}
}

// func (s *Server) handleRadio() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		radioGroupName := r.Header.Get("HX-Trigger")
// 		err := r.ParseForm()
// 		if err != nil {
// 			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
// 			return
// 		}
// 		selected := strings.Join(r.Form[radioGroupName], " ")

// 		session, err := s.sess.Get(r, s.conf.CookieName)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		session.Values[radioGroupName] = []string{selected}
// 		session.Save(r, w)

// 		w.Header().Set("Content-Type", "text/html")
// 		components.RadioGroupWithTextbox().Render(r.Context(), w)
// 	}
// }

func (s *Server) handleAddFollowupTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		pages.FollowupTest().Render(r.Context(), w)
	}
}
