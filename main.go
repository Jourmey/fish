package main

func main() {
	game, gameErr := NewGame()
	if gameErr != nil {
		panic(gameErr)
	}

	if runErr := game.Run(); runErr != nil {
		panic(runErr)
	}
}
