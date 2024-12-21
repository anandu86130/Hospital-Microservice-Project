package di

import (
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	"github.com/anandu86130/Hospital-api-gateway/internal/admin"
	"github.com/anandu86130/Hospital-api-gateway/internal/chat"
	"github.com/anandu86130/Hospital-api-gateway/internal/doctor"
	"github.com/anandu86130/Hospital-api-gateway/internal/server"
	"github.com/anandu86130/Hospital-api-gateway/internal/user"
)

func Init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration file: %v", err) // Provide detailed error output
	}
	if cfg.SECRETKEY == "" {
		log.Fatalf("JWT secret key (SECRETKEY) is not configured.")
	}

	// Initialize server and routes
	svr := server.NewServer()
	svr.R.LoadHTMLGlob("template/*")
	if err := user.NewUserRoute(svr.R, *cfg); err != nil {
		log.Fatalf("Failed to initialize user routes: %v", err)
	}
	admin.NewAdminRoute(svr.R, *cfg)
	doctor.NewDoctorRoutes(svr.R, *cfg)
	chat.NewChatRoutes(svr.R, *cfg)

	// Start the server
	svr.StartServer(cfg.APIPORT)
}