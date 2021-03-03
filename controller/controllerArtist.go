package controller

import (
	API "Api-Go/handler"
	"net/http"

	"goji.io/pat"
)

func Get_By_Artist(w http.ResponseWriter, r *http.Request) {
	var artist = pat.Param(r, "artist")
	response := API.Fetch_Get_By_Artist(artist)
	w.Write(response)
}
