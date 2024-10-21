package main

import (
	"aspgo/api"
	WebApplication "aspgo/core"
	"aspgo/cors"
	"aspgo/services"
	"log"
	"net/http"
	"time"
)

func main() {

	builder := WebApplication.CreateBuilder()

	builder.AddController(api.NewHelloController)

	builder.AddService(services.NewUserService)

	app := builder.Build()

	app.UseMiddleware(cors.NewCorsMiddleware(
		&cors.AccessControlAllow{
			Origin:  "*",
			Methods: "GET, POST, PUT, DELETE, OPTIONS",
			Headers: "Content-Type",
		},
	))

	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("Started %s %s", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			duration := time.Since(start).Milliseconds()
			log.Printf("Completed %s %s in %dms", r.Method, r.URL.Path, duration)
		})
	})

	app.Run(":8080")
}
