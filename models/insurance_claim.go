package models

type InsuranceClaim struct {
	ID          string  `json:"id" bson:"id"`
	PatientID   string  `json:"patient_id" bson:"patient_id"`
	Amount      float64 `json:"amount" bson:"amount"`
	Date        string  `json:"date" bson:"date"`
	Status      string  `json:"status" bson:"status"`
	Description string  `json:"description" bson:"description"`
}

