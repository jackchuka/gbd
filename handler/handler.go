package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jackchuka/gbd/configs"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Configs configs.Configs
}

type queryRequest struct {
	Query    string `json:"query"`
	TypeWord string `json:"type_word,omitempty"`
}

func (h *Handler) APIHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 16384))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	var req queryRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		log.Println(string(body))
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	log.Println("Query started...")
	rows, err := h.getStringFromQuery(req.Query)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(ParseWords(rows, req))
}

func (h *Handler) getStringFromQuery(query string) (sqlx.Rows, error) {
	db, err := sqlx.Connect(h.Configs.Driver, h.Configs.DSN)
	if err != nil {
		return sqlx.Rows{}, err
	}

	rows, err := db.Queryx(query)
	if err != nil {
		return sqlx.Rows{}, err
	}
	return *rows, nil
}
