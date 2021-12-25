package globalStructs

import "time"

type Playlist struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	Songs       []Song    `json:"songs"`
	Created     time.Time `json:"created"`
	Shared      bool      `json:"shared"`
}

type ShortPlaylist struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Shared      bool      `json:"shared"`
	Created     time.Time `json:"created"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
}
