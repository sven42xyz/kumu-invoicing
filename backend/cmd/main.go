package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"chapter42.de/m/internal/database"
    "chapter42.de/m/internal/handlers"
	"chapter42.de/m/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Session store to manage access of logged in users
var store = sessions.NewCookieStore([]byte("placeholder"))

func main() {
    // Initialize database connection
    db, err := database.NewMySQLDB("application_auth:auth@tcp(172.17.224.1:3306)/kumu")
    if err != nil {
        log.Fatal("Error initializing database:", err)
    }
    defer db.Close()

    // Router
    r := mux.NewRouter()

    // Handlers -> eventually export as "package routes"
    r.HandleFunc("/", handlers.RootHandler()).Methods("GET")
    r.HandleFunc("/auth/register", handlers.RegisterHandler(db)).Methods("POST")
    r.HandleFunc("/auth/login", handlers.LoginHandler(db, store)).Methods("POST")
    r.HandleFunc("/auth/logout", handlers.LogoutHandler(db, store)).Methods("POST")
    r.HandleFunc("/auth/user", handlers.UserHandler(db, store)).Methods("PUT")
    r.HandleFunc("/auth/user", handlers.UserHandler(db, store)).Methods("DELETE")

    // Rate limiter middleware
    limiter := middleware.NewRateLimiter(10, 1) // Allow 10 requests per second with a burst of 1 request
    r.Use(limiter.Middleware)

    // Server
    srv := &http.Server{
        Addr:    ":3030",
        Handler: r,
    }

    // Start server in a separate goroutine
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe: %v", err)
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    log.Println("Server shutting down...")

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Shutdown server
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server gracefully stopped")
}
