package main

import (
	"fmt"
	"net/http"

	"github.com/Kotyarich/find-your-pet/api"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := NewRouter()
	HandlerData := PrepareHandlerData()
	api.ConfigureHandlers(r, HandlerData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
