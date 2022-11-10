package ginserver

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func handler() {
	if err := recover(); err != nil {
		erro := fmt.Sprintf("%v", err)
		log.Println(("Erro panic na aplicação: " + erro))
		log.Println(string(debug.Stack()))
	}
	r := setupRouter()
	addr := ":3001"

	server := http.Server{
		Addr:        addr,
		Handler:     r,
		ReadTimeout: 600 * time.Second,
		IdleTimeout: 30 * time.Second,
	}
	server.ListenAndServe()
}
