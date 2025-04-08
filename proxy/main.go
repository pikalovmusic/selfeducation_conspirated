package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)
	r.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
		log.Println("API request")
	})

	fmt.Println("Server started at :8080...")
	http.ListenAndServe(":8080", r)
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprint(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
