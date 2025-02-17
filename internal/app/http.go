package app

import (
	"agenti/internal/currency"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/currencies", currency.Index)
	return mux
}
