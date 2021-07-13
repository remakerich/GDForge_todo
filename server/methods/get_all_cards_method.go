package methods

import (
	"log"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

var all todo.AllCards

func (s *Server) GetAllCards(ctx context.Context, card *todo.Card) (*todo.AllCards, error) {

	log.Println("Client requested all cards")
	return &all, nil
}
