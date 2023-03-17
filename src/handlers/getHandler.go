package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"server/src/dependencies"

	"github.com/redis/go-redis/v9"
)

func getHandler(env *dependencies.ServerDep, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	
	value, err := env.Redis.Get(r.Context(), key).Result()
	if errors.Is(err, redis.Nil) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Key not found")
		return
	} else if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "%v=%v", key, value)
}
