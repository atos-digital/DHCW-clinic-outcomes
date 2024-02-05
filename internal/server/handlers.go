package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/server/models"
	"github.com/atos-digital/DHCW-clinic-outcomes/internal/store/db"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui"
	"github.com/atos-digital/DHCW-clinic-outcomes/ui/pages"
)

const (
	formID   = "outcomes-form-id"
	formData = "outcomes-form-data"
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
		save, err := s.db.GetAllStates()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var subs []db.Submission
		// Ticket 51
		// Get all submissions

		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values[formID] = nil
		session.Values[formData] = []byte{}
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		ui.Index(pages.HomePage(subs, save)).Render(r.Context(), w)
	}
}

func (s *Server) handleNewForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.ClinicOutcomesFormPayload

		if r.Header.Get("HX-Trigger") != "new-form" {
			session, err := s.sess.Get(r, s.conf.CookieName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b := session.Values[formData]
			if b != nil {
				json.Unmarshal(b.([]byte), &data)
			}
		}

		w.Header().Set("Content-Type", "text/html")
		ui.Index(pages.OutcomesFormPage(models.State(data))).Render(r.Context(), w)
	}
}

func (s *Server) handleDraftForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ticket 50

		// Get id from URL

		// Get state from database

		// Get session store

		// Save id and state in the session

		// Set content type and render the form
	}
}

func (s *Server) handleAutosaveForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the body to get the latest form data.
		var data models.ClinicOutcomesFormPayload
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		if data.AddTest != nil {
			data.TestsRequired = append(data.TestsRequired, "")
			data.TestsUndertakenBy = append(data.TestsUndertakenBy, "")
			data.TestsBy = append(data.TestsBy, "Choose Option")
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
		session.Values[formData] = b
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		pages.OutcomesForm(models.State(data)).Render(r.Context(), w)
	}
}

func (s *Server) handleSaveForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.ClinicOutcomesFormPayload
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := session.Values[formID]
		if id != nil {
			err = s.db.UpdateState(id.(string), data)
		} else {
			err = s.db.StoreState(data)
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "Error storing form data", http.StatusInternalServerError)
			return
		}

		session.Values[formID] = nil
		session.Values[formData] = []byte{}
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Location", "/")
	}
}

func (s *Server) handleSubmitForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.ClinicOutcomesFormPayload
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		submission, err := models.Submit(models.State(data))

		if e, ok := err.(models.ErrorSubmit); ok && e.Error() == "Missing fields" {
			w.Header().Set("HX-Retarget", "#error-summary")
			w.Header().Set("HX-Reswap", "outerHTML show:window:top")
			pages.ErrorSummary(e.Errors).Render(r.Context(), w)
			return
		}

		err = s.db.StoreSubmission(submission)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error submitting form data", http.StatusInternalServerError)
			return
		}

		session, err := s.sess.Get(r, s.conf.CookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values[formData] = []byte{}
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Location", "/")
	}
}

func (s *Server) handleSubmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ticket 52
		// Get the submission from the database and render it.
	}
}
