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