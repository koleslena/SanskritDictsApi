package service

import (
	"SanskritDictsApi/utils"
	"database/sql"
	"fmt"
	"log"

	"github.com/mattn/go-sqlite3"
)

type KeyData struct {
	Key  string  `json:"key"`
	Lnum float32 `json:"lnum"`
	Data string  `json:"data"`
}

type Dict struct {
	dbName string
	db     *sql.DB
	dd     sqlite3.SQLiteDriver
}

func NewDictSuggestions(dictName string) (*Dict, error) {
	return newDict(utils.PathToSuggestions(dictName))
}

func NewDict(dictName string) (*Dict, error) {
	return newDict(utils.PathToSearch(dictName))
}

func newDict(pathDictName string, dictName string) (*Dict, error) {
	db, err := sql.Open("sqlite3", pathDictName)
	if err != nil {
		return nil, err
	}
	db.Exec("PRAGMA case_sensitive_like = true")
	return &Dict{
		dbName: dictName,
		db:     db,
	}, nil
}

func (d *Dict) loadData(query string) ([]KeyData, error) {
	d.db.Exec("PRAGMA case_sensitive_like = true")
	rows, err := d.db.Query(query)
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	data := make([]KeyData, 0)
	for rows.Next() {
		var key = KeyData{}
		err = rows.Scan(&key.Key, &key.Lnum, &key.Data)
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

func (d *Dict) loadString(query string) ([]string, error) {
	d.db.Exec("PRAGMA case_sensitive_like = true")
	rows, err := d.db.Query(query)
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	data := make([]string, 0)
	for rows.Next() {
		var key string
		err = rows.Scan(&key)
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

func (d *Dict) GetSuggestions(term string, limit int) ([]KeyData, error) {
	return d.loadData(fmt.Sprintf("SELECT * from %s m where m.key like '%s%%' order by lnum LIMIT %d", d.dbName, term, limit))
}

func (d *Dict) GetSuggestion(term string) ([]KeyData, error) {
	return d.loadData(fmt.Sprintf("SELECT * from %s m where m.key like '%s' order by lnum LIMIT %d", d.dbName, term, 1))
}

func (d *Dict) GetResult(nums string) ([]KeyData, error) {
	return d.loadData(fmt.Sprintf("SELECT * from %s m where m.lnum in (%s)", d.dbName, nums))
}

func (d *Dict) GetResultForNum(term string) ([]KeyData, error) {
	return d.loadData(fmt.Sprintf("SELECT * from %s m where m.key like '%s' order by lnum LIMIT 1", d.dbName, term))
}

func (d *Dict) GetResultList(lnum float32, count float32) ([]string, error) {
	sprintf := fmt.Sprintf("SELECT m.key from %s m where m.lnum >= %f and m.lnum <= %f group by m.key order by m.lnum ", d.dbName, lnum-count, lnum+count)
	return d.loadString(sprintf)
}

func (d *Dict) Close() {
	if d != nil && d.db != nil {
		d.db.Close()
	}
}
