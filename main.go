package main

import (
	"fmt"
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

func newClient() (*tmdb.TMDb, error) {
	apiKey, err := readApiKey()
	if err != nil {
		return nil, err
	}

	config := tmdb.Config{
		APIKey: apiKey,
	}
	return tmdb.Init(config), nil
}

func demo() error {
	client, err := newClient()
	if err != nil {
		return err
	}

	movieID := 550 // Example movie ID for "Fight Club"
	movie, err := client.GetMovieInfo(movieID, nil)
	if err != nil {
		return err
	}

	fmt.Println(movie)

	return nil
}

func main() {
	if err := demo(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
