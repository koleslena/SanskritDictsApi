package utils

import "strings"

const pathToDB = "./data/"
const keysFile string = "keys.sqlite"
const dictFile string = ".sqlite"

func PathToSuggestions(dictName string) string {
	sb := PathToDB(dictName)
	sb.WriteString(keysFile)
	return sb.String()
}

func PathToSearch(dictName string) string {
	sb := PathToDB(dictName)
	sb.WriteString(dictFile)
	return sb.String()
}

func PathToDB(dictName string) *strings.Builder {
	var sb = strings.Builder{}
	sb.WriteString(pathToDB)
	sb.WriteString(dictName)
	return &sb
}
