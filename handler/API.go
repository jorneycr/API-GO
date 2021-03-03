package handler

import (
	// Driver para SQLite3
	"Api-Go/modelo"
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Fetch_Get_All_Songs() []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre;"
	return Fetch_Generic_Songs(query)
}

func Fetch_Get_By_Artist(artist string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.artist = \"" + artist + "\";"
	return Fetch_Generic_Songs(query)
}

func Fetch_Get_By_Song(song string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.song = \"" + song + "\";"
	return Fetch_Generic_Songs(query)
}

func Fetch_Get_By_Genre(genre string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where g.name = \"" + genre + "\";"
	return Fetch_Generic_Songs(query)
}

func Fetch_Generic_Songs(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	collection := []modelo.Songs{}
	for rows.Next() {
		var id int
		var artist string
		var song string
		var genre string
		var length int
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		if err != nil {
			log.Fatal(err)
		}

		temp := modelo.Songs{
			Id:     id,
			Artist: artist,
			Song:   song,
			Genre:  genre,
			Length: length,
		}
		collection = append(collection, temp)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := modelo.SongsCollection{
		Songs: collection,
	}

	z, _ := json.Marshal(y)

	return z
}
