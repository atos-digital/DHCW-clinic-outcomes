package models

type OutcomeDetails struct {
	ConsultationDate string
	ConsultationTime string
	ConsultationType string
	Specialties      string
	Clinicians       string
}
type CancerPathway struct {
	Checked bool
	Option  string
	Other   string
}

type OutcomeOptions struct {
	PatientOption                string
	SeeOnSymptomMonths           string
	DidNotAnswer                 string
	DidNotAttend                 string
	CouldNotAttend               string
	ReferToDiagnosticsData       string
	ReferToAnotherData           string
	ReferToTherapyData           string
	DiscussAtMdtData             string
	ReferToTreatment             string
	ReferRoTreatmentSact         string
	ReferRoTreatmentRadiotherapy string
	ReferRoTreatmentOther        string
	ReferRoTreatmentData         string
	OutpatientProcedureData      string
}

type Test struct {
	TestsRequired   string
	UndertakenBy    string
	TestsRequiredBy string
}

type FollowUp struct {
	FollowUp                    string
	Pathway                     string
	SameClinician               string
	SameClinicianAnswer         string
	SameClinic                  string
	SameClinicAnswer            string
	SeeInUnit                   string
	SeeInNum                    string
	Hospital                    string
	AppointmentDP               string
	Condition                   string
	PreferredConsultationMethod string
	TestsRequired               string
	Tests                       []Test
}

type OutcomesState struct {
	OutcomeDetails   OutcomeDetails
	CancerPathway    CancerPathway
	OutcomeOptions   OutcomeOptions
	FollowUp         FollowUp
	OtherInformation string
}
