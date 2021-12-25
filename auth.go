package globalStructs

import "time"

type Token struct {
	UserID   string    `json:"user_id"`
	TokenStr string    `json:"token_str"`
	Expire   time.Time `json:"expire"`
}

type Creds struct {
	ID       string `json:"id" bson:"_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Password []byte `json:"password" bson:"password"`
}
