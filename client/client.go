package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/GDForge/client/calls"
	"example.com/GDForge/todo"
	"google.golang.org/grpc"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer connection.Close()
	client := todo.NewTodoServiceClient(connection)

	calls.Greeting()
	calls.ListOfCommands()

commands:
	for {
		fmt.Print("\n> ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		switch strings.TrimSpace(command) {
		case "new":
			calls.NewCard(client)
		case "all":
			calls.GetAllCards(client)
		case "rem":
			calls.RemoveCard(client)
		case "done":
			calls.MarkCardAsDone(client)
		case "help":
			calls.ListOfCommands()
		case "quit":
			fmt.Println("Client terminated")
			break commands
		default:
			fmt.Println("Please enter a valid command")
		}
	}
}
