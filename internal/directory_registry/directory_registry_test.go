package directoryregistry

import (
	"fmt"
	"testing"
)

func TestPathToRepo(t *testing.T) {

	t.Run("path to repo", func(t *testing.T) {
		p := PathToRepository()

		fmt.Printf("path to the repo is %s", p)
	})

}

func TestPathToGames(t *testing.T) {

	t.Run("path to repo", func(t *testing.T) {
		p := PathToGames()

		fmt.Printf("path to the games dir is %s", p)
	})

}
