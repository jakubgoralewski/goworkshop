package users

type User struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"wiek"`
	phone   string
}

func NewUser(id uint64, name string, surname string, age int) *User {
	return &User{
		Id:      id,
		Name:    name,
		Surname: surname,
		Age:     age,
	}
}
