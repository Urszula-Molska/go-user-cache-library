package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
    ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() { 

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables from .env file. Please ensure that the file exists and is correctly configured.")
	}

    var userCache = cache.New(1*time.Minute, 5*time.Minute)

    http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {

        userID := r.URL.Path[len("/user/"):]

        //Check if the user data is in the cache
	    if userData, found := userCache.Get(userID); found {
		fmt.Println("Retrieving user data from cache memory:\n", userData)
        fmt.Fprintf(w, "userDataFromCache: ")
		json.NewEncoder(w).Encode(userData)
		return
        }
        //if there is no user data in cache then get them from the database and save in catch memory
        fmt.Fprintf(w, "userDataFromMongo: ")
         user, err:= getUserFromDatabaseByID(userID)
         if err != nil {
            http.Error(w, " User not found in Database", http.StatusNotFound)
            return
        };
        json.NewEncoder(w).Encode(user)
        fmt.Println("Retrieving user data from DATABASE MONGO DB:\n", user)
        //saving userdata in cache memory
        userCache.Set(userID, user, cache.DefaultExpiration)
    })

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
    log.Fatalf("Failed to start the server on port 8080. Please check if the port is available and not in use by another application. Error: %v", err)
	};
}

	func getUserFromDatabaseByID (userID string) (*User, error) {

	URI:=os.Getenv("URI_MONGO")

    mongoUri := fmt.Sprintf("mongodb+srv://%s",URI)

    // CONNECT TO YOUR ATLAS CLUSTER:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}
	fmt.Println("Connected to MongoDB!");

    collection := client.Database("Users").Collection("users")

	// Find User by ID in Mongo Database.
	var user User
	err = collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user in the database: %v", err)
	}

	return &user, nil}




