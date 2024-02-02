package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type OutcomeDetails struct {
	EventDate      string
	EventTime      string
	EventType      string
	EventSpecialty string
	EventClinician string
}
type CancerPathway struct {
	Checked bool
	Option  string
	Other   string
}

type OutcomeOptions struct {
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
	TestsRequired   string
	UndertakenBy    string
	TestsRequiredBy string
}

type FollowUp struct {
	FollowUp                    string
	Pathway                     string
	SameClinician               string
	SameClinicianAnswer         string
	SameClinic                  string
	SameClinicAnswer            string
	SeeInUnit                   string
	SeeInNum                    string
	Hospital                    string
	AppointmentDP               string
	Condition                   string
	PreferredConsultationMethod string
	TestsRequired               string
	Tests                       []Test
}

type OutcomesState struct {
	OutcomeDetails   OutcomeDetails
	CancerPathway    CancerPathway
	OutcomeOptions   OutcomeOptions
	FollowUp         FollowUp
	OtherInformation string
}

func (o OutcomesState) Submit() (OutcomesSubmit, error) {
	var submit OutcomesSubmit

	// EventDetails
	dateString := fmt.Sprintf("%s %s", o.OutcomeDetails.EventDate, o.OutcomeDetails.EventTime)
	dateTime, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return OutcomesSubmit{}, err
	}

	submit.EventDetails.DateTime = dateTime
	submit.EventDetails.Type = o.OutcomeDetails.EventType
	submit.EventDetails.Specialty = o.OutcomeDetails.EventSpecialty
	submit.EventDetails.SeniorResponsibleClinician = o.OutcomeDetails.EventClinician

	// CancerPathway
	if !o.CancerPathway.Checked {
		submit.CancerPathway = "NA"
	} else if o.CancerPathway.Option == "Other" {
		submit.CancerPathway = fmt.Sprintf("Other: %s", o.CancerPathway.Other)
	} else {
		submit.CancerPathway = o.CancerPathway.Option
	}

	// Outcome
	submit.Outcome.Answer = o.OutcomeOptions.OutcomeOption

	switch submit.Outcome.Answer {
	case "See on Symptom":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.SeeOnSymptomDetails
	case "Did Not Answer":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.DidNotAnswerDetails
	case "Did Not Attend":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.DidNotAttendDetails
	case "Could Not Attend":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.CouldNotAttendDetails
	case "Refer to Diagnostics":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToDiagnosticsDetails
	case "Refer to another consultant / specialty":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToAnotherDetails
	case "Refer to Therapies":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToTherapiesDetails
	case "Refer to Treatment":
		ans := ""
		if o.OutcomeOptions.ReferToTreatmentSact == "on" {
			ans += "SACT "
		}
		if o.OutcomeOptions.ReferToTreatmentRadiotherapy == "on" {
			ans += "Radiotherapy "
		}
		if o.OutcomeOptions.ReferToTreatmentOther == "on" {
			ans += fmt.Sprintf("Other: %s", o.OutcomeOptions.ReferToTreatmentDetails)
		}
		submit.Outcome.FollowUpAnswer = strings.TrimSuffix(ans, " ")
	case "Refer to treatment - SACT":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToTreatmentSact
	case "Refer to treatment - Radiotherapy":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToTreatmentRadiotherapy
	case "Refer to treatment - Other":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.ReferToTreatmentOther
	case "Discuss at MDT":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.DiscussAtMdtDetails
	case "Listed for Outpatient Procedure":
		submit.Outcome.FollowUpAnswer = o.OutcomeOptions.OutpatientProcedureDetails
	default:
		submit.Outcome.FollowUpAnswer = "NA"
	}

	// FollowUp
	submit.FollowUpRequired = o.FollowUp.FollowUp == "on"

	if submit.FollowUpRequired {
		submit.FollowUp.Pathway = o.FollowUp.Pathway

		if o.FollowUp.SameClinician == "No" {
			submit.FollowUp.SameClinician = fmt.Sprintf("No: %s", o.FollowUp.SameClinicianAnswer)
		} else {
			submit.FollowUp.SameClinician = o.FollowUp.SameClinician
		}

		if o.FollowUp.SameClinic == "No" {
			submit.FollowUp.SameClinic = fmt.Sprintf("No: %s", o.FollowUp.SameClinicAnswer)
		} else {
			submit.FollowUp.SameClinic = o.FollowUp.SameClinic
		}

		submit.FollowUp.SeeIn = o.FollowUp.SeeInNum + " " + o.FollowUp.SeeInUnit
		seeInNum, err := strconv.Atoi(o.FollowUp.SeeInNum)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return OutcomesSubmit{}, err
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
		submit.FollowUp.AppointmentDP = o.FollowUp.AppointmentDP
		submit.FollowUp.ClinicalCondition = o.FollowUp.Condition

		submit.FollowUp.PreferredConsultationMethod = o.FollowUp.PreferredConsultationMethod
		if o.FollowUp.TestsRequired == "Yes" {
			submit.FollowUp.Tests = o.FollowUp.Tests
			// TODO(viv): check if all fields populated before adding
		}
	}

	// OtherInformation
	submit.OtherInformation = o.OtherInformation
	return submit, nil
}
