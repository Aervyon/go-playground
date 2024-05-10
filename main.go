package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aervyon/go-playground/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
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

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []utils.UserModel{}
		transaction := db.Model(&utils.UserModel{}).Find(&users, "")
		fmt.Printf("%v\n", transaction.RowsAffected)

		if transaction.Error != nil {
			w.WriteHeader(500)
			w.Write([]byte(transaction.Error.Error()))
		}

		if transaction.RowsAffected == 0 {
			log.Println("Got 0 rows for users")
		}

		fmt.Println(users)
		render.JSON(w, r, users)
	})

	log.Println("Listening on port 3457")
	http.ListenAndServe(":3457", r)
}

/*func UsersResponse(users []*utils.UserModel) []render.Renderer {
	list := []render.Renderer{}

	for _, user := range users {
		list = append(list, &utils.User{Username: user.Username})
	}
}*/
