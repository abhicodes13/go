package handlers

import (
	"encoding/json"
	"net/http"

	"task-api/internal/models"

	"github.com/jackc/pgx/v5"
)

func Tasks(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:

			var task models.Task

			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			err = db.QueryRow(
				r.Context(),
				`INSERT INTO tasks (title, completed)
				 VALUES ($1, $2)
				 RETURNING id`,
				task.Title,
				task.Completed,
			).Scan(&task.ID)

			if err != nil {
				http.Error(w, "Failed to create task", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}