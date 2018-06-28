package handler

import (
	"encoding/json"
	"log"
	"strings"
	"github.com/jmoiron/sqlx"
)

var commonWords = map[string]int{
	"a":    1,
	"an":   1,
	"and":  1,
	"are":  1,
	"as":   1,
	"at":   1,
	"be":   1,
	"for":  1,
	"from": 1,
	"has":  1,
	"he":   1,
	"in":   1,
	"is":   1,
	"it":   1,
	"its":  1,
	"of":   1,
	"on":   1,
	"that": 1,
	"the":  1,
	"to":   1,
	"was":  1,
	"were": 1,
	"will": 1,
	"with": 1,
	"#":    1,
	"-":    1,
	"_":    1,
	".":    1,
	",":    1,
	":":    1,
	"":     1,
}

// ParseWords will return weight JSON
func ParseWords(rows sqlx.Rows, req queryRequest) []byte {
	var words []string
	var buffer strings.Builder
	// Scan hack
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				v := string(b)
				words = append(words, v)
				buffer.WriteString(v)
				buffer.WriteString(" ")
			}
		}
	}

	words = strings.Split(strings.ToLower(buffer.String()), " ")

	m := map[string]int{}
	for _, word := range words {
		// filter common words
		if _, ok := commonWords[word]; ok {
			continue
		}
		m[word] = m[word] + 1
	}

	result, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return result
}
