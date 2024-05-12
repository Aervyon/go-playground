package main

import (
	"log"
	"net/http"

	myHttp "github.com/Aervyon/go-playground/http"
	"github.com/Aervyon/go-playground/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	// Do migrations
	db.AutoMigrate(&utils.UserModel{})

	r := chi.NewRouter()
	log.Println("Using middlewares: Logger, recoverer, requestID, and heartbeat")
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/health"))

	// user stuffs
	r.Post("/signup", myHttp.CreateUser(db))
	r.Post("/auth", myHttp.AuthUser(db))
	r.Get("/users", myHttp.GetUsers(db))

	log.Println("Listening on port 3457")
	http.ListenAndServe(":3457", r)
}

/*func UsersResponse(users []*utils.UserModel) []render.Renderer {
	list := []render.Renderer{}

	for _, user := range users {
		list = append(list, &utils.User{Username: user.Username})
	}
}*/
