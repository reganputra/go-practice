package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	createdAt time.Time
}

// create constructor function to initialize User struct

func NewUser(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields must be provided")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (user *User) OutputDetail() {
	fmt.Println("Outputting user details...")
	fmt.Println("First Name:", user.FirstName)
	fmt.Println("Last Name:", user.LastName)
	fmt.Println("Birthdate:", user.Birthdate)
}

func (user *User) Remove() {
	fmt.Println("Removing user...")
	user.FirstName = ""
	user.LastName = ""
}
