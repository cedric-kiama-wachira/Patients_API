package main

import (
	"log"
	"net/http"
	"patients_api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	pgdb, _, err := handlers.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize databases: %v", err)
	}
	defer pgdb.Close()

	r := mux.NewRouter()

	// Patient Management System
	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients/{id}", handlers.GetPatient).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")

	// Appointment Scheduling System
	r.HandleFunc("/appointments", handlers.CreateAppointment).Methods("POST")
	r.HandleFunc("/appointments/{id}", handlers.GetAppointment).Methods("GET")
	r.HandleFunc("/appointments/{id}", handlers.UpdateAppointment).Methods("PUT")
	r.HandleFunc("/appointments/{id}", handlers.DeleteAppointment).Methods("DELETE")
	r.HandleFunc("/patients/{id}/appointments", handlers.GetPatientAppointments).Methods("GET")

	// Medical Records System
	r.HandleFunc("/patients/{id}/medical-records", handlers.AddMedicalRecord).Methods("POST")
	r.HandleFunc("/patients/{id}/medical-records", handlers.GetMedicalRecords).Methods("GET")
	r.HandleFunc("/patients/{id}/lab-results", handlers.AddLabResult).Methods("POST")
	r.HandleFunc("/patients/{id}/lab-results", handlers.GetLabResults).Methods("GET")

	// Doctor-Patient Communication Platform
	r.HandleFunc("/messages", handlers.SendMessage).Methods("POST")
	r.HandleFunc("/messages/{id}", handlers.GetMessage).Methods("GET")
	r.HandleFunc("/patients/{id}/messages", handlers.GetPatientMessages).Methods("GET")

	// Billing and Insurance Management System
	r.HandleFunc("/billing", handlers.CreateBillingRecord).Methods("POST")
	r.HandleFunc("/billing/{id}", handlers.GetBillingRecord).Methods("GET")
	r.HandleFunc("/billing/{id}", handlers.UpdateBillingRecord).Methods("PUT")
	r.HandleFunc("/billing/{id}", handlers.DeleteBillingRecord).Methods("DELETE")

	r.HandleFunc("/insurance-claims", handlers.CreateInsuranceClaim).Methods("POST")
	r.HandleFunc("/insurance-claims/{id}", handlers.GetInsuranceClaim).Methods("GET")
	r.HandleFunc("/insurance-claims/{id}", handlers.UpdateInsuranceClaim).Methods("PUT")
	r.HandleFunc("/insurance-claims/{id}", handlers.DeleteInsuranceClaim).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

