package health

import (
	"log/slog"
	"net/http"

	"github.com/CodeMaverick-143/Golang_learning/internal/utils/response"
)

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Health check request")
		_ = response.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}
