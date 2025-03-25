package main

import (
	. "SanskritDictsApi/cmd/web"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	defer CloseAll()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/index", IndexHandler)
	mux.HandleFunc("/api/getSuggest", SuggestHandler)
	mux.HandleFunc("/api/search", SearchHandler)
	mux.HandleFunc("/api/list", SearchListHandler)
	mux.HandleFunc("/api/transliterate", TransliterateHandler)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	handler := cors.AllowAll().Handler(mux)
	err := http.ListenAndServe("127.0.0.1:4000", handler)
	log.Fatal(err)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

func addCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		// w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, req)
	})
}
