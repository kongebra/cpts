package util

import "os"

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return ":" + port
}