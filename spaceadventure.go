package main

import fmt

func main(){
	fmt.Println("welcome to the solar system!/n")
	fmt.Println("There are 9 planets to explore!/n")
	fmt.Println("What is your name?/n")
	name := ""
	fmt.Scanln(name)
	fmt.Println("Nice to meet you " name " My name is Eliza, I'm an old friend of Alexa/n")
	fmt.Println("Let's go on an adventure!")
	fmt.Println("Shall I randomly choose a planet for you to vist? (Y/N)/n")
	randomplanet :=""
	fmt.Scanln(randomplanet)
	fmt.Println("Name the planet you would like to vist:")
	desiredplanet := ""
	fmt.Scanln(desiredplanet)
	fmt.Println("Traveling to " desiredplanet "...")



}