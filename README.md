# API GOLANG
 
Test: Jorney Lopez


GitHub: https://github.com/jorneylm/API-go

# Get o clone Project
git@github.com:jorneylm/API-go.git


# Install
Driver for SQLite3
go get -u -v github.com/mattn/go-sqlite3


# Rutas

All Songs
http://localhost:8080/songs/

Get by Artist add "Colornoise" 
http://localhost:8080/artist/Colornoise

Get by Songs add "Gala"
http://localhost:8080/song/Gala

Get by Genres add "Rock"
http://localhost:8080/genres/1


Important: The server listens on port:8080

If you need to change, open main.go and go to http.ListenAndServe(":8080", mux). Put your port.


