package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ClinicOutcomesFormState struct {
	Details          DetailsState
	CancerPathway    CancerPathwayState
	Outcome          OutcomeState
	FollowUp         FollowUpState
	OtherInformation string
}

type DetailsState struct {
	EventDate      string
	EventTime      string
	EventType      string
	EventSpecialty string
	EventClinician string
}

type CancerPathwayState struct {
	Checked bool
	Option  string
	Other   string
}

type OutcomeState struct {
	OutcomeOption                string
	SeeOnSymptomDetails          string
	DidNotAnswerDetails          string
	DidNotAttendDetails          string
	CouldNotAttendDetails        string
	ReferToDiagnosticsDetails    string
	ReferToAnotherDetails        string
	ReferToTherapiesDetails      string
	ReferToTreatment             string
	ReferToTreatmentSact         string
	ReferToTreatmentRadiotherapy string
	ReferToTreatmentOther        string
	ReferToTreatmentDetails      string
	DiscussAtMdtDetails          string
	OutpatientProcedureDetails   string
}

type Test struct {
	TestsRequired     string
	TestsUndertakenBy string
	TestsRequiredBy   string
}

type FollowUpState struct {
	FollowUp                    string
	Pathway                     string
	SameClinician               string
	SameClinicianNo             string
	SameClinic                  string
	SameClinicNo                string
	SeeInNum                    string
	SeeInUnit                   string
	Hospital                    string
	AppointmentPriority         string
	Condition                   string
	PreferredConsultationMethod string
	TestsRequiredBeforeFollowup string
	Tests                       []Test
}

func (o ClinicOutcomesFormState) Submit() (ClinicOutcomesFormSubmit, error) {
	var submit ClinicOutcomesFormSubmit

	// EventDetails
	dateString := fmt.Sprintf("%s %s", o.Details.EventDate, o.Details.EventTime)
	dateTime, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ClinicOutcomesFormSubmit{}, err
	}

	submit.DetailsSubmit.DateTime = dateTime
	submit.DetailsSubmit.Type = o.Details.EventType
	submit.DetailsSubmit.Specialty = o.Details.EventSpecialty
	submit.DetailsSubmit.Clinician = o.Details.EventClinician

	// CancerPathway
	if !o.CancerPathway.Checked {
		submit.CancerPathway = "NA"
	} else if o.CancerPathway.Option == "Other" {
		submit.CancerPathway = fmt.Sprintf("Other: %s", o.CancerPathway.Other)
	} else {
		submit.CancerPathway = o.CancerPathway.Option
	}

	// Outcome
	submit.Outcome.Answer = o.Outcome.OutcomeOption

	switch submit.Outcome.Answer {
	case "See on Symptom":
		submit.Outcome.AnswerDetails = o.Outcome.SeeOnSymptomDetails
	case "Did Not Answer":
		submit.Outcome.AnswerDetails = o.Outcome.DidNotAnswerDetails
	case "Did Not Attend":
		submit.Outcome.AnswerDetails = o.Outcome.DidNotAttendDetails
	case "Could Not Attend":
		submit.Outcome.AnswerDetails = o.Outcome.CouldNotAttendDetails
	case "Refer to Diagnostics":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToDiagnosticsDetails
	case "Refer to another consultant / specialty":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToAnotherDetails
	case "Refer to Therapies":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToTherapiesDetails
	case "Refer to Treatment":
		ans := ""
		if o.Outcome.ReferToTreatmentSact == "on" {
			ans += "SACT "
		}
		if o.Outcome.ReferToTreatmentRadiotherapy == "on" {
			ans += "Radiotherapy "
		}
		if o.Outcome.ReferToTreatmentOther == "on" {
			ans += fmt.Sprintf("Other: %s", o.Outcome.ReferToTreatmentDetails)
		}
		submit.Outcome.AnswerDetails = strings.TrimSuffix(ans, " ")
	case "Refer to treatment - SACT":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToTreatmentSact
	case "Refer to treatment - Radiotherapy":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToTreatmentRadiotherapy
	case "Refer to treatment - Other":
		submit.Outcome.AnswerDetails = o.Outcome.ReferToTreatmentOther
	case "Discuss at MDT":
		submit.Outcome.AnswerDetails = o.Outcome.DiscussAtMdtDetails
	case "Listed for Outpatient Procedure":
		submit.Outcome.AnswerDetails = o.Outcome.OutpatientProcedureDetails
	default:
		submit.Outcome.AnswerDetails = "NA"
	}

	// FollowUp
	submit.FollowUpRequired = o.FollowUp.FollowUp == "on"

	if submit.FollowUpRequired {
		submit.FollowUp.Pathway = o.FollowUp.Pathway

		if o.FollowUp.SameClinician == "No" {
			submit.FollowUp.SameClinician = fmt.Sprintf("No: %s", o.FollowUp.SameClinicianNo)
		} else {
			submit.FollowUp.SameClinician = o.FollowUp.SameClinician
		}

		if o.FollowUp.SameClinic == "No" {
			submit.FollowUp.SameClinic = fmt.Sprintf("No: %s", o.FollowUp.SameClinicNo)
		} else {
			submit.FollowUp.SameClinic = o.FollowUp.SameClinic
		}

		submit.FollowUp.SeeIn = o.FollowUp.SeeInNum + " " + o.FollowUp.SeeInUnit
		seeInNum, err := strconv.Atoi(o.FollowUp.SeeInNum)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return ClinicOutcomesFormSubmit{}, err
		}
		switch o.FollowUp.SeeInUnit {
		case "Weeks":
			submit.FollowUp.DateTime = time.Now().AddDate(0, 0, seeInNum*7)
		case "Months":
			submit.FollowUp.DateTime = time.Now().AddDate(0, seeInNum, 0)
		case "Years":
			submit.FollowUp.DateTime = time.Now().AddDate(seeInNum, 0, 0)
		}

		submit.FollowUp.Hospital = o.FollowUp.Hospital
		submit.FollowUp.AppointmentPriority = o.FollowUp.AppointmentPriority
		submit.FollowUp.ClinicalCondition = o.FollowUp.Condition

		submit.FollowUp.PreferredConsultationMethod = o.FollowUp.PreferredConsultationMethod
		if o.FollowUp.TestsRequiredBeforeFollowup == "Yes" {
			submit.FollowUp.Tests = o.FollowUp.Tests
			// TODO(viv): check if all fields populated before adding
		}
	}

	// OtherInformation
	submit.OtherInformation = o.OtherInformation
	return submit, nil
}
