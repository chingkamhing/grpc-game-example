package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/mortenson/grpc-game-example/pkg/backend"
	"github.com/mortenson/grpc-game-example/pkg/bot"
	"github.com/mortenson/grpc-game-example/pkg/frontend"
)

func main() {
	currentPlayer := backend.Player{
		Name:            "Alice",
		Icon:            'A',
		IdentifierBase:  backend.IdentifierBase{uuid.New()},
		CurrentPosition: backend.Coordinate{X: -1, Y: -5},
	}
	game := backend.NewGame()
	game.AddEntity(&currentPlayer)
	// game.AddEntity(&backend.Player{
	// 	Name:            "Bob",
	// 	Icon:            'B',
	// 	IdentifierBase:  backend.IdentifierBase{uuid.New()},
	// 	CurrentPosition: backend.Coordinate{X: 0, Y: 0},
	// })
	bots := bot.NewBots(game)
	bots.AddBot("Dave")

	view := frontend.NewView(game)
	view.CurrentPlayer = currentPlayer.ID()

	game.Start()
	view.Start()
	bots.Start()

	err := <-view.Done
	if err != nil {
		log.Fatal(err)
	}
}
