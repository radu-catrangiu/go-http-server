package handlers

import (
	"log"
	"net/http"
	"server/src/dependencies"
)

func RegisterHandlers(env *dependencies.ServerDep) {
	log.Println("Registering handlers...")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingHandler(env, w, r)
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		setHandler(env, w, r)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		getHandler(env, w, r)
	})
}
