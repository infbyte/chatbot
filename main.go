/* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /*
[[=================================================================]]
If you are editing this repository make sure to issue a pull request
to further the development of this chatbot.
Others would like to learn as well :)
-infbyte

Update listings:
Recent update 2/16/18 - About has been removed should maybe be in a loop
[[=================================================================]]
*/ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */ /* */

package main

import (
	"bufio"
	"fmt"
	"os"

	//. "github.com/logrusorgru/aurora"
	"github.com/fatih/color"
)

func main() {
	//var blue = color.Cyan
	//var red = color.Red

	redtype := color.New(color.FgRed, color.Bold)

	color.Cyan("Hello there. My name is Clara. I am a chatbot.")

	scanner := bufio.NewScanner(os.Stdin)

	// Outside instances
	var text string
	var name string
	var about string
	var greet string

	//	greetList := "hi" || "hey" || "hola"
	//for text != "/clara exit" { // break the loop if entered

	// Color formatting
	fmt.Printf("Enter")
	redtype.Printf("'Y'")
	fmt.Println(", to start chatting")

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

		fmt.Println("Lovely name,", name)
		fmt.Println("To learn more about me say /about, otherwise start a conversation.")

		scanner.Scan()
		greet = scanner.Text()

		if greet == "hi" || greet == "hey" || greet == "hola" {
			fmt.Println("Hey to you!")
		} else {
			if greet == "boo" || greet == "lame" {
				fmt.Println("Mean >:(")
			} else {
				fmt.Println("No greeting issued. Shall be ignored.")
			}
		}
		scanner.Scan()
		about = scanner.Text()

		if about == "/about" || about == "about" {
			fmt.Println("Thanks for asking ", name)
			fmt.Println("I am a GitHub repository by infbyte.\n >> Currently in development. <<")
		}

	} else {
		fmt.Println("Don't want to talk? That's fine. See you soon!")
	}

	//}
}
