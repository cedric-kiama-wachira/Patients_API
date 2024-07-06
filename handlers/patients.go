package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "patients_api/models"
)

// CreatePatient creates a new patient record
func CreatePatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    err := json.NewDecoder(r.Body).Decode(&patient)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `INSERT INTO patients (id, birthdate, first, last, gender) VALUES ($1, $2, $3, $4, $5)`
    _, err = pgdb.Exec(query, patient.ID, patient.Birthdate, patient.First, patient.Last, patient.Gender)
    if err != nil {
        http.Error(w, "Failed to create patient: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(patient)
}

// GetPatient retrieves a patient by ID
func GetPatient(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var patient models.Patient
    query := `SELECT id, birthdate, first, last, gender FROM patients WHERE id = $1`
    err := pgdb.QueryRow(query, id).Scan(&patient.ID, &patient.Birthdate, &patient.First, &patient.Last, &patient.Gender)
    if err == sql.ErrNoRows {
        http.Error(w, "Patient not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Failed to retrieve patient: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(patient)
}

// UpdatePatient updates an existing patient record
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var patient models.Patient
    err := json.NewDecoder(r.Body).Decode(&patient)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `UPDATE patients SET birthdate = $1, first = $2, last = $3, gender = $4 WHERE id = $5`
    _, err = pgdb.Exec(query, patient.Birthdate, patient.First, patient.Last, patient.Gender, params["id"])
    if err != nil {
        http.Error(w, "Failed to update patient: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(patient)
}

// DeletePatient deletes a patient record
func DeletePatient(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    query := `DELETE FROM patients WHERE id = $1`
    _, err := pgdb.Exec(query, params["id"])
    if err != nil {
        http.Error(w, "Failed to delete patient: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
