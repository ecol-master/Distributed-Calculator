package main

import (
	"encoding/json"
	"io"
	"os"
)

type User struct {
	IsAuth   bool
	Login    string
	Password string
	ID       int
}

const userConfigFilename = "user_data.json"

func loadUserFromFile() (*User, error) {
	file, err := os.OpenFile(userConfigFilename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
