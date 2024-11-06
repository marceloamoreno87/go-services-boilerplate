package helpers

import (
	"net/http"
)

func GetAllFilters(r *http.Request) map[string]string {
	filters := make(map[string]string)

	// Define valores padrão
	defaultLimit := "10"
	defaultPage := "1"

	// Adiciona parâmetros da URL ao mapa de filtros
	for key, value := range r.URL.Query() {
		filters[key] = value[0]
	}

	// Verifica se 'limit' está presente e não está em branco, caso contrário, define o valor padrão
	if limit, exists := filters["limit"]; !exists || limit == "" {
		filters["limit"] = defaultLimit
	}

	// Verifica se 'page' está presente e não está em branco, caso contrário, define o valor padrão
	if page, exists := filters["page"]; !exists || page == "" {
		filters["page"] = defaultPage
	}

	return filters
}
