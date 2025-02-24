// filepath: /Users/rohit/Documents/go-chassis-flight-management-system/flight-service/main.go
package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	// Create Iris app instance
	InstallPlugin()
	app := iris.New()

	// Set up middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(iris.Compression)
	app.Use(crs)

	// Define routes
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello from flight-service"})
	})

	// Create Iris server wrapper

	// Register the custom Iris server with Go-Chassis
	chassis.RegisterSchema("rest", app)

	// Initialize Go-Chassis
	if err := chassis.Init(); err != nil {
		panic("Go-Chassis initialization failed: " + err.Error())
	}

	// Start Go-Chassis
	chassis.Run()
}
