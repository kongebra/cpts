package util

import "os"

/*
GetPort checks if the system environment port is set, if not set's port to 3000
 */
func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return ":" + port
}

/*
GetMongoURL checks if the system environment url is set, if not set's url to localhost:27017
 */
func GetMongoURL() string {
	var url = os.Getenv("MONGO_URL")

	if url == "" {
		url = "localhost:27017"
	}

	return url
}