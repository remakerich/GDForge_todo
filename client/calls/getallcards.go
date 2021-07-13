package calls

import (
	"fmt"
	"log"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

func GetAllCards(client todo.TodoServiceClient) {

	allCards, error := client.GetAllCards(context.Background(), &todo.Card{})
	if error != nil {
		log.Fatalf("Error when calling GetAllCards: %v", error)
	}

	log.Println("\nAll cards:")
	for _, card := range allCards.Cards {
		fmt.Printf("(id: %d) '%s' (done: %t)\n",
			card.Id, card.Description, card.Done)
	}
}
