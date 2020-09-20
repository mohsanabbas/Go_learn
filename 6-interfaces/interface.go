package main

import "fmt"

type spanishBot struct{}
type englishBot struct{}

type bot interface {
	getAbuse() string
}

// dont wanna get abused multple time use interface and get it once

func main() {
	sb := spanishBot{}
	eb := englishBot{}
	printAbuse(sb)
	printAbuse(eb)

}

// dont wanna use printAbuse fot spanishBot and englishBot differently
func printAbuse(b bot) {
	fmt.Println(b.getAbuse())
}

// if there is no use of reciever param we can just use reciever type ex. spanishBot
func (spanishBot) getAbuse() string {
	return "grande volumen de mie***!!"
}
func (englishBot) getAbuse() string {
	return "Pile of sh** !!"
}
