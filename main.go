package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./statics"))

	http.Handle("/statics/", http.StripPrefix("/statics/", fs))
	db := http.FileServer(http.Dir("./java"))

	http.Handle("/java/", http.StripPrefix("/java/", db))
	http.HandleFunc("/", HomeHandler)
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
