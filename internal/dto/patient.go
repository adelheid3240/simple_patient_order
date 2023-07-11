package dto

type ListPatientsResp struct {
	Patients []Patient `json:"patients"`
}

type Patient struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}
