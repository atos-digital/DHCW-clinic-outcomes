package models

type OutcomesForm struct {
	ConsultationDate string `json:"consultation_date"`
	ConsultationTime string `json:"consultation_time"`
	ConsultationType string `json:"consultation_type"`
	Specialties      string `json:"specialties"`
	Clinicians       string `json:"clinicians"`

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

	ReferToTherapyData string `json:"refer_to_therapy_data"`

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
	}
}
