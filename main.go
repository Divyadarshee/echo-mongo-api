package main

import ( // importing required dependencies
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	e := echo.New() // initialize an Echo application with the New() function

	// run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//e.GET("/", func(c echo.Context) error { // GET function to the route = "/" path and an handler
	//	return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"}) // function that returns a JSON of "Hello from Echo & mongoDB".
	//	// echo.Map is a shortcut for map[string]interface{} useful for JSON returns
	//})

	e.Logger.Fatal(e.Start(":6000")) // Start function is used to run the application on port 6000

}
