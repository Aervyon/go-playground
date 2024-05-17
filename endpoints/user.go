package endpoints

import (
	"errors"
	"net/http"

	"log"

	"github.com/Aervyon/go-playground/database"
	"github.com/Aervyon/go-playground/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

var (
	ResponseUnauthorized = render.M{"code": http.StatusUnauthorized, "message": "Unauthorized"}
)

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
		existingUser := &models.UserModel{}
		db.Model(&models.UserModel{}).Find(&existingUser, "username = ?", username)
		if existingUser.ID != "" {
			w.WriteHeader(http.StatusIMUsed)
			render.JSON(w, r, render.M{"code": http.StatusIMUsed, "message": "Username or email taken"})
			return
		}

		user, err := database.CreateUser(db, username, password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{"code": http.StatusInternalServerError, "message": "Failed to make user"})
			log.Println("Error saving user", username, "to db:", err.Error())
			return
		}

		render.JSON(w, r, render.M{"code": http.StatusOK, "message": "Saved your info. Welcome " + user.Username})
	}
}

func GetSelfAccount(db *gorm.DB, session *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := session.GetString(r.Context(), "session")
		if uid == "" {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, ResponseUnauthorized)
			return
		}

		var user models.UserModel
		transaction := db.Model(&models.UserModel{}).Limit(1).Find(&user, "ID = ?", uid)
		if transaction.Error != nil && errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
			log.Println("No record found for authentication user", uid)
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, render.M{
				"code":    http.StatusNotFound,
				"message": "Account " + uid + " not found",
			})
		}
		if transaction.Error != nil {
			log.Println(transaction.Error.Error())
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"code":    http.StatusInternalServerError,
				"message": "Error Occurred getting your account",
				"error":   transaction.Error.Error(),
			})
			return
		}

		render.JSON(w, r, render.M{
			"code":    http.StatusOK,
			"message": "OK",
			"account": user,
		})
	}
}

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []models.UserModel{}
		transaction := db.Model(&models.UserModel{}).Find(&users, "")

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
