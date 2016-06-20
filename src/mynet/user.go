package mynet

type User struct {
	Id      int
	Name    string
	Passwd  string
	Friends string
	Other   string
}

func NewUser(id int, name string, pass string, fri string, other string) *User {
	return &User{id, name, pass, fri, other}
}
