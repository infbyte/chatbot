package main

import (
	"bufio"
	"fmt"
	"os"
	//"bufio"
	//"os"
	//. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("Hello there. My name is Clara. I am a chatbot.")

	scanner := bufio.NewScanner(os.Stdin)

	// Outside instances
	var text string
	var name string
	var about string

	//for text != "/clara exit" { // break the loop if entered

	fmt.Println("Enter 'Y' to start chatting")
	scanner.Scan()
	text = scanner.Text()

	if text == "owo" {
		fmt.Println("What's this?")
	}

	if text == "y" || text == "Y" {
		// fmt.Println("Your text was: ", text)

		fmt.Println("Thanks for chatting with me. What's your name?")

		scanner.Scan()
		name = scanner.Text()

		fmt.Println("Lovely name, ", name)
		fmt.Println("To learn more about me say /about, otherwise start a conversation.")

		scanner.Scan()
		about = scanner.Text()

		if about == "/about" || about == "about" {
			fmt.Println("Thanks for asking ", name)
			fmt.Println("I am a GitHub repository by infbyte.\n >> Currently in development. <<")
		}

	} else {
		fmt.Println("Cya!")
	}

	//}
}
