package main

import (
	"Democratic_shopping_mall/config"
	"Democratic_shopping_mall/router"

	
)


func main() {
	config.InitConfig()
	r := router.SetUpRouter()

	port := config.Appconfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}

