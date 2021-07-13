package methods

import (
	"log"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

func (s *Server) MarkCardAsDone(ctx context.Context,
	card *todo.Card) (*todo.Card, error) {

	// var temp todo.AllCards
	checkindex := -1
	for index, item := range all.Cards {
		if item.Id == card.Id {
			// temp.Cards = append(temp.Cards, item)
			all.Cards[index].Done = true
			checkindex = index
		}
	}
	if checkindex == -1 {
		log.Printf("Can't mark card (id: %d) as done: card not found", card.Id)
		return &todo.Card{Description: "not found"}, nil
	} else {
		log.Printf("Card (id: %d) successfully marked as done", card.Id)
		return &todo.Card{}, nil
	}
}
