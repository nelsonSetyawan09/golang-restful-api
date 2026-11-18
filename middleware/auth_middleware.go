package middleware

import (
	"belajar-golang-restful-api/config"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
	Config  *config.Config
}

func NewAuthMiddleware(handler http.Handler, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
		Config:  cfg,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("X-API-Key")

	if apiKey == middleware.Config.APIKey {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}

	helper.WriteToResponseBody(w, webResponse)
}
