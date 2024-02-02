package models

import (
	"encoding/json"
)

type ClinicOutcomesFormPayload struct {
	EventDate      string `json:"event_date"`
	EventTime      string `json:"event_time"`
	EventType      string `json:"event_type"`
	EventSpecialty string `json:"event_specialty"`
	EventClinician string `json:"event_clinician"`

	CancerPathway       string `json:"cancer_pathway"`
	CancerPathwayOption string `json:"cancer_pathway_option"`
	CancerPathwayOther  string `json:"cancer_pathway_other"`

	OutcomeOption                string `json:"outcome_option"`
	SeeOnSymptomDetails          string `json:"see_on_symptom_details"`
	DidNotAnswerDetails          string `json:"did_not_answer_details"`
	DidNotAttendDetails          string `json:"did_not_attend_details"`
	CouldNotAttendDetails        string `json:"could_not_attend_details"`
	ReferToDiagnosticsDetails    string `json:"refer_to_diagnostics_details"`
	ReferToAnotherDetails        string `json:"refer_to_another_details"`
	ReferToTherapiesDetails      string `json:"refer_to_therapies_details"`
	ReferToTreatmentSact         string `json:"refer_to_treatment_sact"`
	ReferToTreatmentRadiotherapy string `json:"refer_to_treatment_radiotherapy"`
	ReferToTreatmentOther        string `json:"refer_to_treatment_other"`
	ReferToTreatmentDetails      string `json:"refer_to_treatment_details"`
	DiscussAtMdtDetails          string `json:"discuss_at_mdt_details"`
	OutpatientProcedureDetails   string `json:"outpatient_procedure_details"`

	FollowUp                    string `json:"follow_up"`
	Pathway                     string `json:"pathway"`
	SameClinician               string `json:"same_clinician"`
	SameClinicianNo             string `json:"same_clinician_no"`
	SameClinic                  string `json:"same_clinic"`
	SameClinicNo                string `json:"same_clinic_no"`
	SeeInNum                    string `json:"see_in_num"`
	SeeInUnit                   string `json:"see_in_unit"`
	Hospital                    string `json:"hospital"`
	AppointmentPriority         string `json:"appointment_priority"`
	Condition                   string `json:"condition"`
	PreferredConsultationMethod string `json:"preferred_consultation_method"`

	TestsRequiredBeforeFollowup string      `json:"tests_required_before_followup"`
	TestsRequired               ArrayString `json:"tests_required"`
	TestsUndertakenBy           ArrayString `json:"tests.undertaken"`
	TestsBy                     ArrayString `json:"tests.by"`
	AddTest                     *string     `json:"add_test"`

	OtherInformation string `json:"other_information"`
}

type ArrayString []string

func (a *ArrayString) UnmarshalJSON(b []byte) error {
	var v interface{}
	json.Unmarshal(b, &v)
	switch res := v.(type) {
	case string:
		*a = []string{res}
	default:
		return json.Unmarshal(b, (*[]string)(a))
	}
	return nil
}
