package service

import (
	"SanskritDictsApi/utils"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type ShabdaData struct {
	Word  string `json:"word"`
	Forms string `json:"forms"`
	Linga string `json:"linga"`
}

type Shabda struct {
	tableName string
	db        *sql.DB
}

func NewShabdaService() *Shabda {
	shabda, err := newShabda(utils.PathToShabda())
	if err != nil {
		log.Println(err)
		return nil
	}
	return shabda
}

func newShabda(pathDBName string, tableName string) (*Shabda, error) {
	db, err := sql.Open("sqlite3", pathDBName)
	if err != nil {
		return nil, err
	}
	db.Exec("PRAGMA case_sensitive_like = true")
	return &Shabda{
		tableName: tableName,
		db:        db,
	}, nil
}

func (d *Shabda) loadData(query string) ([]ShabdaData, error) {
	d.db.Exec("PRAGMA case_sensitive_like = true")
	rows, err := d.db.Query(query)
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	data := make([]ShabdaData, 0)
	for rows.Next() {
		var key = ShabdaData{}
		err = rows.Scan(&key.Word, &key.Linga, &key.Forms)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		data = append(data, key)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Shabda) GetResult(term string, linga string) ([]ShabdaData, error) {
	if len(linga) != 0 {
		query := "select s.sha_word as word, s.sha_linga as linga, s.sha_forms as forms from %s s where s.sha_word like '%s' and s.sha_linga = '%s' order by sha_id LIMIT 1"
		return d.loadData(fmt.Sprintf(query, d.tableName, term, strings.ToUpper(linga)))
	} else {
		query := "select s.sha_word as word, s.sha_linga as linga, s.sha_forms as forms from %s s where s.sha_word like '%s' order by sha_id"
		return d.loadData(fmt.Sprintf(query, d.tableName, term))
	}
}

func (d *Shabda) GetResultSuggestions(term string) ([]ShabdaData, error) {
	query := "select s.sha_word as word, s.sha_linga as linga, '' as forms from %s s where s.sha_word like '%s%%' order by sha_id"
	return d.loadData(fmt.Sprintf(query, d.tableName, term))
}
