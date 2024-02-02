package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/server/models"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui/pages"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui/tables"
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

func (s *Server) handlePageIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subs, err := s.db.GetAllSubmissions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		ui.Index(pages.Home(subs)).Render(r.Context(), w)
	}
}

func (s *Server) handleViewSubmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		sub, err := s.db.GetSubmission(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		tables.SubmittedFormAnswers(sub).Render(r.Context(), w)
	}
}

func (s *Server) handlePageOutcomes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b := session.Values["outcomes-form-data"]
		w.Header().Set("Content-Type", "text/html")
		var data models.OutcomesForm
		if b != nil {
			json.Unmarshal(b.([]byte), &data)
		}
		ui.Index(pages.Outcomes(data.State())).Render(r.Context(), w)
	}
}

func (s *Server) handleOutcomesForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the body to get the latest form data.
		var data models.OutcomesForm
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		if data.AddTest != nil {
			data.FollowUpTestsRequired = append(data.FollowUpTestsRequired, "")
			data.FollowUpTestsUndertaken = append(data.FollowUpTestsUndertaken, "")
			data.FollowUpTestsBy = append(data.FollowUpTestsBy, "Day Prior to the Clinic")
		}
		// Back into bytes
		b, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Get the session store
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Update and save
		session.Values["outcomes-form-data"] = b
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		pages.Outcomes(data.State()).Render(r.Context(), w)
	}
}

func (s *Server) handleSaveOutcomes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the body to get the latest form data.
		var data models.OutcomesForm
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Store the data in the database.
		err = s.db.StoreState(data.State())
		if err != nil {
			log.Println(err)
			http.Error(w, "Error storing form data", http.StatusInternalServerError)
			return
		}

		// Get the session store
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Update and save
		session.Values["outcomes-form-data"] = []byte{}
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Location", "/")
	}
}

func (s *Server) handleSubmitOutcomes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the body to get the latest form data.
		var data models.OutcomesForm
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		submission, err := data.State().Submit()
		log.Println(submission)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error submitting form data", http.StatusInternalServerError)
			return
		}

		// Store the data in the database.
		err = s.db.StoreSubmission(submission)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error submitting form data", http.StatusInternalServerError)
			return
		}

		// Get the session store
		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Update and save
		session.Values["outcomes-form-data"] = []byte{}
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Location", "/")
	}
}
