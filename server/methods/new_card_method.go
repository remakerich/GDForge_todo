package methods

import (
	"log"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

var count = 1

func (s *Server) NewCard(ctx context.Context,
	card *todo.Card) (*todo.Card, error) {

	newCard := todo.Card{
		Id:          int32(count),
		Description: card.Description,
		Done:        false}
	count++

	all.Cards = append(all.Cards, &newCard)

	log.Printf("New card with id %d generated: '%s'",
		newCard.Id, newCard.Description)
	return &newCard, nil
}
