package web

import (
	"SanskritDictsApi/cmd/consts"
	"SanskritDictsApi/cmd/service"
	"SanskritDictsApi/utils"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Title struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

const maxSuggestions = 10

var dicts = service.NewDicts()

var tpl = template.Must(template.ParseFiles("./ui/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func TransliterateHandler(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	output := r.URL.Query().Get("output")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if len(term) != 0 && len(output) > 0 {
		result := service.NewTransliteration(consts.SLP1, output).Transliterate(term)
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(term)
	}
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	dict := r.URL.Query().Get("dict")
	if len(dict) == 0 {
		dict = "mw"
	}
	dictSet, err := dicts.GetDict(dict)
	if err != nil {
		log.Println("Error initiation of dict ", err)
		json.NewEncoder(w).Encode([]string{})
	} else {
		term := r.URL.Query().Get("term")
		input := r.URL.Query().Get("input")
		log.Println("term =>", term)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		if len(term) != 0 {
			if len(input) != 0 && input != consts.SLP1 {
				term = service.NewTransliteration(input, consts.SLP1).Transliterate(term)
				suggestions, err := dictSet.DictSuggestions.GetSuggestions(term, maxSuggestions)
				var result []Title
				transliteration := service.NewTransliteration(consts.SLP1, input)
				for _, x := range suggestions {
					result = append(result, Title{Name: transliteration.Transliterate(x.Key), Value: x.Key})
				}
				if err != nil {
					log.Println("Error load suggestions ", err)
				} else {
					json.NewEncoder(w).Encode(result)
				}
			} else {
				suggestions, err := dictSet.DictSuggestions.GetSuggestions(term, maxSuggestions)
				var result []Title
				for _, x := range suggestions {
					result = append(result, Title{Name: x.Key, Value: x.Key})
				}
				if err != nil {
					log.Println("Error load suggestions ", err)
				} else {
					json.NewEncoder(w).Encode(result)
				}
			}
		} else {
			json.NewEncoder(w).Encode([]string{})
		}
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var keyData service.KeyData
	err := decoder.Decode(&keyData)
	if err != nil {
		log.Println("Error get body ", err)
	}
	dict := r.URL.Query().Get("dict")
	if len(dict) == 0 {
		dict = "mw"
	}
	dictSet, err := dicts.GetDict(dict)
	if err != nil {
		log.Println("Error initiation of dict ", err)
		json.NewEncoder(w).Encode([]string{})
	} else {
		term := r.URL.Query().Get("term")
		nums := keyData.Data
		log.Println("search term =>", term)
		log.Println("search nums =>", nums)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		if len(nums) != 0 {
			result, err := dictSet.DictSearch.GetResult(utils.CleanData(nums))
			if err != nil {
				log.Println("Error load result ", err)
			} else {
				json.NewEncoder(w).Encode(result)
			}
		} else if len(term) != 0 {
			if strings.ToLower(dict) != consts.MW {
				result, err := dictSet.DictSuggestions.GetSearchResult(term)
				if err != nil {
					log.Println("Error load result ", err)
				} else {
					json.NewEncoder(w).Encode(result)
				}
			} else {
				suggestion, err := dictSet.DictSuggestions.GetSuggestion(term)
				if err != nil {
					log.Println("Error load suggestion ", err)
				} else if len(suggestion) == 1 {
					result, err := dictSet.DictSearch.GetResult(utils.CleanData(suggestion[0].Data))
					if err != nil {
						log.Println("Error load result ", err)
					} else {
						json.NewEncoder(w).Encode(result)
					}
				}
			}
		} else {
			json.NewEncoder(w).Encode([]service.KeyData{})
		}
	}
}

func SearchListHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var keyData service.KeyData
	err := decoder.Decode(&keyData)
	if err != nil {
		log.Println("Error get body ", err)
	}
	dict := r.URL.Query().Get("dict")
	if len(dict) == 0 {
		dict = "mw"
	}
	dictSet, err := dicts.GetDict(dict)
	if err != nil {
		log.Println("Error initiation of dict ", err)
		json.NewEncoder(w).Encode([]string{})
	} else {
		term := r.URL.Query().Get("term")
		nums := keyData.Data
		log.Println("list term =>", term)
		log.Println("list nums =>", nums)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		if len(nums) != 0 {
			result, err := dictSet.DictSearch.GetResult(utils.CleanData(nums))
			if err != nil {
				log.Println("Error load result ", err)
			} else {
				json.NewEncoder(w).Encode(result)
			}
		} else if len(term) != 0 {
			result, err := dictSet.DictSearch.GetResultForNum(term)
			if err != nil {
				log.Println("Error load result ", err)
			} else if len(result) == 1 {
				list, err := dictSet.DictSearch.GetResultList(result[0].Lnum, 15)
				if err != nil {
					log.Println("Error load result ", err)
				} else {
					json.NewEncoder(w).Encode(list)
				}
			}
		} else {
			json.NewEncoder(w).Encode([]string{})
		}
	}
}

func CloseAll() {
	dicts.CloseAll()
}
