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

func Submit(state ClinicOutcomesFormState) (ClinicOutcomesFormSubmit, error) {
	var submit ClinicOutcomesFormSubmit

	// EventDetails
	dateString := fmt.Sprintf("%s %s", state.Details.EventDate, state.Details.EventTime)
	dateTime, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ClinicOutcomesFormSubmit{}, err
	}

	submit.Details.DateTime = dateTime
	submit.Details.Type = state.Details.EventType
	submit.Details.Specialty = state.Details.EventSpecialty
	submit.Details.Clinician = state.Details.EventClinician

	// CancerPathway
	if !state.CancerPathway.Checked {
		submit.CancerPathway = "NA"
	} else if state.CancerPathway.Option == "Other" {
		submit.CancerPathway = fmt.Sprintf("Other: %s", state.CancerPathway.Other)
	} else {
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
		if state.Outcome.ReferToTreatmentSact == "on" {
			ans += "SACT "
		}
		if state.Outcome.ReferToTreatmentRadiotherapy == "on" {
			ans += "Radiotherapy "
		}
		if state.Outcome.ReferToTreatmentOther == "on" {
			ans += fmt.Sprintf("Other: %s", state.Outcome.ReferToTreatmentDetails)
		}
		submit.Outcome.AnswerDetails = strings.TrimSuffix(ans, " ")
	case "Refer to treatment - SACT":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToTreatmentSact
	case "Refer to treatment - Radiotherapy":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToTreatmentRadiotherapy
	case "Refer to treatment - Other":
		submit.Outcome.AnswerDetails = state.Outcome.ReferToTreatmentOther
	case "Discuss at MDT":
		submit.Outcome.AnswerDetails = state.Outcome.DiscussAtMdtDetails
	case "Listed for Outpatient Procedure":
		submit.Outcome.AnswerDetails = state.Outcome.OutpatientProcedureDetails
	default:
		submit.Outcome.AnswerDetails = "NA"
	}

	// FollowUp
	submit.FollowUpRequired = state.FollowUp.FollowUp == "on"

	if submit.FollowUpRequired {
		submit.FollowUp.Pathway = state.FollowUp.Pathway

		if state.FollowUp.SameClinician == "No" {
			submit.FollowUp.SameClinician = fmt.Sprintf("No: %s", state.FollowUp.SameClinicianNo)
		} else {
			submit.FollowUp.SameClinician = state.FollowUp.SameClinician
		}

		if state.FollowUp.SameClinic == "No" {
			submit.FollowUp.SameClinic = fmt.Sprintf("No: %s", state.FollowUp.SameClinicNo)
		} else {
			submit.FollowUp.SameClinic = state.FollowUp.SameClinic
		}

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

		submit.FollowUp.Hospital = state.FollowUp.Hospital
		submit.FollowUp.AppointmentPriority = state.FollowUp.AppointmentPriority
		submit.FollowUp.ClinicalCondition = state.FollowUp.Condition

		submit.FollowUp.PreferredConsultationMethod = state.FollowUp.PreferredConsultationMethod
		if state.FollowUp.TestsRequiredBeforeFollowup == "Yes" {
			submit.FollowUp.Tests = state.FollowUp.Tests
			// TODO(viv): check if all fields populated before adding
		}
	}

	// OtherInformation
	submit.OtherInformation = state.OtherInformation
	return submit, nil
}
