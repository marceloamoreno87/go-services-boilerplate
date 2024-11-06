package helpers

import (
	"net/http"
	"time"
)

type TokensInterface interface {
	GetAccessToken() string
	GetRefreshToken() string
}

func SetTokensCookie(w http.ResponseWriter, tokens TokensInterface) {

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokens.GetAccessToken(),
		HttpOnly: true,                             // Evita acesso via JavaScript
		Secure:   true,                             // Apenas para HTTPS
		SameSite: http.SameSiteLaxMode,             // Ajuda a prevenir CSRF
		Path:     "/",                              // O cookie estará disponível para todas as rotas
		Expires:  time.Now().Add(15 * time.Minute), // Expira em 15 minutos (ajuste conforme necessário)
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.GetRefreshToken(),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour), // Expira em 7 dias (ajuste conforme necessário)
	})
}
