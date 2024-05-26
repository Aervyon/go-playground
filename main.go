package main

import (
	"log"
	"net/http"

	"github.com/Aervyon/go-playground/endpoints"
	"github.com/Aervyon/go-playground/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/alexedwards/scs/gormstore"
	"github.com/alexedwards/scs/v2"
)

func main() {
	databaseType := "sqlite-mem"

	var dbDriver gorm.Dialector

	switch databaseType {
	case "sqlite":
		log.Println("using db driver: sqlite")
		dbDriver = sqlite.Open("test.db")
	case "sqlite-mem":
		log.Println("Using db driver: sqlite memory")
		dbDriver = sqlite.Open(":memory:")
	case "postgres":
		panic("postgres is unsupported")
	default:
		log.Println("Using default db driver: sqlite")
		dbDriver = sqlite.Open("test.db")
	}

	db, err := gorm.Open(dbDriver, &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	log.Println("Connected to database")

	sessionManager := scs.New()
	if sessionManager.Store, err = gormstore.New(db); err != nil {
		log.Fatal(err)
	}

	// Do migrations
	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.Token{})

	r := chi.NewRouter()
	log.Println("Using middlewares: Logger, recoverer, requestID, CORS, and heartbeat")
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"User-Agent",
			"Host",
			"Referer",
			"Origin",
			"Cache-Control",
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		AllowCredentials: true,
	}))

	// user stuffs
	r.Post("/api/signup", endpoints.CreateUser(db))
	r.Post("/api/auth", endpoints.AuthUser(db, sessionManager))
	r.Get("/api/account", endpoints.GetSelfAccount(db, sessionManager))
	r.Get("/api/users", endpoints.GetUsers(db))

	log.Println("Listening on port 3457")
	http.ListenAndServe(":3457", sessionManager.LoadAndSave(r))
}
