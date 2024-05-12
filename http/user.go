package http

import (
	"net/http"

	"log"

	"github.com/Aervyon/go-playground/utils"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			return
		}
		if !r.Form.Has("username") || !r.Form.Has("password") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ code: 400, message: \"Signup requires username and password\" }"))
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")
		existingUser := &utils.UserModel{}
		db.Model(&utils.UserModel{}).Find(&existingUser, "username = ?", username)
		if existingUser.ID != "" {
			w.WriteHeader(http.StatusIMUsed)
			w.Write([]byte("{ code: 226, message: \"Username or email taken\" }"))
			return
		}

		user, err := utils.CreateUser(db, username, password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{ code: 500, message: \"Failed to make user\" }"))
			log.Println("Error saving user", username, "to db:", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{code: 200, message: \"Saved your info. Welcome " + user.Username + "!\"}"))
	}
}
