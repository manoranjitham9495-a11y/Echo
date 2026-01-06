package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define the structure for your data
type FavoriteThing struct {
	Name     string `bson:"name"`
	Category string `bson:"category"` // Film or Dinosaur
	Detail   string `bson:"detail"`
}

func main() {
	// Replace <connection_string> with your actual MongoDB Atlas URI
	// Make sure to replace <password> with your database user password!
	const uri = "mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority"

	// Set up connection options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection Error:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ping Error:", err)
	}

	fmt.Println("Connected to MongoDB Atlas successfully!")

	// Access the database and collection
	collection := client.Database("LearningDB").Collection("Contributors")

	// Create data (As per your task: Dinosaur or Film)
	myData := FavoriteThing{
		Name:     "T-Rex",
		Category: "Dinosaur",
		Detail:   "The King of the Lizards - definitely not a Pterodactyl!",
	}

	// Insert Data
	result, err := collection.InsertOne(ctx, myData)
	if err != nil {
		log.Fatal("Insertion Error:", err)
	}

	fmt.Printf("Data inserted with ID: %v\n", result.InsertedID)
}
