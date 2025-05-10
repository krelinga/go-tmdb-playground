package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func getPosterURL(config *tmdb.Configuration, path string) string {
	size := config.Images.PosterSizes[len(config.Images.PosterSizes)-1]
	parts := []string{config.Images.SecureBaseURL, size, path}
	return strings.Join(parts, "")
}

func GetBackdropURL(config *tmdb.Configuration, path string) string {
	size := config.Images.BackdropSizes[len(config.Images.BackdropSizes)-1]
	parts := []string{config.Images.SecureBaseURL, size, path}
	return strings.Join(parts, "")
}

func demo() error {
	client, err := newClient()
	if err != nil {
		return err
	}
	config, err := client.GetConfiguration()
	if err != nil {
		return err
	}

	fmt.Print("Partial Title: ")
	partialTitle, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	movies, err := client.SearchMovie(partialTitle, nil)
	if err != nil {
		return err
	}
	fmt.Println("Search Results:")
	for i, movie := range movies.Results {
		if i >= 5 {
			break
		}
		fmt.Printf("%d) %s (ID: %d)\n", i, movie.Title, movie.ID)
	}
	fmt.Print("Select a movie by number: ")
	var choice int
	if _, err = fmt.Scanf("%d", &choice); err != nil {
		return err
	}
	if choice < 0 || choice >= len(movies.Results) {
		return fmt.Errorf("invalid choice: %d", choice)
	}
	movieID := movies.Results[choice].ID
	info, err := client.GetMovieInfo(movieID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Title: %s\n", info.Title)
	fmt.Printf("Release Date: %s\n", info.ReleaseDate)
	fmt.Printf("Overview: %s\n", info.Overview)
	fmt.Printf("Genres:\n")
	for _, genre := range info.Genres {
		fmt.Printf("- %s\n", genre.Name)
	}
	fmt.Printf("Poster Path: %s\n", getPosterURL(config, info.PosterPath))
	fmt.Printf("Back Drop Path: %s\n", GetBackdropURL(config, info.BackdropPath))

	return nil
}

func main() {
	if err := demo(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
