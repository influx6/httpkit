package users

import "time"

//@httpapi
type User struct {
	PublicID string    `json:"public_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}
