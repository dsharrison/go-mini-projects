package main

func main() {
	cards := newDeckFromFile("MyCards1")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
