package main

import (
	"idstar.com/app/routers"
	"idstar.com/docs"

	config "idstar.com/app/configs"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
