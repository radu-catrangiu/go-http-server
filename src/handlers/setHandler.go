package handlers

import (
	"fmt"
	"net/http"
	"server/src/dependencies"
)

func setHandler(env *dependencies.ServerDep, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	
	err := env.Redis.Set(r.Context(), key, value, 0).Err()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "Value set %s - %s", key, value)
}
