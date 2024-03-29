package models

type ClinicOutcomesFormState struct {
	Details          DetailsState
	CancerPathway    CancerPathwayState
	Outcome          OutcomeState
	FollowUp         FollowUpState
	OtherInformation string
	Errors           []string
}

type DetailsState struct {
	EventDate      string
	EventTime      string
	EventType      string
	EventSpecialty string
	EventClinician string
}

type CancerPathwayState struct {
	Checked bool
	Option  string
	Other   string
}

type OutcomeState struct {
	OutcomeOption                string
	SeeOnSymptomDetails          string
	DidNotAnswerDetails          string
	DidNotAttendDetails          string
	CouldNotAttendDetails        string
	ReferToDiagnosticsDetails    string
	ReferToAnotherDetails        string
	ReferToTherapiesDetails      string
	ReferToTreatment             string
	ReferToTreatmentSact         bool
	ReferToTreatmentRadiotherapy bool
	ReferToTreatmentOther        bool
	ReferToTreatmentDetails      string
	DiscussAtMdtDetails          string
	OutpatientProcedureDetails   string
	FollowUpChecked              bool
}

type Test struct {
	TestsRequired     string
	TestsUndertakenBy string
	TestsRequiredBy   string
}

type FollowUpState struct {
	Checked                     bool
	Pathway                     string
	SameClinician               string
	SameClinicianNo             string
	SameClinic                  string
	SameClinicNo                string
	SeeInNum                    string
	SeeInUnit                   string
	Hospital                    string
	AppointmentPriority         string
	Condition                   string
	PreferredConsultationMethod string
	TestsRequiredBeforeFollowup string
	Tests                       []Test
}

func State(payload ClinicOutcomesFormPayload) ClinicOutcomesFormState {
	num := len(payload.TestsRequired)
	if num == 0 {
		num = 1
	}
	tests := make([]Test, num)
	for i, v := range payload.TestsRequired {
		tests[i] = Test{
			TestsRequired:     v,
			TestsUndertakenBy: payload.TestsUndertakenBy[i],
			TestsRequiredBy:   payload.TestsBy[i],
		}
	}

	FollowUpChecked := payload.FollowUp == "on" && payload.OutcomeOption != "Patient Initiated Follow Up"

	return ClinicOutcomesFormState{
		Details: DetailsState{
			EventDate:      payload.EventDate,
			EventTime:      payload.EventTime,
			EventType:      payload.EventType,
			EventSpecialty: payload.EventSpecialty,
			EventClinician: payload.EventClinician,
		},
		CancerPathway: CancerPathwayState{
			Checked: payload.CancerPathway == "on",
			Option:  payload.CancerPathwayOption,
			Other:   payload.CancerPathwayOther,
		},
		Outcome: OutcomeState{
			OutcomeOption:       payload.OutcomeOption,
			SeeOnSymptomDetails: payload.SeeOnSymptomDetails,

			DidNotAnswerDetails:   payload.DidNotAnswerDetails,
			DidNotAttendDetails:   payload.DidNotAttendDetails,
			CouldNotAttendDetails: payload.CouldNotAttendDetails,

			ReferToDiagnosticsDetails: payload.ReferToDiagnosticsDetails,
			ReferToAnotherDetails:     payload.ReferToAnotherDetails,
			ReferToTherapiesDetails:   payload.ReferToTherapiesDetails,

			ReferToTreatmentSact:         payload.ReferToTreatmentSact == "on",
			ReferToTreatmentRadiotherapy: payload.ReferToTreatmentRadiotherapy == "on",
			ReferToTreatmentOther:        payload.ReferToTreatmentOther == "on",
			ReferToTreatmentDetails:      payload.ReferToTreatmentDetails,

			DiscussAtMdtDetails:        payload.DiscussAtMdtDetails,
			OutpatientProcedureDetails: payload.OutpatientProcedureDetails,
			FollowUpChecked:            FollowUpChecked,
		},
		FollowUp: FollowUpState{
			Checked:                     FollowUpChecked,
			Pathway:                     payload.Pathway,
			SameClinician:               payload.SameClinician,
			SameClinicianNo:             payload.SameClinicianNo,
			SameClinic:                  payload.SameClinic,
			SameClinicNo:                payload.SameClinicNo,
			SeeInUnit:                   payload.SeeInUnit,
			SeeInNum:                    payload.SeeInNum,
			Hospital:                    payload.Hospital,
			AppointmentPriority:         payload.AppointmentPriority,
			Condition:                   payload.Condition,
			PreferredConsultationMethod: payload.PreferredConsultationMethod,
			TestsRequiredBeforeFollowup: payload.TestsRequiredBeforeFollowup,
			Tests:                       tests,
		},
		OtherInformation: payload.OtherInformation,
		Errors:           payload.Errors,
	}
}
