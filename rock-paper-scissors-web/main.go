package main

import (
	"log"
	"net/http"
	"rock-paper-scissors-web/handlers"
)

func main() {
	// Crear router
	router := http.NewServeMux()

	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchado en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
