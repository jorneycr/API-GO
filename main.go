package main

import (
	"Api-Go/controller"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs"), controller.Get_All_Songs)
	mux.HandleFunc(pat.Get("/songs/artist/:artist"), controller.Get_By_Artist)
	mux.HandleFunc(pat.Get("/songs/song/:song"), controller.Get_By_Song)
	mux.HandleFunc(pat.Get("/songs/genre/:genre"), controller.Get_By_Genre)
	http.ListenAndServe("localhost:8080", mux)
}
