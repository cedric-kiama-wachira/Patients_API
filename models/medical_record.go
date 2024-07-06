package models

type MedicalRecord struct {
	ID          string `json:"id" bson:"id"`
	PatientID   string `json:"patient_id" bson:"patient_id"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

type LabResult struct {
	ID          string `json:"id" bson:"id"`
	PatientID   string `json:"patient_id" bson:"patient_id"`
	Result      string `json:"result" bson:"result"`
	Date        string `json:"date" bson:"date"`
}
