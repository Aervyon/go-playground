package http

import (
	"net/http"

	"log"

	"github.com/Aervyon/go-playground/models"
	"github.com/Aervyon/go-playground/utils"
	"github.com/alexedwards/argon2id"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func AuthUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, render.M{"code": http.StatusBadRequest, "message": "Failed to parse form"})
			return
		}

		if !r.Form.Has("username") || !r.Form.Has("password") {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, render.M{"code": 400, "message": "Authentication requires username & password"})
			return
		}

		username := r.Form.Get("username")
		var user utils.UserModel
		db.Model(&utils.UserModel{}).First(&user, "username = ?", username)
		if user.ID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, render.M{"code": http.StatusUnauthorized, "message": "Authentication Failed"})
			return
		}

		match, _, err := argon2id.CheckHash(r.Form.Get("password"), user.Password)
		if err != nil {
			log.Println("Error checking user's", username, "hash:", err)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{"code": http.StatusInternalServerError, "message": "Failed checking password", "error": err.Error()})
			return
		}

		if !match {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, render.M{"code": http.StatusUnauthorized, "message": "Authentication Failed"})
			return
		}

		token := models.NewToken(user.ID)

		transaction := db.Model(&models.Token{}).Create(token)
		if transaction.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{"code": http.StatusInternalServerError, "message": "access token generation failed"})
		}

		render.JSON(w, r, render.M{"code": 201, "message": "Created Token", "token": token.Token, "tokenType": "Bearer"})
	}
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, render.M{"code": http.StatusBadRequest, "message": "Failed to parse form"})
			return
		}
		if !r.Form.Has("username") || !r.Form.Has("password") {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, render.M{"code": 400, "message": "Signup requires username & password"})
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")
		existingUser := &utils.UserModel{}
		db.Model(&utils.UserModel{}).Find(&existingUser, "username = ?", username)
		if existingUser.ID != "" {
			w.WriteHeader(http.StatusIMUsed)
			render.JSON(w, r, render.M{"code": http.StatusIMUsed, "message": "Username or email taken"})
			return
		}

		user, err := utils.CreateUser(db, username, password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{"code": http.StatusInternalServerError, "message": "Failed to make user"})
			log.Println("Error saving user", username, "to db:", err.Error())
			return
		}

		render.JSON(w, r, render.M{"code": http.StatusOK, "message": "Saved your info. Welcome " + user.Username})
	}
}

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []utils.UserModel{}
		transaction := db.Model(&utils.UserModel{}).Find(&users, "")

		if transaction.Error != nil {
			w.WriteHeader(500)
			render.JSON(w, r, render.M{"message": "Failed to get users", "code": 500, "error": transaction.Error.Error()})
		}

		if transaction.RowsAffected == 0 {
			log.Println("Got 0 rows for users")
		}
		render.JSON(w, r, render.M{"message": "OK", "code": http.StatusOK, "users": users, "count": transaction.RowsAffected})
	}
}
