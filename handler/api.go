package handler

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
	"github.com/jmoiron/sqlx"
)

var commonWords = map[string]int{
	"new":  1,
	"old":  1,
	"size": 1,
}

// ParseWords will return weight JSON
func ParseWords(rows sqlx.Rows, req queryRequest) []byte {
	log.Println(req)
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
				buffer.WriteString(v + " ")
			}
		}
	}

	if req.TypeWord {
		words = []string{}
		tokens := tokenize.NewTreebankWordTokenizer().Tokenize(buffer.String())

		tagger := tag.NewPerceptronTagger()
		for _, tok := range tagger.Tag(tokens) {
			if tok.Tag == "NNP" {
				words = append(words, tok.Text)
			}
		}
	}

	m := map[string]int{}
	for _, word := range words {
		// filter common words
		if _, ok := commonWords[strings.ToLower(word)]; ok {
			continue
		}
		if len(word) < 3 {
			continue
		}
		m[word] = m[word] + 1
	}

	if req.Min > 0 {
		for key, value := range m {
			if value < req.Min {
				delete(m, key)
			}
		}
	}

	result, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return result
}
