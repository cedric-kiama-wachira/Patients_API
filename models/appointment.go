package models

type Appointment struct {
	ID          string `json:"id" bson:"id"`
	PatientID   string `json:"patient_id" bson:"PATIENT"`
	StartTime   string `json:"start_time" bson:"START"`
	EndTime     string `json:"end_time" bson:"STOP"`
	Description string `json:"description" bson:"DESCRIPTION"`
}

