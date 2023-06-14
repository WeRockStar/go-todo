package reader

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Status string

const (
	Yes Status = "Yes"
	No  Status = "No"
)

type Todo struct {
	ID          int
	Description string
	Status      Status
}

func ReadAll(name string) ([]Todo, error) {
	f, err := os.Open(name)
	if err != nil {
		panic("cannot open file")
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	records = records[1:]
	if len(records) == 0 {
		return []Todo{}, nil
	}
	var todos []Todo
	for _, record := range records {
		if id, err := strconv.Atoi(record[0]); err != nil {
			return nil, err
		} else {
			description := record[1]
			status := Status(record[2])
			todos = append(todos, Todo{
				ID:          id,
				Description: description,
				Status:      status,
			})
		}
	}

	return todos, nil
}
