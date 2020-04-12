package main

import (
	"fmt"
	"net/http"

	"github.com/Kotyarich/find-your-pet/api"
	"github.com/Kotyarich/find-your-pet/router"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := router.NewRouter()
	r.StrictSlash(true)
	HandlerData := PrepareHandlerData()
	c := router.CorsSetup()
	corsHandler := c.Handler(r)
	api.ConfigureHandlers(r, HandlerData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), corsHandler)
}
