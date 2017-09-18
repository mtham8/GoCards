package main

// go build --> compiles a bunch of go source code files
// go run --> compiles and executes one or two files
// go fmt --> formats all the code in each file in the current directory
// go install --> compiles and "installs" a package
// go get --> downloads the raw source code of someone else's package
// go test --> run any tests associated with the current project
// types: bool, string, int, float64

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades") // does not modify existing slice

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
