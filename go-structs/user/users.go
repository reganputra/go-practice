package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// create constructor function to initialize User struct

func NewUser(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields must be provided")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// create constructor function to initialize Admin struct

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Aku",
			birthdate: "1/1/2000",
			createdAt: time.Now(),
		},
	}
}

func (user *User) OutputDetail() {
	fmt.Println("Outputting user details...")
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birthdate:", user.birthdate)
}

func (user *User) Remove() {
	fmt.Println("Removing user...")
	user.firstName = ""
	user.lastName = ""
}
