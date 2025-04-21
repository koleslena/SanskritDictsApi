package utils

import "strings"
import "SanskritDictsApi/cmd/consts"

const pathToDB = "./data/"
const keysFile string = "keys"
const dictFile string = ".sqlite"
const amaraDB string = "amara.db"
const shabdaDB string = "shabda.db"
const shabda string = "shabda"

func PathToSuggestions(dictName string) (string, string) {
	sb := PathToDB(dictName)
	if dictName == consts.MW {
		sb.WriteString(keysFile)
		sb.WriteString(dictFile)
		return sb.String(), dictName + keysFile
	}
	sb.WriteString(dictFile)
	return sb.String(), dictName
}

func PathToShabda() (string, string) {
	var sb = strings.Builder{}
	sb.WriteString(pathToDB)
	sb.WriteString(shabdaDB)
	return sb.String(), shabda
}

func PathToSearch(dictName string) (string, string) {
	sb := PathToDB(dictName)
	sb.WriteString(dictFile)
	return sb.String(), dictName
}

func PathToAmaraDB() *strings.Builder {
	var sb = strings.Builder{}
	sb.WriteString(pathToDB)
	sb.WriteString(amaraDB)
	return &sb
}

func PathToDB(dictName string) *strings.Builder {
	var sb = strings.Builder{}
	sb.WriteString(pathToDB)
	sb.WriteString(dictName)
	return &sb
}
