package calls

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"example.com/GDForge/todo"
	"golang.org/x/net/context"
)

func RemoveCard(client todo.TodoServiceClient) {

	fmt.Print("Card id: ")
	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	temp, _ := strconv.ParseInt(strings.TrimSpace(userInput), 0, 64)
	IDtoRemove := int32(temp)

	cardToRemove := todo.Card{
		Id: IDtoRemove,
	}
	Status, error := client.RemoveCard(context.Background(), &cardToRemove)
	if error != nil {
		log.Fatalf("Error when calling RemoveCard method: %v", error)
	}
	if Status.Description == "not found" {
		log.Println()
		fmt.Println("Card not found")
	} else {
		log.Println()
		fmt.Println("Card removed")
	}
}
