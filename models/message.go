package models

type Message struct {
	ID        string `json:"id" bson:"id"`
	PatientID string `json:"patient_id" bson:"patient_id"`
	DoctorID  string `json:"doctor_id" bson:"doctor_id"`
	Content   string `json:"content" bson:"content"`
	Date      string `json:"date" bson:"date"`
}
