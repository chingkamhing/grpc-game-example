package cmd

// Runs a game server.

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/chingkamhing/grpc-game-example/pkg/backend"
	"github.com/chingkamhing/grpc-game-example/pkg/bot"
	"github.com/chingkamhing/grpc-game-example/pkg/server"
	"github.com/chingkamhing/grpc-game-example/proto"

	"google.golang.org/grpc"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "T Shooter game (server)",
	Run:   runServer,
}

var port int
var password string
var numBots int

func init() {
	cmdServer.Flags().IntVar(&port, "port", defaultPort, "The port to listen on.")
	cmdServer.Flags().StringVar(&password, "password", "", "The server password.")
	cmdServer.Flags().IntVar(&numBots, "bots", 0, "The number of bots to add to the server.")

	rootCmd.AddCommand(cmdServer)
}

func runServer(cmd *cobra.Command, args []string) {
	log.Printf("listening on port %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	game := backend.NewGame()

	bots := bot.NewBots(game)
	for i := 0; i < numBots; i++ {
		bots.AddBot(fmt.Sprintf("Bob %d", i))
	}

	game.Start()
	bots.Start()

	s := grpc.NewServer()
	server := server.NewGameServer(game, password)
	proto.RegisterGameServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
