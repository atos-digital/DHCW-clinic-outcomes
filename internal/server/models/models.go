package models

import (
	"encoding/json"
)

type ClinicOutcomesForm struct {
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
	TestsUndertakenBy           ArrayString `json:"tests_undertaken"`
	TestsBy                     ArrayString `json:"tests_by"`
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

func (o ClinicOutcomesForm) State() ClinicOutcomesFormState {
	num := len(o.TestsRequired)
	if num == 0 {
		num = 1
	}
	tests := make([]Test, num)
	for i, v := range o.TestsRequired {
		tests[i] = Test{
			TestsRequired:     v,
			TestsUndertakenBy: o.TestsUndertakenBy[i],
			TestsRequiredBy:   o.TestsBy[i],
		}
	}

	return ClinicOutcomesFormState{
		Details: DetailsState{
			EventDate:      o.EventDate,
			EventTime:      o.EventTime,
			EventType:      o.EventType,
			EventSpecialty: o.EventSpecialty,
			EventClinician: o.EventClinician,
		},
		CancerPathway: CancerPathwayState{
			Checked: o.CancerPathway == "on",
			Option:  o.CancerPathwayOption,
			Other:   o.CancerPathwayOther,
		},
		Outcome: OutcomeState{
			OutcomeOption:       o.OutcomeOption,
			SeeOnSymptomDetails: o.SeeOnSymptomDetails,

			DidNotAnswerDetails:   o.DidNotAnswerDetails,
			DidNotAttendDetails:   o.DidNotAttendDetails,
			CouldNotAttendDetails: o.CouldNotAttendDetails,

			ReferToDiagnosticsDetails: o.ReferToDiagnosticsDetails,
			ReferToAnotherDetails:     o.ReferToAnotherDetails,
			ReferToTherapiesDetails:   o.ReferToTherapiesDetails,

			ReferToTreatmentSact:         o.ReferToTreatmentSact,
			ReferToTreatmentRadiotherapy: o.ReferToTreatmentRadiotherapy,
			ReferToTreatmentOther:        o.ReferToTreatmentOther,
			ReferToTreatmentDetails:      o.ReferToTreatmentDetails,

			DiscussAtMdtDetails:        o.DiscussAtMdtDetails,
			OutpatientProcedureDetails: o.OutpatientProcedureDetails,
		},
		FollowUp: FollowUpState{
			FollowUp:                    o.FollowUp,
			Pathway:                     o.Pathway,
			SameClinician:               o.SameClinician,
			SameClinicianNo:             o.SameClinicianNo,
			SameClinic:                  o.SameClinic,
			SameClinicNo:                o.SameClinicNo,
			SeeInUnit:                   o.SeeInUnit,
			SeeInNum:                    o.SeeInNum,
			Hospital:                    o.Hospital,
			AppointmentPriority:         o.AppointmentPriority,
			Condition:                   o.Condition,
			PreferredConsultationMethod: o.PreferredConsultationMethod,
			TestsRequiredBeforeFollowup: o.TestsRequiredBeforeFollowup,
			Tests:                       tests,
		},
		OtherInformation: o.OtherInformation,
	}
}
