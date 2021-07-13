package calls

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

func NewCard(client todo.TodoServiceClient) {

	fmt.Print("Card description: ")
	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	newDescription := todo.Card{
		Description: strings.TrimSpace(userInput),
	}

	newCard, error := client.NewCard(context.Background(), &newDescription)
	if error != nil {
		log.Fatalf("Error when calling NewCard: %v", error)
	}

	log.Printf("\nCard generated:\n(id: %d) '%s' (done: %t)",
		newCard.Id, newCard.Description, newCard.Done)
}
