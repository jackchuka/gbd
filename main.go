package main

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/jackchuka/gbd/api"
	"github.com/jackchuka/gbd/configs"
)

func main() {
	config := configs.GetConfigs()

	log.Println(config)

	db, err := sqlx.Connect(config.Driver, config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT name FROM items ORDER BY created DESC LIMIT 40")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
				buffer.WriteString(v)
			}
		}
	}

	log.Println(string(api.ParseWords(buffer.String())))
}
