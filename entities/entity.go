package entities

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

type Admin struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

type SuAdmin struct {
	Id       uint
	Name     string
	Email    string
	Password string
}
