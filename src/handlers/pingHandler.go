package handlers

import (
	"fmt"
	"net/http"
	"server/src/dependencies"
)

func pingHandler(env *dependencies.ServerDep, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
