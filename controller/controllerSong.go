package controller

import (
	API "Api-Go/handler"
	"net/http"

	"goji.io/pat"
)

func Get_All_Songs(w http.ResponseWriter, r *http.Request) {
	z := API.Fetch_Get_All_Songs()
	w.Write(z)
}

func Get_By_Song(w http.ResponseWriter, r *http.Request) {
	var song = pat.Param(r, "song")
	response := API.Fetch_Get_By_Song(song)
	w.Write(response)
}
