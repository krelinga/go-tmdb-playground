package main

import (
	"os"
)

func readApiKey() (string, error) {
	data, err := os.ReadFile(".api_key")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
