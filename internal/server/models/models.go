package models

import (
	"encoding/json"
)

type OutcomesForm struct {
	AddTest                      *string     `json:"add_test"`
	ConsultationDate             string      `json:"consultation_date"`
	ConsultationTime             string      `json:"consultation_time"`
	ConsultationType             string      `json:"consultation_type"`
	Specialties                  string      `json:"specialties"`
	Clinicians                   string      `json:"clinicians"`
	CancerPathway                string      `json:"cancer_pathway"`
	CancerPathwayOption          string      `json:"cancer_pathway_options"`
	CancerPathwayOther           string      `json:"cancer_pathway_other"`
	PatientOption                string      `json:"patient_option"`
	SeeOnSymptomMonths           string      `json:"see_on_symptom_months"`
	DidNotAnswer                 string      `json:"did_not_answer"`
	DidNotAttend                 string      `json:"did_not_attend"`
	CouldNotAttend               string      `json:"could_not_attend"`
	ReferToDiagnosticsData       string      `json:"refer_to_diagnostics_data"`
	ReferToAnotherData           string      `json:"refer_to_another_data"`
	ReferRoTreatmentSact         string      `json:"refer_to_treatment_sact"`
	ReferRoTreatmentRadiotherapy string      `json:"refer_to_treatment_radiotherapy"`
	ReferRoTreatmentOther        string      `json:"refer_to_treatment_other"`
	ReferRoTreatmentData         string      `json:"refer_to_treatment_data"`
	ReferToTherapyData           string      `json:"refer_to_therapy_data"`
	DiscussAtMdtData             string      `json:"discuss_at_mdt_data"`
	OutpatientProcedureData      string      `json:"outpatient_procedure_data"`
	FollowUp                     string      `json:"follow_up"`
	Pathway                      string      `json:"pathway"`
	SameClinician                string      `json:"same_clinician"`
	SameClinicianAnswer          string      `json:"same_clinician_answer"`
	SeeInUnit                    string      `json:"see_in_unit"`
	SeeInNum                     string      `json:"see_in_num"`
	SameClinic                   string      `json:"same_clinic"`
	SameClinicAnswer             string      `json:"same_clinic_answer"`
	Hospital                     string      `json:"hospital"`
	Condition                    string      `json:"condition"`
	AppointmentDP                string      `json:"appointment_dp"`
	PreferredConsultationMethod  string      `json:"preferred_consultation_method"`
	TestsRequired                string      `json:"tests_required"`
	FollowUpTestsRequired        ArrayString `json:"tests.required"`
	FollowUpTestsUndertaken      ArrayString `json:"tests.undertaken"`
	FollowUpTestsBy              ArrayString `json:"tests.by"`
	OtherInformation             string      `json:"other_information"`
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
			ConsultationDate: o.ConsultationDate,
			ConsultationTime: o.ConsultationTime,
			ConsultationType: o.ConsultationType,
			Specialties:      o.Specialties,
			Clinicians:       o.Clinicians,
		},
		CancerPathway: CancerPathway{
			Checked: o.CancerPathway == "on",
			Option:  o.CancerPathwayOption,
			Other:   o.CancerPathwayOther,
		},
		OutcomeOptions: OutcomeOptions{
			PatientOption:      o.PatientOption,
			SeeOnSymptomMonths: o.SeeOnSymptomMonths,

			DidNotAnswer:   o.DidNotAnswer,
			DidNotAttend:   o.DidNotAttend,
			CouldNotAttend: o.CouldNotAttend,

			ReferToDiagnosticsData: o.ReferToDiagnosticsData,
			ReferToAnotherData:     o.ReferToAnotherData,
			ReferToTherapyData:     o.ReferToTherapyData,

			ReferRoTreatmentSact:         o.ReferRoTreatmentSact,
			ReferRoTreatmentRadiotherapy: o.ReferRoTreatmentRadiotherapy,
			ReferRoTreatmentOther:        o.ReferRoTreatmentOther,
			ReferRoTreatmentData:         o.ReferRoTreatmentData,

			DiscussAtMdtData:        o.DiscussAtMdtData,
			OutpatientProcedureData: o.OutpatientProcedureData,
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
