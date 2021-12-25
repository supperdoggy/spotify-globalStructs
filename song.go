package globalStructs

import "time"

type Song struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Album       string    `json:"album"`
	Band        string    `json:"band"`
	ReleaseDate time.Time `json:"release_date"`
	Path        string    `json:"path"`
	OwnerID     string    `json:"owner_id"`
}

type SongData struct {
	ID   string `json:"id" bson:"_id"`
	Data []byte `json:"data" bson:"data"`
}

type LikedSong struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	SongID string `json:"song_id"`
}
