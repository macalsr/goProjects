package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Print("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = !isHeistOn
		fmt.Print("Plan a better disguise next time?")
	}
	fmt.Println("isHeistOn is currently:", isHeistOn)
	openedVault := rand.Intn(100)

	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == true && openedVault < 70 {
		fmt.Println("The vault can't be open!")
	}
	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {

		case 0:
			isHeistOn = false
			fmt.Print("Looks like you tripped an alarm... run?!")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")

		case 2:
			isHeistOn = false
			fmt.Print("You forgot to pick up your gloves... Maybe tomorrow? ")
		case 3:
			isHeistOn = false
			fmt.Print("There's dogs inside the vault!! Run!")
		default:
			fmt.Print("Start the getaway car!")
		}

	}
	if isHeistOn == true {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Print(amtStolen)
	}
}
