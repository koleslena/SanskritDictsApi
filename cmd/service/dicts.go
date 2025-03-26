package service

import (
	"log"
	"strings"
)

type Dicts struct {
	dictMap map[string]*DictSet
}

type DictSet struct {
	DictSuggestions *Dict
	DictSearch      *Dict
}

func NewDicts() *Dicts {
	log.Println("new ")
	return &Dicts{make(map[string]*DictSet)}
}

func (ds *Dicts) GetDict(dictName string) (*DictSet, error) {
	dictName = strings.ToLower(strings.TrimSpace(dictName))
	if dict, inMap := ds.dictMap[dictName]; inMap {
		if dict.DictSuggestions != nil && dict.DictSearch != nil {
			return dict, nil
		}
	}
	dict, err := createDict(dictName)
	if err != nil {
		return nil, err
	}
	ds.dictMap[dictName] = dict
	return dict, nil
}

func createDict(dictName string) (*DictSet, error) {
	var dictSuggestions, initDictSuggestErr = NewDictSuggestions(dictName)
	if initDictSuggestErr != nil {
		return &DictSet{}, initDictSuggestErr
	}
	var dictSearch, initDictSearchErr = NewDict(dictName)
	if initDictSearchErr != nil {
		return &DictSet{}, initDictSearchErr
	}
	return &DictSet{dictSuggestions, dictSearch}, nil
}

func (ds *Dicts) CloseAll() {
	for _, dict := range ds.dictMap {
		if dict.DictSuggestions != nil {
			dict.DictSuggestions.Close()
		}
		if dict.DictSearch != nil {
			dict.DictSearch.Close()
		}
	}
}
