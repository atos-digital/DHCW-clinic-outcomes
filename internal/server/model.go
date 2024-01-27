package server

import "github.com/atos-digital/DHCW-clinic-outcomes/ui/pages"

type OutcomesForm struct {
	ConsultationDate string `json:"consultation_date"`
	ConsultationTime string `json:"consultation_time"`
	ConsultationType string `json:"consultation_type"`
	Specialties      string `json:"specialties"`
	Clinicians       string `json:"clinicians"`

	CancerPathway       string `json:"cancer_pathway"`
	CancerPathwayOption string `json:"cancer_pathway_options"`

	OutcomesOption     string `json:"outcomes_option"`
	SeeInNum           string `json:"see_in_num"`
	SeeInUnit          string `json:"see_in_unit"`
	SeeOnSymptomMonths string `json:"see_on_symptom_months"`
}

func (o OutcomesForm) State() pages.OutcomesState {
	return pages.OutcomesState{
		ConsultationDate: o.ConsultationDate,
		ConsultationTime: o.ConsultationTime,
		ConsultationType: o.ConsultationType,
		Specialties:      o.Specialties,
		Clinicians:       o.Clinicians,
		CancerPathway: pages.CancerPathway{
			Checked: o.CancerPathway == "on",
			Option:  o.CancerPathwayOption,
		},
	}
}
