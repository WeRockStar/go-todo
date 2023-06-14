package main

import (
	"fmt"
	"github.com/werockstar/go-todo/reader"
	"os"
)

func main() {
	todos, err := reader.ReadAll("todos.csv")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	var menu int
	fmt.Println("Welcome to Todo App")
	for {
		fmt.Println("Enter: 1 View All Todo")
		fmt.Println("Enter: 2 Add Todo")
		fmt.Println("Enter: 3 Mark as Done")
		fmt.Println("Enter: 4 View Done Only")
		fmt.Println("Enter: 5 View In-progress Only")
		fmt.Println("Enter: 0 Exit")
		fmt.Print("Enter: ")
		_, err := fmt.Scanf("%d\n", &menu)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(1)
		}

		switch menu {
		case 1:
			fmt.Println("=============== Todo List ===============")
			for _, todo := range todos {
				fmt.Printf("%d. %s %s\n", todo.ID, todo.Description, todo.Status)
			}
			fmt.Println("================= End ==================")
		case 0:
			os.Exit(0)
		}
	}
}
