package main

import (
	"idstar.com/app/routers"
	"idstar.com/docs"

	config "idstar.com/app/configs"
)

func main() {
	// Initialize connection to Database
	config.InitDB()

	r := routers.SetupRouter()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger User API"
	docs.SwaggerInfo.Description = "This is a sample Swagger in Golang with GIN Framework."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5001/api"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.Run(":5001")
}
