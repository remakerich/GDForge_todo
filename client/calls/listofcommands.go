package calls

import (
	"fmt"
)

func ListOfCommands() {
	fmt.Println(
		`
	Available commands:

		new - create new card
		all - show all cards
		rem - remove card
		done - mark card as done

		help - show commands
		quit - terminate client`)
}
