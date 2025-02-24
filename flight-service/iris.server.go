package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/kataras/iris/v12"
)

func init() {
	InstallPlugin()
}

func InstallPlugin() {
	server.InstallPlugin("rest", newIrisServer)
}

type irisServer struct {
	app    *iris.Application
	opts   server.Options
	server *http.Server
}

func newIrisServer(opts server.Options) server.ProtocolServer {
	return &irisServer{
		opts:   opts,
		server: nil,
	}
}

func (i *irisServer) Register(schema interface{}, options ...server.RegisterOption) (string, error) {
	app, ok := schema.(*iris.Application) // Ensure correct type
	if !ok {
		return "", fmt.Errorf("failed to register schema: invalid type")
	}
	i.app = app
	fmt.Println("Iris schema registered successfully")
	return "irisServer", nil
}

// Start runs the Iris server in a goroutine
func (i *irisServer) Start() error {
	if i.app == nil {
		return fmt.Errorf("iris application is not registered")
	}
	go func() {
		err := i.app.Listen(":8080")
		if err != nil {
			fmt.Println("Error starting Iris server:", err)
		}
	}()
	fmt.Println("Iris Server started on port 8080")
	return nil
}

// Stop method to implement graceful shutdown
func (i *irisServer) Stop() error {
	if i.app == nil {
		return fmt.Errorf("iris application is not initialized")
	}
	fmt.Println("Stopping Iris Server...")
	return i.app.Shutdown(context.Background())
}

// String method to return the server name
func (i *irisServer) String() string {
	return "iris"
}
