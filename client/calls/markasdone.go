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

func MarkCardAsDone(client todo.TodoServiceClient) {

	fmt.Print("Card id: ")
	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	temp, _ := strconv.ParseInt(strings.TrimSpace(userInput), 0, 64)
	IDtoMarkAsDone := int32(temp)

	cardDone := todo.Card{
		Id: IDtoMarkAsDone,
	}
	Status, error := client.MarkCardAsDone(context.Background(), &cardDone)
	if error != nil {
		log.Fatalf("Error when calling MarkCardAsDone method: %v", error)
	}
	if Status.Description == "not found" {
		log.Println()
		fmt.Println("Card not found")
	} else {
		log.Println()
		fmt.Println("Card marked as done")
	}
}
