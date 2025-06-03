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
