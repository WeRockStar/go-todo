package todo

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

type File struct {
	name string
}

type Todo struct {
	ID          int
	Description string
	Status      Status
}

func New(name string) *File {
	return &File{name: name}
}

func (f *File) ReadAll() ([]Todo, error) {
	file, err := os.Open(f.name)
	if err != nil {
		panic("cannot open file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	reader := csv.NewReader(file)
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

func (f *File) Add(todo Todo) error {
	file, err := os.OpenFile(f.name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("cannot open file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	w := csv.NewWriter(file)

	todos, err := f.ReadAll()
	if err != nil {
		return err
	}
	todo.ID = len(todos) + 1
	todo.Status = "No"
	record := []string{strconv.Itoa(todo.ID), todo.Description, string(todo.Status)}
	err = w.Write(record)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
