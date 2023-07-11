package dto

type CreatePatientOrder struct {
	Message string `json:"message"`
}

type ListPatientOrderResp struct {
	Orders []PatientOrder `json:"orders,omitempty"`
}

type PatientOrder struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
