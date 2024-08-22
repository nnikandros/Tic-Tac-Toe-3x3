package directoryregistry

import (
	"os"
	"path/filepath"
)

func PathToRepository() string {

	p, _ := os.Getwd()
	filepath.Dir(p)
	return filepath.Dir(filepath.Dir(p))

}

func PathToGames() string {

	p := PathToRepository()

	games_dir := filepath.Join(p, "games")

	return games_dir

}

// Create the directories necessary for storing. Currently we only save the games so we need app/games dir
func CreateDirs() error {
	var GAMES_DIR = filepath.Join(APP_DIR, "games")

	err := os.Mkdir(GAMES_DIR, 0755)

	if err != nil {
		return err
	}

	return nil

}

var p, _ = os.Getwd()

var APP_DIR = filepath.Dir(filepath.Dir(filepath.Dir(p)))

var GAMES_DIR = filepath.Join(APP_DIR, "games")
