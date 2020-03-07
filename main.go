package main

func main() {
	newCardsDeck := NewDeckFromFile("test.txt")
	newCardsDeck.Shuffle()
	newCardsDeck.Print()
}
