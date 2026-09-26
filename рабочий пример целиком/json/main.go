package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	t := Task{ID: 1, Title: "Сверстать карточку", Done: false}
	out, err := json.Marshal(t)
	if err != nil {
		fmt.Println("ошибка:", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
