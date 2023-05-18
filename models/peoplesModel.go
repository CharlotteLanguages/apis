package models

import (
	"charlotte_backend/entities"
	"database/sql"
)

type PeopleModel struct {
	Db *sql.DB
}

// show everything
func (peoplesModel PeopleModel) FindAll() (peoples []entities.People, err error) {
	rows, err := peoplesModel.Db.Query("SELECT * FROM PERSON")
	if err != nil {
		return nil, err
	} else {
		var peoples []entities.People
		for rows.Next() {
			var idPerson int64
			var name string
			var lastname string
			var birthdate string
			var gender string
			var email string
			var username string
			var password string

			err2 := rows.Scan(&idPerson, &name, &lastname, &birthdate, &gender, &email, &username, &password)

			if err2 != nil {
				return nil, err2
			} else {
				people := entities.People{
					Id:        idPerson,
					Name:      name,
					Lastname:  lastname,
					Birthdate: birthdate,
					Gender:    gender,
					Email:     email,
					Username:  username,
					Password:  password,
				}
				peoples = append(peoples, people)
			}
		}
		return peoples, nil
	}
}

// Search for username
func (peoplesModel PeopleModel) Search(keyword string) (peoples []entities.People, err error) {
	rows, err := peoplesModel.Db.Query("SELECT * FROM PERSON where username like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var peoples []entities.People
		for rows.Next() {
			var idPerson int64
			var name string
			var lastname string
			var birthdate string
			var gender string
			var email string
			var username string
			var password string

			err2 := rows.Scan(&idPerson, &name, &lastname, &birthdate, &gender, &email, &username, &password)

			if err2 != nil {
				return nil, err2
			} else {
				people := entities.People{
					Id:        idPerson,
					Name:      name,
					Lastname:  lastname,
					Birthdate: birthdate,
					Gender:    gender,
					Email:     email,
					Username:  username,
					Password:  password,
				}
				peoples = append(peoples, people)
			}
		}
		return peoples, nil
	}
}

// Create people
func (peoplesModel PeopleModel) Create(people *entities.People) (err error) {
	result, err := peoplesModel.Db.Exec("INSERT INTO PERSON (name, lastname, birthdate, gender, email, username, password) VALUES (?, ?, ?, ?, ?, ?, ?)",
		people.Name, people.Lastname, people.Birthdate, people.Gender, people.Email, people.Username, people.Password)
	if err != nil {
		return err
	} else {
		people.Id, _ = result.LastInsertId()
		return nil
	}
}

// Update people
func (peoplesModel PeopleModel) Update(people *entities.People) (int64, error) {
	result, err := peoplesModel.Db.Exec("UPDATE PERSON SET name = ?, lastname = ?, birthdate = ?, gender = ?, email = ?, username = ?, password = ? WHERE idPerson = ?",
		people.Name, people.Lastname, people.Birthdate, people.Gender, people.Email, people.Username, people.Password, people.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// Delete people
func (peoplesModel PeopleModel) Delete(id int64) (int64, error) {
	result, err := peoplesModel.Db.Exec("DELETE FROM PERSON WHERE idPerson = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
