package main

import "log"

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()

	starting, end := cards.mDeal(3)
	firstDeck, EndDeck := deal(cards, 3)

	starting.print()
	end.print()

	firstDeck.print()
	EndDeck.print()

	fn := "card_game.txt"
	err := cards.saveToFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	cards_from_file := readDeckFromFile("card_game.txt")
	cards_from_file.shuffle()
	cards_from_file.print()

}
