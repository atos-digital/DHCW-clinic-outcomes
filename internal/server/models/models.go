package models

import (
	"encoding/json"
)

type OutcomesForm struct {
	EventDate      string `json:"event_date"`
	EventTime      string `json:"event_time"`
	EventType      string `json:"event_type"`
	EventSpecialty string `json:"event_specialty"`
	EventClinician string `json:"event_clinician"`

	CancerPathway       string `json:"cancer_pathway"`
	CancerPathwayOption string `json:"cancer_pathway_option"`
	CancerPathwayOther  string `json:"cancer_pathway_other"`

	PatientOption                string `json:"patient_option"`
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
	SameClinicianAnswer         string `json:"same_clinician_answer"`
	SeeInUnit                   string `json:"see_in_unit"`
	SeeInNum                    string `json:"see_in_num"`
	SameClinic                  string `json:"same_clinic"`
	SameClinicAnswer            string `json:"same_clinic_answer"`
	Hospital                    string `json:"hospital"`
	Condition                   string `json:"condition"`
	AppointmentDP               string `json:"appointment_dp"`
	PreferredConsultationMethod string `json:"preferred_consultation_method"`

	TestsRequired           string      `json:"tests_required"`
	FollowUpTestsRequired   ArrayString `json:"tests.required"`
	FollowUpTestsUndertaken ArrayString `json:"tests.undertaken"`
	FollowUpTestsBy         ArrayString `json:"tests.by"`
	AddTest                 *string     `json:"add_test"`

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

func (o OutcomesForm) State() OutcomesState {
	num := len(o.FollowUpTestsRequired)
	if num == 0 {
		num = 1
	}
	tests := make([]Test, num)
	for i, v := range o.FollowUpTestsRequired {
		tests[i] = Test{
			TestsRequired:   v,
			UndertakenBy:    o.FollowUpTestsUndertaken[i],
			TestsRequiredBy: o.FollowUpTestsBy[i],
		}
	}

	return OutcomesState{
		OutcomeDetails: OutcomeDetails{
			EventDate:      o.EventDate,
			EventTime:      o.EventTime,
			EventType:      o.EventType,
			EventSpecialty: o.EventSpecialty,
			EventClinician: o.EventClinician,
		},
		CancerPathway: CancerPathway{
			Checked: o.CancerPathway == "on",
			Option:  o.CancerPathwayOption,
			Other:   o.CancerPathwayOther,
		},
		OutcomeOptions: OutcomeOptions{
			PatientOption:       o.PatientOption,
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
		FollowUp: FollowUp{
			FollowUp:                    o.FollowUp,
			Pathway:                     o.Pathway,
			SameClinician:               o.SameClinician,
			SameClinicianAnswer:         o.SameClinicianAnswer,
			SameClinic:                  o.SameClinic,
			SameClinicAnswer:            o.SameClinicAnswer,
			SeeInUnit:                   o.SeeInUnit,
			SeeInNum:                    o.SeeInNum,
			Hospital:                    o.Hospital,
			AppointmentDP:               o.AppointmentDP,
			Condition:                   o.Condition,
			PreferredConsultationMethod: o.PreferredConsultationMethod,
			TestsRequired:               o.TestsRequired,
			Tests:                       tests,
		},
		OtherInformation: o.OtherInformation,
	}
}
