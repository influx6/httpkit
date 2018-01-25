package fixtures

import (
	"encoding/json"

	"github.com/gokit/httpkit/example"
)

// json fixtures ...
var (
	UserJSON = `{


    "public_id":	"ur3rzf9veoxi9c5vsbnfi6gr9dwb7c",

    "username":	"RaymondHall",

    "email":	"JohnThomas@Voonte.mil",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

}`

	UserCreateJSON = `{


    "public_id":	"4e1phhg5ilwodfzffdopvkevvcr1di",

    "username":	"possimus",

    "email":	"uGonzales@Photospace.org",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

}`

	UserUpdateJSON = `{


    "public_id":	"37w81h78rz3n8taa5sprlnzbos5ibw",

    "username":	"fGomez",

    "email":	"PeterNguyen@Realbuzz.net",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

}`
)

// LoadCreateJSON returns a new instance of a users.User.
func LoadCreateJSON(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return users.User{}, err

	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a users.User.
func LoadUpdateJSON(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {

		return users.User{}, err

	}

	return elem, nil
}

// LoadUserJSON returns a new instance of a users.User.
func LoadUserJSON(content string) (users.User, error) {
	var elem users.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return users.User{}, err
	}

	return elem, nil
}
