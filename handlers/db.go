package handlers

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var pgdb *sql.DB
var mongodb *mongo.Database

func InitDB() (*sql.DB, *mongo.Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// PostgreSQL connection
	pgConnStr := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	pgdb, err = sql.Open("postgres", pgConnStr)
	if err != nil {
		return nil, nil, err
	}

	err = pgdb.Ping()
	if err != nil {
		return nil, nil, err
	}

	// MongoDB connection
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}

	mongodb = client.Database(os.Getenv("MONGO_DB"))

	// Initialize collections
	SetAppointmentsCollection(mongodb.Collection("encounters"))
	SetMedicalRecordsCollections(mongodb.Collection("medical_records"), mongodb.Collection("lab_results"))
	SetMessagesCollection(mongodb.Collection("messages"))
	SetBillingCollection(mongodb.Collection("billing"))
	SetInsuranceCollection(mongodb.Collection("insurance_claims"))

	return pgdb, mongodb, nil
}

