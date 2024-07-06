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

var appointmentsCollection *mongo.Collection

// SetAppointmentsCollection sets the MongoDB collection for appointments
func SetAppointmentsCollection(collection *mongo.Collection) {
	appointmentsCollection = collection
}

// CreateAppointment handles POST /appointments
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	appointment.ID = uuid.New().String()

	_, err = appointmentsCollection.InsertOne(context.Background(), appointment)
	if err != nil {
		http.Error(w, "Failed to create appointment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

// GetAppointment handles GET /appointments/{id}
func GetAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appointmentID := params["id"]

	var appointment models.Appointment
	err := appointmentsCollection.FindOne(context.Background(), bson.M{"id": appointmentID}).Decode(&appointment)
	if err != nil {
		http.Error(w, "Appointment not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

// UpdateAppointment handles PUT /appointments/{id}
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appointmentID := params["id"]

	var updatedAppointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&updatedAppointment)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	updatedAppointment.ID = appointmentID

	_, err = appointmentsCollection.ReplaceOne(context.Background(), bson.M{"id": appointmentID}, updatedAppointment)
	if err != nil {
		http.Error(w, "Failed to update appointment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedAppointment)
}

// DeleteAppointment handles DELETE /appointments/{id}
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appointmentID := params["id"]

	_, err := appointmentsCollection.DeleteOne(context.Background(), bson.M{"id": appointmentID})
	if err != nil {
		http.Error(w, "Failed to delete appointment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetPatientAppointments handles GET /patients/{id}/appointments
func GetPatientAppointments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	cursor, err := appointmentsCollection.Find(context.Background(), bson.M{"patient_id": patientID})
	if err != nil {
		http.Error(w, "Failed to retrieve appointments: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var appointments []models.Appointment
	for cursor.Next(context.Background()) {
		var appointment models.Appointment
		err := cursor.Decode(&appointment)
		if err != nil {
			http.Error(w, "Failed to decode appointment: "+err.Error(), http.StatusInternalServerError)
			return
		}
		appointments = append(appointments, appointment)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

