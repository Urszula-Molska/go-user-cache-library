# go-user-cache-library

A program in Go language creating a cache for users with a REST API, using the go-cache and mongo-driver libraries. It listens on port 8080 and provide one endpoint: /user/{id}. If user data is not found in the cache, the program retrieves it from the MongoDB database, stores it in the cache, and returns it to the client.

# Remarks!!!

godotenv is not a standard Go package and needs to be installed separately.
Install the godotenv package and update go.mod file using the following commands:

**go get github.com/joho/godotenv**

**go mod tidy**

create .env file with URI link to Mongo Database (MongoDB Atlas connection string)
https://www.mongodb.com/docs/manual/reference/connection-string/?utm_source=compass&utm_medium=product

URI_MONGO= addressToTheMongoDatabase

then use the command to finally run the program

**_go run ._**
