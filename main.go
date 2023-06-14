package main

import (
	"fmt"
	todo "github.com/werockstar/go-todo/todo"
	"os"
)

func main() {
	t := todo.New("todos.csv")
	records, err := t.ReadAll()
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
			for _, t := range records {
				fmt.Printf("%d. %s %s\n", t.ID, t.Description, t.Status)
			}
			fmt.Println("================= End ==================")
		case 2:
			var description string
			fmt.Print("Enter Todo Description: ")
			_, _ = fmt.Scanf("%s\n", &description)
			err := t.Add(todo.Todo{ID: 3, Description: description, Status: todo.Status("No")})
			if err != nil {
				return
			}
		case 0:
			os.Exit(0)
		}
	}
}
