package models

import (
	"encoding/json"
	"log"
)

type OutcomesForm struct {
	AddTest          *string `json:"add_test"`
	ConsultationDate string  `json:"consultation_date"`
	ConsultationTime string  `json:"consultation_time"`
	ConsultationType string  `json:"consultation_type"`
	Specialties      string  `json:"specialties"`
	Clinicians       string  `json:"clinicians"`

	CancerPathway       string `json:"cancer_pathway"`
	CancerPathwayOption string `json:"cancer_pathway_options"`
	CancerPathwayOther  string `json:"cancer_pathway_other"`

	PatientOption      string `json:"patient_option"`
	SeeOnSymptomMonths string `json:"see_on_symptom_months"`

	DidNotAnswer   string `json:"did_not_answer"`
	DidNotAttend   string `json:"did_not_attend"`
	CouldNotAttend string `json:"could_not_attend"`

	ReferToDiagnosticsData string `json:"refer_to_diagnostics_data"`
	ReferToAnotherData     string `json:"refer_to_another_data"`

	ReferToTreatment string `json:"refer_to_treatment"`

	ReferRoTreatmentSact         string `json:"refer_to_treatment_sact"`
	ReferRoTreatmentRadiotherapy string `json:"refer_to_treatment_radiotherapy"`
	ReferRoTreatmentOther        string `json:"refer_to_treatment_other"`
	ReferRoTreatmentData         string `json:"refer_to_treatment_data"`
	ReferToTherapyData           string `json:"refer_to_therapy_data"`

	DiscussAtMdtData        string `json:"discuss_at_mdt_data"`
	OutpatientProcedureData string `json:"outpatient_procedure_data"`

	FollowUp            string `json:"follow_up"`
	Pathway             string `json:"pathway"`
	SameClinician       string `json:"same_clinician"`
	SameClinicianAnswer string `json:"same_clinician_answer"`

	SeeInUnit string `json:"see_in_unit"`
	SeeInNum  string `json:"see_in_num"`

	SameClinic       string `json:"same_clinic"`
	SameClinicAnswer string `json:"same_clinic_answer"`

	Hospital                    string `json:"hospital"`
	Condition                   string `json:"condition"`
	AppointmentDP               string `json:"appointment_dp"`
	PreferredConsultationMethod string `json:"preferred_consultation_method"`
	TestsRequired               string `json:"tests_required"`
	OtherInformation            string `json:"other_information"`

	FollowUpTestsRequired   ArrayString `json:"tests.required"`
	FollowUpTestsUndertaken ArrayString `json:"tests.undertaken"`
	FollowUpTestsBy         ArrayString `json:"tests.by"`
}

type ArrayString []string

func (a *ArrayString) UnmarshalJSON(b []byte) error {
	log.Println(string(b))
	if len(b) == 0 || string(b) == `""` {
		log.Println(string(b), "is nil or empty")
		*a = []string{}
		return nil
	}
	if json.Valid(b) {
		log.Println(string(b), "is valid json")
		return json.Unmarshal(b, (*[]string)(a))
	}
	*a = []string{string(b)}
	return nil
}

type OutcomesState struct {
	ConsultationDate string
	ConsultationTime string
	ConsultationType string
	Specialties      string
	Clinicians       string

	CancerPathway CancerPathway

	PatientOption      string
	SeeOnSymptomMonths string

	DidNotAnswer   string
	DidNotAttend   string
	CouldNotAttend string

	ReferToDiagnosticsData string
	ReferToAnotherData     string
	ReferToTherapyData     string

	DiscussAtMdtData             string
	ReferToTreatment             string
	ReferRoTreatmentSact         string
	ReferRoTreatmentRadiotherapy string
	ReferRoTreatmentOther        string
	ReferRoTreatmentData         string
	OutpatientProcedureData      string

	FollowUp            string
	Pathway             string
	SameClinician       string
	SameClinicianAnswer string

	SeeInUnit     string
	SeeInNum      string
	Hospital      string
	Condition     string
	AppointmentDP string

	SameClinic       string
	SameClinicAnswer string

	PreferredConsultationMethod string
	TestsRequired               string
	Tests                       []Test
	OtherInformation            string
}

type CancerPathway struct {
	Checked bool
	Option  string
	Other   string
}

type Test struct {
	TestsRequired   string
	UndertakenBy    string
	TestsRequiredBy string
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
		ConsultationDate: o.ConsultationDate,
		ConsultationTime: o.ConsultationTime,
		ConsultationType: o.ConsultationType,
		Specialties:      o.Specialties,
		Clinicians:       o.Clinicians,

		CancerPathway: CancerPathway{
			Checked: o.CancerPathway == "on",
			Option:  o.CancerPathwayOption,
			Other:   o.CancerPathwayOther,
		},

		PatientOption:      o.PatientOption,
		SeeOnSymptomMonths: o.SeeOnSymptomMonths,

		DidNotAnswer:   o.DidNotAnswer,
		DidNotAttend:   o.DidNotAttend,
		CouldNotAttend: o.CouldNotAttend,

		ReferToDiagnosticsData: o.ReferToDiagnosticsData,
		ReferToAnotherData:     o.ReferToAnotherData,
		ReferToTreatment:       o.ReferToTreatment,

		ReferRoTreatmentSact:         o.ReferRoTreatmentSact,
		ReferRoTreatmentRadiotherapy: o.ReferRoTreatmentRadiotherapy,
		ReferRoTreatmentOther:        o.ReferRoTreatmentOther,
		ReferRoTreatmentData:         o.ReferRoTreatmentData,
		ReferToTherapyData:           o.ReferToTherapyData,

		DiscussAtMdtData:        o.DiscussAtMdtData,
		OutpatientProcedureData: o.OutpatientProcedureData,

		FollowUp:            o.FollowUp,
		Pathway:             o.Pathway,
		SameClinician:       o.SameClinician,
		SameClinicianAnswer: o.SameClinicianAnswer,

		SeeInUnit:     o.SeeInUnit,
		SeeInNum:      o.SeeInNum,
		Hospital:      o.Hospital,
		Condition:     o.Condition,
		AppointmentDP: o.AppointmentDP,

		SameClinic:       o.SameClinic,
		SameClinicAnswer: o.SameClinicAnswer,

		PreferredConsultationMethod: o.PreferredConsultationMethod,
		TestsRequired:               o.TestsRequired,
		OtherInformation:            o.OtherInformation,
		Tests:                       tests,
	}
}
