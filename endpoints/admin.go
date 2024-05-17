package endpoints

import (
	"log"
	"net/http"

	"github.com/Aervyon/go-playground/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

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
