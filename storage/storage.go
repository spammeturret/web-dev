package storage

import (
	"encoding/json"

	"fmt"
	"os"

	"github.com/spammeturret/go-react-todo/pkg"
)

func LoadTodoFromJson(filePath string) ([]pkg.Todo, error) {
	// Read the JSON file.
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return []pkg.Todo{}, err
	}
	// Declare a variable to hold the unmarshaled data.
	var data []pkg.Todo
	// Unmarshal the JSON data into the MyData struct.
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return []pkg.Todo{}, err
	}
	//DEBUG
	// fmt.Printf("fileContent value: %v\n", string(fileContent))
	// fmt.Printf("data value: %v\n", data)
	// fmt.Printf("error value: %v\n", err)
	return data, nil
}

func SaveTodoToFile(filePath string, Todos []pkg.Todo) ([]pkg.Todo, error) {
	// var data []pkg.Todo
	// data, err := LoadTodoFromJson(filePath)
	// if err != nil {
	// 	return []pkg.Todo{}, err
	// }
	// data = append(data, newTodo)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []pkg.Todo{}, err
	}
	defer file.Close()
	jsonData, err := json.Marshal(Todos)
	if err != nil {
		return []pkg.Todo{}, err
	}
	_, err = file.Write([]byte(jsonData))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return []pkg.Todo{}, err
	}

	fmt.Println("File content overwritten successfully.")
	return Todos, err
}
