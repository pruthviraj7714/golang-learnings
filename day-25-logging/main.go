package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

var logger *zap.Logger

func initLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Info("Incoming Request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
		)

		next.ServeHTTP(w, r)

	})
}

func main() {
	//sync the logger before exiting - sync flushes any buffered log entries - important for development
	initLogger()
	defer logger.Sync()

	//logging levels
	logger.Info("info log")
	logger.Error("error log")
	logger.Warn("warning log")
	logger.Debug("debug log")
	logger.Info("user info", zap.String("user", "John"), zap.Int("id", 1))

	user := struct {
		ID   string
		Name string
	}{
		ID:   "1",
		Name: "John",
	}

	//add context to logs
	logger.Info("User created",
		zap.String("user_id", user.ID),
		zap.String("name", user.Name),
	)

	r := chi.NewRouter()

	r.Use(loggingMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("serving homepage", zap.String("url", r.URL.Path))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
		})
	})

	http.ListenAndServe(":8080", r)

}
