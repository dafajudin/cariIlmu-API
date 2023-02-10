package main

import (
	"cariIlmu-API/src/database"
	"cariIlmu-API/src/routes"
)

func main(){

	//init database
	database.Init()

	e := routes.GetApiRoutes()
	
	e.Logger.Fatal(e.Start(":8080"))
}