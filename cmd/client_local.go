package cmd

// Starts a local instance of the game with bots.

import (
	"fmt"
	"log"
	"os"

	termutil "github.com/andrew-d/go-termutil"
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/chingkamhing/grpc-game-example/pkg/backend"
	"github.com/chingkamhing/grpc-game-example/pkg/bot"
	"github.com/chingkamhing/grpc-game-example/pkg/frontend"
)

var cmdLocal = &cobra.Command{
	Use:   "local",
	Short: "T Shooter game (local)",
	Run:   runLocal,
}

var clientBots int

func init() {
	cmdLocal.Flags().IntVar(&clientBots, "bots", 1, "The number of bots to play against.")

	rootCmd.AddCommand(cmdLocal)
}

func runLocal(cmd *cobra.Command, args []string) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		panic("this program must be run in a terminal")
	}

	currentPlayer := backend.Player{
		Name:            "Alice",
		Icon:            'A',
		IdentifierBase:  backend.IdentifierBase{uuid.New()},
		CurrentPosition: backend.Coordinate{X: -1, Y: -5},
	}
	game := backend.NewGame()
	game.AddEntity(&currentPlayer)

	view := frontend.NewView(game)
	view.CurrentPlayer = currentPlayer.ID()

	bots := bot.NewBots(game)
	for i := 0; i < clientBots; i++ {
		bots.AddBot(fmt.Sprintf("Bob %d", i))
	}

	game.Start()
	view.Start()
	bots.Start()

	err := <-view.Done
	if err != nil {
		log.Fatal(err)
	}
}
