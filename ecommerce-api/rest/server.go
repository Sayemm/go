package rest

import (
	"context"
	"ecommerce/config"
	"ecommerce/rest/handlers/cart"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct { //Dependency
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	cartHandler    *cart.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	cartHandler *cart.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler, //Inject
		userHandler:    userHandler,
		cartHandler:    cartHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	// Health check endpoint (to verify server is running)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy","service":"%s","version":"%s"}`,
			server.cnf.ServiceName, server.cnf.Version)
	})

	// Root endpoint (welcome message)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"Welcome to %s API","version":"%s"}`,
			server.cnf.ServiceName, server.cnf.Version)
	})

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.cartHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	address := fmt.Sprintf(":%d", server.cnf.HttpPort)

	httpServer := &http.Server{
		Addr:         address,
		Handler:      wrappedMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		fmt.Printf("Server running on port %d\n", server.cnf.HttpPort)
		fmt.Println("\nPress Ctrl+C to stop")

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// Give active requests 30 seconds to finish
	// This is "graceful" shutdown - don't drop ongoing requests
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("✓ Server stopped gracefully")
}
