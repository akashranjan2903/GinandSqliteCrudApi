package controllers

import (
	"encoding/json"
	"fmt"
	"os"
)

func (b bloglist) SavetoJson() {
	// create files
	data, err := json.Marshal(b.blogStore)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("db.json", data, 0o644); err != nil {
		panic(err)
	}
}

func (b bloglist) AddnewId() int {
	if len(b.blogStore) == 0 {
		return 1
	}
	return b.blogStore[len(b.blogStore)-1].Id + 1
}
func (b *bloglist) LoadFromJson() {
	// Check if the file exists
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		os.Create("db.json")
	}

	// convert the file to a byte array
	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		// Unmarshal the data from the file to t.todoStore
		err = json.Unmarshal(data, &b.blogStore)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No data found")
	}
}
