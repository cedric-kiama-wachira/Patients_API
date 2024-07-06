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

var billingCollection *mongo.Collection

func SetBillingCollection(collection *mongo.Collection) {
	billingCollection = collection
}

func CreateBillingRecord(w http.ResponseWriter, r *http.Request) {
	var billing models.Billing
	err := json.NewDecoder(r.Body).Decode(&billing)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	billing.ID = generateID()
	_, err = billingCollection.InsertOne(context.TODO(), billing)
	if err != nil {
		http.Error(w, "Failed to create billing record", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(billing)
}

func GetBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var billing models.Billing
	err := billingCollection.FindOne(context.TODO(), bson.M{"id": params["id"]}).Decode(&billing)
	if err != nil {
		http.Error(w, "Failed to retrieve billing record: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(billing)
}

func UpdateBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedBilling models.Billing
	err := json.NewDecoder(r.Body).Decode(&updatedBilling)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := bson.M{"id": params["id"]}
	update := bson.M{"$set": updatedBilling}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result models.Billing
	err = billingCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
	if err != nil {
		http.Error(w, "Failed to update billing record", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := billingCollection.DeleteOne(context.TODO(), bson.M{"id": params["id"]})
	if err != nil {
		http.Error(w, "Failed to delete billing record", http.StatusInternalServerError)
		return
	}
}

