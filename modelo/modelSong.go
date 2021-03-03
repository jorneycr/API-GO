package modelo

type Songs struct {
	Id     int
	Artist string
	Song   string
	Genre  string
	Length int
}

type SongsCollection struct {
	Songs []Songs
}
