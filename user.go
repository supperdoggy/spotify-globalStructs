package globalStructs

import "time"

type Statuses struct {
	Blocked bool `json:"blocked" bson:"blocked"`
	Admin   bool `json:"admin" bson:"admin"`
}

type User struct {
	ID string `json:"id" bson:"_id"`

	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`

	Statuses Statuses `json:"statuses" bson:"statuses"`

	LastOnline time.Time `json:"last_online" bson:"last_online"`
}
