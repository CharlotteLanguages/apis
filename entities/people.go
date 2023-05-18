package entities

import (
	"fmt"
)

type People struct {
	Id        int64  `json: "idPerson" sql:"primary_key; auto_increment"`
	Name      string `json: "name"`
	Lastname  string `json: "lastname"`
	Birthdate string `json: "birthdate"`
	Gender    string `json: "gender"`
	Email     string `json: "email"`
	Username  string `json: "username"`
	Password  string `json: "password"`
}

func (people People) ToString() string {
	return fmt.Sprintf("id: %d name: %s lastname: %s birthdate: %s gender: %s email: %s username: %s password: %s",
		people.Id, people.Name, people.Lastname, people.Birthdate, people.Gender, people.Email, people.Username, people.Password)
}
