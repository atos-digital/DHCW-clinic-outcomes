package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ClinicOutcomesFormSubmit struct {
	Details          DetailsSubmit
	CancerPathway    string
	Outcome          OutcomeSubmit
	FollowUpRequired bool
	FollowUp         FollowUpSubmit
	OtherInformation string
}

type DetailsSubmit struct {
	DateTime  time.Time
	Type      string
	Specialty string
	Clinician string
}

type OutcomeSubmit struct {
	Answer        string
	AnswerDetails string
}

type FollowUpSubmit struct {
	Pathway                     string
	SameClinician               string
	SameClinic                  string
	SeeIn                       string
	DateTime                    time.Time
	Hospital                    string
	AppointmentPriority         string
	ClinicalCondition           string
	PreferredConsultationMethod string
	Tests                       []Test
}

type ErrorSubmit struct {
	Errors []string
}

func (e ErrorSubmit) Error() string {
	if len(e.Errors) > 0 {
		return "Missing fields"
	}
	return ""
}

func Submit(state ClinicOutcomesFormState) (ClinicOutcomesFormSubmit, error) {
	var submit ClinicOutcomesFormSubmit
	var errors ErrorSubmit

	// EventDetails

	var dateTime time.Time
	var err error

	if state.Details.EventDate == "" {
		errors.Errors = append(errors.Errors, "Enter the event date")
	}
	if state.Details.EventTime == "" {
		errors.Errors = append(errors.Errors, "Enter the event time")
	}
	if state.Details.EventDate != "" && state.Details.EventTime != "" {
		dateString := fmt.Sprintf("%s %s", state.Details.EventDate, state.Details.EventTime)
		dateTime, err = time.Parse("2006-01-02 15:04", dateString)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return ClinicOutcomesFormSubmit{}, err
		}
	}

	submit.Details.DateTime = dateTime

	if state.Details.EventType == "Choose Option" {
		errors.Errors = append(errors.Errors, "Select consultation type")
	} else {
		submit.Details.Type = state.Details.EventType
	}

	if state.Details.EventSpecialty == "Choose Option" {
		errors.Errors = append(errors.Errors, "Select specialty")
	} else {
		submit.Details.Specialty = state.Details.EventSpecialty
	}

	if state.Details.EventClinician == "Choose Option" {
		errors.Errors = append(errors.Errors, "Select senior responsible clinician")
	} else {
		submit.Details.Clinician = state.Details.EventClinician
	}

	// CancerPathway

	switch {
	case !state.CancerPathway.Checked:
		submit.CancerPathway = ""
	case state.CancerPathway.Option == "Other":
		submit.CancerPathway = fmt.Sprintf("Other: %s", state.CancerPathway.Other)
	default:
		submit.CancerPathway = state.CancerPathway.Option
	}

	// Outcome
	submit.Outcome.Answer = state.Outcome.OutcomeOption

	switch submit.Outcome.Answer {
	case "See on Symptom":
		submit.Outcome.AnswerDetails = state.Outcome.SeeOnSymptomDetails
	case "Did Not Answer":
		submit.Outcome.AnswerDetails = state.Outcome.DidNotAnswerDetails
	case "Did Not Attend":
		submit.Outcome.AnswerDetails = state.Outcome.DidNotAttendDetails
	case "Could Not Attend":
		submit.Outcome.AnswerDetails = state.Outcome.CouldNotAttendDetails
	case "Refer to Diagnostics":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToDiagnosticsDetails
	case "Refer to another consultant / specialty":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToAnotherDetails
	case "Refer to Therapies":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToTherapiesDetails
	case "Refer to Treatment":
		ans := ""
		if state.Outcome.ReferToTreatmentSact {
			ans += "SACT, "
		}
		if state.Outcome.ReferToTreatmentRadiotherapy {
			ans += "Radiotherapy, "
		}
		if state.Outcome.ReferToTreatmentOther {
			ans += "Other, "
		}
		ans = strings.TrimSuffix(ans, ", ")
		ans += ": " + state.Outcome.ReferToTreatmentDetails
		submit.Outcome.AnswerDetails = ans
	case "Discuss at MDT":
		submit.Outcome.AnswerDetails = state.Outcome.DiscussAtMdtDetails
	case "Listed for Outpatient Procedure":
		submit.Outcome.AnswerDetails = state.Outcome.OutpatientProcedureDetails
	}

	// FollowUp
	submit.FollowUpRequired = state.FollowUp.Checked

	if submit.FollowUpRequired {
		switch {
		case state.FollowUp.Pathway == "":
			errors.Errors = append(errors.Errors, "Select follow up pathway")
		default:
			submit.FollowUp.Pathway = state.FollowUp.Pathway
		}

		switch {
		case state.FollowUp.SameClinician == "":
			errors.Errors = append(errors.Errors, "Select follow up clinician")
		case state.FollowUp.SameClinician == "No":
			submit.FollowUp.SameClinician = fmt.Sprintf("No: %s", state.FollowUp.SameClinicianNo)
		default:
			submit.FollowUp.SameClinician = state.FollowUp.SameClinician
		}

		switch {
		case state.FollowUp.SameClinic == "No":
			submit.FollowUp.SameClinic = fmt.Sprintf("No: %s", state.FollowUp.SameClinicNo)
		default:
			submit.FollowUp.SameClinic = state.FollowUp.SameClinic
		}

		switch {
		case state.FollowUp.SeeInNum == "":
			errors.Errors = append(errors.Errors, "Select follow up date")
		default:
			submit.FollowUp.SeeIn = state.FollowUp.SeeInNum + " " + state.FollowUp.SeeInUnit
			seeInNum, err := strconv.Atoi(state.FollowUp.SeeInNum)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return ClinicOutcomesFormSubmit{}, err
			}
			switch state.FollowUp.SeeInUnit {
			case "Weeks":
				submit.FollowUp.DateTime = time.Now().AddDate(0, 0, seeInNum*7)
			case "Months":
				submit.FollowUp.DateTime = time.Now().AddDate(0, seeInNum, 0)
			case "Years":
				submit.FollowUp.DateTime = time.Now().AddDate(seeInNum, 0, 0)
			}
		}

		submit.FollowUp.Hospital = state.FollowUp.Hospital
		switch {
		case state.FollowUp.AppointmentPriority == "":
			errors.Errors = append(errors.Errors, "Enter the appointment priority")
		default:
			submit.FollowUp.AppointmentPriority = state.FollowUp.AppointmentPriority
		}
		submit.FollowUp.ClinicalCondition = state.FollowUp.Condition

		switch {
		case state.FollowUp.PreferredConsultationMethod == "":
			errors.Errors = append(errors.Errors, "Select the preferred consultation method")
		default:
			submit.FollowUp.PreferredConsultationMethod = state.FollowUp.PreferredConsultationMethod
		}

		if state.FollowUp.TestsRequiredBeforeFollowup == "Yes" {
			followUpError := false
			for _, test := range state.FollowUp.Tests {
				if test.TestsRequired == "" {
					followUpError = true
					break
				}
				if test.TestsUndertakenBy == "" {
					followUpError = true
					break
				}
				if test.TestsRequiredBy == "Choose Option" {
					followUpError = true
					break
				}
			}
			if followUpError {
				errors.Errors = append(errors.Errors, "All test request fields are required")
			}
			if !followUpError {
				submit.FollowUp.Tests = state.FollowUp.Tests
			}
		}
	}

	// OtherInformation

	submit.OtherInformation = state.OtherInformation

	if len(errors.Errors) > 0 {
		return ClinicOutcomesFormSubmit{}, errors
	} else {
		return submit, nil
	}
}
