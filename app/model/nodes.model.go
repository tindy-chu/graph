package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type TNode []string
type TNodes map[string]TNode

func (nodes *TNodes) Get() {
	file, err := os.Open("graph.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	decode := *json.NewDecoder(file)
	if err := decode.Decode(&nodes); err != nil {
		fmt.Println(err)
	}
}
