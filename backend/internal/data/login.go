package data

import (
	"encoding/json"
	"io"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) JsonDecode(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}