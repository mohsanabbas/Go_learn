package main

func main() {
	cards := newDeck()
	// 	cards.saveToFile("deck")
	// cards := getDeckFromFile("deck")
	cards.shufle()
	cards.print()
}
