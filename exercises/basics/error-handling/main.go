package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user, err := createUser("Fatih", 22)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user)
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age < 18 {
		return errors.New("age must be at least 18")
	}

	return nil
}

func createUser(name string, age int) (User, error) {
	err := validateAge(age)

	if err != nil {
		return User{}, fmt.Errorf("create user failed: %w", err)
	}

	return User{
		Name: name,
		Age:  age,
	}, nil
}