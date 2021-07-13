package methods

import (
	"log"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

func (s *Server) RemoveCard(ctx context.Context,
	card *todo.Card) (*todo.Card, error) {

	var temp todo.AllCards
	for _, item := range all.Cards {
		if item.Id != card.Id {
			temp.Cards = append(temp.Cards, item)
		}
	}
	if len(all.Cards) == len(temp.Cards) {
		log.Printf("Can't remove card (id: %d): card not found", card.Id)
		return &todo.Card{Description: "not found"}, nil
	} else {
		log.Printf("Card (id: %d) successfully removed", card.Id)
		all.Cards = temp.Cards
		return &todo.Card{}, nil
	}
}
