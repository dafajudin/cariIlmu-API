package main

import (
	"cariIlmu-API/database"
	"cariIlmu-API/routes"
)

func main(){

	//initialization database
	database.Init()

	//return endpoint in routes
	e := routes.GetApiRoutes()
	
	//start server in port 8080
	e.Logger.Fatal(e.Start(":8080"))
}