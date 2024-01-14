# go-user-cache-library
A program in Go language creating a cache for users with a REST API, using the go-cache and mongo-driver libraries. It listens on port 8080 and provide one endpoint: /user/{id}. If user data is not found in the cache, the program retrieves it from the MongoDB database, stores it in the cache, and returns it to the client.
