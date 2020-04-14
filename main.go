package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Kotyarich/find-your-pet/api"
	"github.com/Kotyarich/find-your-pet/router"
	"github.com/spf13/viper"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(os.Getenv("DEBUG")) == "true" {
			log.Printf("%s %s\n Origin: %s\n", r.Method,
				r.URL, r.Header.Get("Origin"))
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	PrepareConfig()
	r := router.NewRouter()
	r.StrictSlash(true)
	HandlerData := PrepareHandlerData()
	c := router.CorsSetup()
	corsHandler := c.Handler(r)
	api.ConfigureHandlers(r, HandlerData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), logRequest(corsHandler))
}
