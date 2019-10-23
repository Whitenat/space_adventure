package main

import (
	"fmt"
	"strings"
)

func welcomeMessage(){
	fmt.Println("welcome to the solar system!")
	fmt.Println("There are 9 planets to explore!")

	fmt.Println("What is your name?")
	var name string    
    fmt.Scanln(&name)
	fmt.Println("Nice to meet you " + name + " My name is Eliza, I'm an old friend of Alexa")
	
	fmt.Println("Let's go on an adventure!")
	
}

func randomPrompt(){
	fmt.Println("Shall I randomly choose a planet for you to vist? (Y/N)")
	var randomplanet string    
	fmt.Scanln(&randomplanet)
	checkvalid(randomplanet)
}

func checkvalid(usrinput string){
	var upperinput string = strings.ToUpper(usrinput)
	if upperinput != "Y" && upperinput != "N"{
		fmt.Println("Please enter a valid input")
		randomPrompt()
	}else{
		if upperinput == "Y"{
			fmt.Println("Traveling to Mars...")
		}else{
			choosePlanetPrompt()
		}
	}
}


func choosePlanetPrompt(){
	fmt.Println("Name the planet you would like to vist:")
	var desiredplanet string    
	fmt.Scanln(&desiredplanet)
	validPlanet(desiredplanet)
	fmt.Println("Traveling to " + desiredplanet + "...")
}

func validPlanet(usrinput string){
	fmtinput := strings.ToUpper(usrinput)
	planets := []string{"MERCURY","VENUS","EARTH", "MARS","JUPITER","SATURN", "URANUS", "NEPTUNE","PLUTO"}
	for _, planet := range planets {
    	if fmtinput == planet {
			return
		}
	}
	fmt.Println("Please enter a valid input")
	choosePlanetPrompt()
}

func main(){
	welcomeMessage()
	randomPrompt()

}

