# API GOLANG
 
Test: Jorney Lopez


GitHub: https://github.com/jorneylm/API-GO

# Get o clone Project
git@github.com:jorneylm/API-GO.git


# Install
Glide (To managment package) I install with Visual Studio Code (VS code Glide), then Open terminal on VS Code >Glide:Install.

go build 

go run *.go


# Rutas

All Songs
http://localhost:8080/songs

Get by Artist add "Colornoise" 
http://localhost:8080/songs/artist/Colornoise

Get by Songs add "Gala"
http://localhost:8080/songs/song/Gala

Get by Genres add "Rock"
http://localhost:8080/songs/genres/Rock


Important: The server listens on port:8080

If you need to change, open main.go and go to http.ListenAndServe(":8080", mux). Put your port.


