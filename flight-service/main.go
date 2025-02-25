// filepath: /Users/rohit/Documents/go-chassis-flight-management-system/flight-service/main.go
package main

import (
	"flight-service/config"
	_ "flight-service/docs"
	"flight-service/routes"

	"github.com/go-chassis/go-chassis/v2"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize logger" + err.Error())
	}
	defer logger.Sync()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server FlightManager server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("http://localhost:8080/swagger/swagger.json"),
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)

	// Create Iris app instance
	InstallPlugin()
	app := iris.New()

	// Set up middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	app.Use(iris.Compression)
	app.Use(crs)
	db := config.ConnectDb()
	// checkining
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello from flight-service"})
	})

	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)

	routes.FlightRoutes(app, logger, db)

	// Register the custom Iris server with Go-Chassis
	chassis.RegisterSchema("rest", app)

	// Initialize Go-Chassis
	if err := chassis.Init(); err != nil {
		panic("Go-Chassis initialization failed: " + err.Error())
	}

	// Start Go-Chassis
	chassis.Run()
}
