package service

import (
	"SanskritDictsApi/cmd/consts"
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

func NewAmaraDict(dictName string) (*Dict, error) {
	return newDict(utils.PathToAmaraDB().String(), dictName)
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

func (d *Dict) getQuery(dict_query string, amara_query string) string {
	var query = dict_query
	if d.dbName == consts.AMARA {
		query = amara_query
	}
	return query
}

const amara_select_start = "SELECT w_word as key, w_id as lnum, w_synonyms || '*' || sh.sh_text_line1 || '**' || sh.sh_text_line2 || '***' || sh.sh_number as data from words m join shlokas sh on m.w_shloka_id = sh.sh_id "
const amara_select_suggest_start = "SELECT w_word as key, w_id as lnum, '' as data from words m join shlokas sh on m.w_shloka_id = sh.sh_id "

func (d *Dict) GetSuggestions(term string, limit int) ([]KeyData, error) {
	dict_query := fmt.Sprintf("SELECT * from %s ", d.dbName) + "m where m.key like '%s%%' order by lnum LIMIT %d"
	amara_query := amara_select_suggest_start + " where m.w_word like '%s%%' order by w_id LIMIT %d"
	return d.loadData(fmt.Sprintf(d.getQuery(dict_query, amara_query), term, limit))
}

func (d *Dict) GetSuggestion(term string) ([]KeyData, error) {
	dict_query := fmt.Sprintf("SELECT * from %s ", d.dbName) + "m where m.key like '%s' order by lnum LIMIT %d"
	amara_query := amara_select_start + " where m.w_word like '%s' order by w_id LIMIT %d"
	return d.loadData(fmt.Sprintf(d.getQuery(dict_query, amara_query), term, 1))
}

func (d *Dict) GetSearchResult(term string) ([]KeyData, error) {
	dict_query := fmt.Sprintf("SELECT * from %s ", d.dbName) + "m where m.key like '%s' order by lnum"
	amara_query := amara_select_start + " where m.w_word like '%s' order by w_id"
	return d.loadData(fmt.Sprintf(d.getQuery(dict_query, amara_query), term))
}

func (d *Dict) GetResult(nums string) ([]KeyData, error) {
	dict_query := fmt.Sprintf("SELECT * from %s ", d.dbName) + "m where m.lnum in (%s)"
	amara_query := amara_select_start + " where m.w_id in (%s)"
	return d.loadData(fmt.Sprintf(d.getQuery(dict_query, amara_query), nums))
}

func (d *Dict) GetResultForNum(term string) ([]KeyData, error) {
	dict_query := fmt.Sprintf("SELECT * from %s ", d.dbName) + "m where m.key like '%s' order by lnum LIMIT 1"
	amara_query := amara_select_start + " where m.w_word like '%s' order by w_id LIMIT 1"
	return d.loadData(fmt.Sprintf(d.getQuery(dict_query, amara_query), term))
}

func (d *Dict) GetResultList(lnum float32, count float32) ([]string, error) {
	dict_query := fmt.Sprintf("SELECT m.key from %s ", d.dbName) + "m.lnum >= %f and m.lnum <= %f group by m.key order by m.lnum"
	amara_query := "SELECT w_word as key from words m where m.w_id >= %f and m.w_id <= %f group by m.w_word order by m.w_id"

	sprintf := fmt.Sprintf(d.getQuery(dict_query, amara_query), lnum-count, lnum+count)
	return d.loadString(sprintf)
}

func (d *Dict) Close() {
	if d != nil && d.db != nil {
		d.db.Close()
	}
}
