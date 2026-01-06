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

// Define your data structure
type Mission struct {
	Name        string    `bson:"name"`
	Destination string    `bson:"destination"`
	LaunchYear  int       `bson:"launch_year"`
	Successful  bool      `bson:"successful"`
	CreatedAt   time.Time `bson:"created_at"`
}

func main() {
	// 1. Setup your connection string
	// Replace <password> with your actual password and <cluster-url> with your Atlas link
	const uri = "mongodb+srv://<username>:<password>@<cluster-url>/?retryWrites=true&w=majority"

	// 2. Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. Connect to MongoDB Atlas
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection Error:", err)
	}
	defer client.Disconnect(ctx)

	// 4. Select your Database and Collection
	collection := client.Database("SpaceDB").Collection("Missions")

	// 5. Create a piece of data
	newMission := Mission{
		Name:        "Artemis II",
		Destination: "Moon",
		LaunchYear:  2025,
		Successful:  true,
		CreatedAt:   time.Now(),
	}

	// 6. Insert into the Database
	result, err := collection.InsertOne(ctx, newMission)
	if err != nil {
		log.Fatal("Insert Error:", err)
	}

	fmt.Printf("Successfully inserted mission with ID: %v\n", result.InsertedID)
}
