package main

func main() {
	cards := newDeckFromFile("ASD")
	cards.shuffle()
	cards.print()
}
