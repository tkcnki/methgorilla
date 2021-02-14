package entity

var Users []*User

type User struct {
	Id   string `json:id`
	Name string `json:name`
	Kana string `json:kana`
}
