package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fredrikaugust/runlog/storage"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *storage.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var body RegisterBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		if body.Email == "" || body.Password == "" {
			http.Error(w, "email and password can't be empty", http.StatusBadRequest)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			slog.Error("failed to hash user password", "error", err.Error())
			return
		}

		if err := db.CreateUser(r.Context(), body.Email, passwordHash); err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			slog.Error("failed to create user", "error", err.Error())
			return
		}

		w.Write([]byte("user created"))
	}
}
