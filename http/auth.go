package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/Aervyon/go-playground/models"
	"github.com/Aervyon/go-playground/utils"
	"github.com/alexedwards/argon2id"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func AuthUser(db *gorm.DB, sessionManager *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		sessionManager.Put(r.Context(), "auth", token.ID+"."+token.Token)

		render.JSON(w, r, render.M{"code": 201, "message": "Created Token", "token": token.Token, "tokenType": "Bearer"})
	}
}

func CheckSession(db *gorm.DB, sessionManager *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := sessionManager.GetString(r.Context(), "auth")

		pair := strings.Split(auth, ".")
		if len(pair) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, render.M{"code": http.StatusUnauthorized, "message": "Unauthorized"})
			return
		}

		// find the session in the database
		var token models.Token
		db.Model(&models.Token{}).Find(&token, "ID = ?", pair[0])
		if token.Hash == "" {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, render.M{"code": http.StatusUnauthorized, "message": "Unauthorized"})
			return
		}

		isValid, _, err := argon2id.CheckHash(pair[1], token.Hash)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"code":    http.StatusInternalServerError,
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, render.M{"code": http.StatusUnauthorized, "message": "Unauthorized"})
			return
		}

		render.JSON(w, r, render.M{"code": http.StatusOK, "message": "Hi user with ID" + token.UID})
	}
}
