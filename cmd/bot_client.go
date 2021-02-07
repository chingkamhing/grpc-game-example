package cmd

// Connects a local bot to a remote server, which is a neat demo.
// The server has no awareness that a bot is controlling the player.

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/chingkamhing/grpc-game-example/pkg/backend"
	"github.com/chingkamhing/grpc-game-example/pkg/bot"
	"github.com/chingkamhing/grpc-game-example/pkg/client"
	"github.com/chingkamhing/grpc-game-example/pkg/frontend"
	"github.com/chingkamhing/grpc-game-example/proto"
	"google.golang.org/grpc"
)

var cmdBot = &cobra.Command{
	Use:   "bot",
	Short: "T Shooter game (bot client)",
	Run:   runBot,
}

var address string

func init() {
	cmdBot.Flags().StringVar(&address, "address", fmt.Sprintf(":%d", defaultPort), "The server address.")

	rootCmd.AddCommand(cmdBot)
}

func runBot(cmd *cobra.Command, args []string) {
	game := backend.NewGame()
	game.IsAuthoritative = false
	view := frontend.NewView(game)
	game.Start()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	grpcClient := proto.NewGameClient(conn)
	client := client.NewGameClient(game, view)

	bots := bot.NewBots(game)
	player := bots.AddBot("Bob")

	err = client.Connect(grpcClient, player.ID(), player.Name, "")
	if err != nil {
		log.Fatalf("connect request failed %v", err)
	}
	client.Start()

	view.Start()
	bots.Start()

	err = <-view.Done
	if err != nil {
		log.Fatal(err)
	}
}
