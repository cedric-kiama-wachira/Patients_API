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

var messagesCollection *mongo.Collection

func SetMessagesCollection(collection *mongo.Collection) {
	messagesCollection = collection
}

// SendMessage handles POST /messages
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	message.ID = uuid.New().String()

	_, err = messagesCollection.InsertOne(context.Background(), message)
	if err != nil {
		http.Error(w, "Failed to send message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

// GetMessage handles GET /messages/{id}
func GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageID := params["id"]

	var message models.Message
	err := messagesCollection.FindOne(context.Background(), bson.M{"id": messageID}).Decode(&message)
	if err != nil {
		http.Error(w, "Failed to retrieve message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// GetPatientMessages handles GET /patients/{id}/messages
func GetPatientMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	patientID := params["id"]

	cursor, err := messagesCollection.Find(context.Background(), bson.M{"patient_id": patientID})
	if err != nil {
		http.Error(w, "Failed to retrieve messages: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var messages []models.Message
	for cursor.Next(context.Background()) {
		var message models.Message
		err := cursor.Decode(&message)
		if err != nil {
			http.Error(w, "Failed to decode message: "+err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

