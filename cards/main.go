package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard() -> := only need for initialization
	// card = "Five of Diamonds"
	// cards := newCard()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
