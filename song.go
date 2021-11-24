package globalStructs

import "time"

type Song struct {
	Name        string    `json:"name"`
	Album       string    `json:"album"`
	Band        string    `json:"band"`
	ReleaseDate time.Time `json:"release_date"`
	Path        string    `json:"path"`
}

type SongData struct {
	ID string `json:"id" mgo:"_id"`
	Data []byte `json:"data" mgo:"data"`
}
