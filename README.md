# grpc-game-example

An example game built using Go, gRPC, and tview.

The blog post for this project is a good reference: https://mortenson.coffee/blog/making-multiplayer-game-go-and-grpc

I built this as a way to learn more about Go and haven't been using the
language for that long so don't judge the code too harshly!

## Game description

This is “tshooter” - a local or online multiplayer shooting game you play
in your terminal. Players can move in a map and fire lasers at other players.
When a player is hit, they respawn on the map and the shooting player’s score
is increased. When a player reaches 10 kills, the round ends and a new round
begins. You can play the game offline with bots, or online with up to eight
players (but that limit is arbitrary).

## Reference and use

Here's a quick reference for common operations on the project:

```bash
# Download go module dependencies
go get ./...
# Build binaries
make build
# Rebuild protobuf
make proto
# Run gofmt
make fmt
```

If you run the commands or binaries directly more command line options are available:

```bash
# Run a server
./tshooter server -port=9999 -bots=2 -password=foo
# Run a local, offline game
./tshooter local -bots=2
# Run a bot as a client
./tshooter bot -address=":9999"
```
