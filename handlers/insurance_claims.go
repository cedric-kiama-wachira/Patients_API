package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"patients_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var insuranceCollection *mongo.Collection

func SetInsuranceCollection(collection *mongo.Collection) {
	insuranceCollection = collection
}

func CreateInsuranceClaim(w http.ResponseWriter, r *http.Request) {
	var claim models.InsuranceClaim
	err := json.NewDecoder(r.Body).Decode(&claim)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	claim.ID = generateID()
	_, err = insuranceCollection.InsertOne(context.TODO(), claim)
	if err != nil {
		http.Error(w, "Failed to create insurance claim", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(claim)
}

func GetInsuranceClaim(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var claim models.InsuranceClaim
	err := insuranceCollection.FindOne(context.TODO(), bson.M{"id": params["id"]}).Decode(&claim)
	if err != nil {
		http.Error(w, "Failed to retrieve insurance claim: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(claim)
}

func UpdateInsuranceClaim(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedClaim models.InsuranceClaim
	err := json.NewDecoder(r.Body).Decode(&updatedClaim)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := bson.M{"id": params["id"]}
	update := bson.M{"$set": updatedClaim}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result models.InsuranceClaim
	err = insuranceCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
	if err != nil {
		http.Error(w, "Failed to update insurance claim", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteInsuranceClaim(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := insuranceCollection.DeleteOne(context.TODO(), bson.M{"id": params["id"]})
	if err != nil {
		http.Error(w, "Failed to delete insurance claim", http.StatusInternalServerError)
		return
	}
}

