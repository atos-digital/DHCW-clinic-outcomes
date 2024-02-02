package models

import "time"

type OutcomesSubmit struct {
	EventDetails     EventDetailsAns
	CancerPathway    string
	Outcome          OutcomeAns
	FollowUpRequired bool
	FollowUp         FollowUpAns
	OtherInformation string
}

type EventDetailsAns struct {
	DateTime                   time.Time
	Type                       string
	Specialty                  string
	SeniorResponsibleClinician string
}

type OutcomeAns struct {
	Answer         string
	FollowUpAnswer string
}

type FollowUpAns struct {
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
