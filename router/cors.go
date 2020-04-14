package router

import (
	"log"
	"os"
	"strconv"

	"github.com/rs/cors"
)

func CorsSetup() *cors.Cors {
	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalln(err)
	}
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://localhost:10888",
		},
		AllowedMethods: []string{
			"HEAD",
			"GET",
			"POST",
			"OPTIONS",
		},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: debugMode,
	})
}
