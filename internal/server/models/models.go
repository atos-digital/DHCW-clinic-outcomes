package models

type OutcomesForm struct {
	ConsultationDate string `json:"consultation_date"`
	ConsultationTime string `json:"consultation_time"`
	ConsultationType string `json:"consultation_type"`
	Specialties      string `json:"specialties"`
	Clinicians       string `json:"clinicians"`

	CancerPathway       string `json:"cancer_pathway"`
	CancerPathwayOption string `json:"cancer_pathway_options"`

	OutcomesOption     string `json:"outcomes_option"`
	SeeOnSymptomMonths string `json:"see_on_symptom_months"`

	DidNotAnswer   string `json:"did_not_answer"`
	DidNotAttend   string `json:"did_not_attend"`
	CouldNotAttend string `json:"could_not_attend"`

	ReferToTreatment string `json:"refer_to_treatment"`

	FollowUp            string `json:"follow_up"`
	Pathway             string `json:"pathway"`
	SameClinician       string `json:"same_clinician"`
	SameClinicianAnswer string `json:"same_clinician_answer"`

	SeeInUnit string `json:"see_in_unit"`
	SeeInNum  string `json:"see_in_num"`

	SameClinic       string `json:"same_clinic"`
	SameClinicAnswer string `json:"same_clinic_answer"`

	PreferredConsultationMethod string `json:"preferred_consultation_method"`
	TestsRequired               string `json:"tests_required"`
}

type OutcomesState struct {
	ConsultationDate string
	ConsultationTime string
	ConsultationType string
	Specialties      string
	Clinicians       string

	CancerPathway CancerPathway

	OutcomesOption     string
	SeeOnSymptomMonths string

	DidNotAnswer   string
	DidNotAttend   string
	CouldNotAttend string

	ReferToTreatment string

	FollowUp            string
	Pathway             string
	SameClinician       string
	SameClinicianAnswer string

	SeeInUnit string
	SeeInNum  string

	SameClinic       string
	SameClinicAnswer string

	PreferredConsultationMethod string
	TestsRequired               string
	Tests                       []Test
}

type CancerPathway struct {
	Checked bool
	Option  string
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
		},

		OutcomesOption:     o.OutcomesOption,
		SeeOnSymptomMonths: o.SeeOnSymptomMonths,

		DidNotAnswer:   o.DidNotAnswer,
		DidNotAttend:   o.DidNotAttend,
		CouldNotAttend: o.CouldNotAttend,

		ReferToTreatment: o.ReferToTreatment,

		FollowUp:            o.FollowUp,
		Pathway:             o.Pathway,
		SameClinician:       o.SameClinician,
		SameClinicianAnswer: o.SameClinicianAnswer,

		SeeInUnit: o.SeeInUnit,
		SeeInNum:  o.SeeInNum,

		SameClinic:       o.SameClinic,
		SameClinicAnswer: o.SameClinicAnswer,

		PreferredConsultationMethod: o.PreferredConsultationMethod,
		TestsRequired:               o.TestsRequired,
	}
}
