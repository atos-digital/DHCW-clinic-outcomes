package models

import "time"

type ClinicOutcomesFormSubmit struct {
	DetailsSubmit    DetailsSubmit
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
