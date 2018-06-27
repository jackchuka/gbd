package api

import (
	"encoding/json"
	"log"
	"strings"
)

// ParseWords will return weight JSON
func ParseWords(text string) []byte {
	words := strings.Split(strings.ToLower(text), " ")

	m := map[string]int{}
	for _, word := range words {
		m[word] = m[word] + 1
	}

	json, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return json
}
