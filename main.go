package main

import (
	"fmt"
	todo "github.com/werockstar/go-todo/todo"
	"os"
)

var (
	bannerText = `
___________        .___      
\__    ___/___   __| _/____  
  |    | /  _ \ / __ |/  _ \ 
  |    |(  <_> ) /_/ (  <_> )
  |____| \____/\____ |\____/ 
                    \/       
`[1:]
	optionsText = `
Usage options:
1. View All Todo
2. Add Todo
3. Mark as Done
4. View Done Only
5. View In-progress Only
0. Exit
`[1:]
)

func main() {
	t := todo.New("todos.csv")
	records, err := t.ReadAll()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	var menu int
	fmt.Println(bannerText)
	fmt.Print(optionsText)
	for {
		fmt.Print("Enter option: ")
		_, err := fmt.Scanf("%d\n", &menu)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(1)
		}

		switch menu {
		case 1:
			for _, t := range records {
				fmt.Printf("%d. %s %s\n", t.ID, t.Description, t.Status)
			}
			break
		case 2:
			var description string
			fmt.Print("Todo: ")
			_, _ = fmt.Scanf("%s", &description)
			err := t.Add(todo.Todo{Description: description})
			if err != nil {
				continue
			}
			break
		case 0:
			os.Exit(0)
		}
	}
}
