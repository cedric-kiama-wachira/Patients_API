package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"patients_api/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var medicalRecordsCollection *mongo.Collection
var labResultsCollection *mongo.Collection

func SetMedicalRecordsCollections(medicalRecords *mongo.Collection, labResults *mongo.Collection) {
	medicalRecordsCollection = medicalRecords
	labResultsCollection = labResults
}

// AddMedicalRecord handles POST /patients/{id}/medical-records
func AddMedicalRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	var medicalRecord models.MedicalRecord
	err := json.NewDecoder(r.Body).Decode(&medicalRecord)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	medicalRecord.ID = uuid.New().String()
	medicalRecord.PatientID = patientID

	_, err = medicalRecordsCollection.InsertOne(context.Background(), medicalRecord)
	if err != nil {
		http.Error(w, "Failed to add medical record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medicalRecord)
}

// GetMedicalRecords handles GET /patients/{id}/medical-records
func GetMedicalRecords(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	cursor, err := medicalRecordsCollection.Find(context.Background(), bson.M{"patient_id": patientID})
	if err != nil {
		http.Error(w, "Failed to retrieve medical records: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var medicalRecords []models.MedicalRecord
	for cursor.Next(context.Background()) {
		var medicalRecord models.MedicalRecord
		err := cursor.Decode(&medicalRecord)
		if err != nil {
			http.Error(w, "Failed to decode medical record: "+err.Error(), http.StatusInternalServerError)
			return
		}
		medicalRecords = append(medicalRecords, medicalRecord)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicalRecords)
}

// AddLabResult handles POST /patients/{id}/lab-results
func AddLabResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	var labResult models.LabResult
	err := json.NewDecoder(r.Body).Decode(&labResult)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	labResult.ID = uuid.New().String()
	labResult.PatientID = patientID

	_, err = labResultsCollection.InsertOne(context.Background(), labResult)
	if err != nil {
		http.Error(w, "Failed to add lab result: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(labResult)
}

// GetLabResults handles GET /patients/{id}/lab-results
func GetLabResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	cursor, err := labResultsCollection.Find(context.Background(), bson.M{"patient_id": patientID})
	if err != nil {
		http.Error(w, "Failed to retrieve lab results: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var labResults []models.LabResult
	for cursor.Next(context.Background()) {
		var labResult models.LabResult
		err := cursor.Decode(&labResult)
		if err != nil {
			http.Error(w, "Failed to decode lab result: "+err.Error(), http.StatusInternalServerError)
			return
		}
		labResults = append(labResults, labResult)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(labResults)
}

