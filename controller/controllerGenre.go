package controller

import (
	API "Api-Go/handler"
	"net/http"

	"goji.io/pat"
)

func Get_By_Genre(w http.ResponseWriter, r *http.Request) {
	var genre = pat.Param(r, "genre")
	response := API.Fetch_Get_By_Genre(genre)
	w.Write(response)
}
