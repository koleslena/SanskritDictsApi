package service

import (
	"SanskritDictsApi/cmd/consts"
	"log"
	"strings"
	"unicode"
)

type Transliteration struct {
	from, to string
}

func NewTransliteration(from, to string) *Transliteration {
	log.Println("new Transliteration")
	return &Transliteration{from, to}
}

// TODO: sāyā°hna-samaye sāyāhan, n. or sāyā°hna, <s>sAyA<srs/>°hna-samaye</s> <s>sAyA<srs/>°hna</s>
func (ts *Transliteration) Transliterate(str string) string {
	str = strings.ReplaceAll(str, "/", "") // TODO clean DB
	switch ts.from + ts.to {
	case consts.DEVANAGARY + consts.SLP1:
		return devaToSlp(str)
	case consts.DEVANAGARY + consts.HK:
		return devaToHk(str)
	case consts.DEVANAGARY + consts.IAST:
		return devaToIAST(str)
	case consts.IAST + consts.DEVANAGARY:
		return IASTToDeva(str)
	case consts.SLP1 + consts.DEVANAGARY:
		return SlpToDeva(str)
	case consts.IAST + consts.HK:
		return IASTToHK(str)
	case consts.HK + consts.DEVANAGARY:
		return HKToDeva(str)
	case consts.HK + consts.IAST:
		return HKToIAST(str)
	case consts.SLP1 + consts.IAST:
		return SlpToIAST(str)
	case consts.SLP1 + consts.HK:
		return SlpToHK(str)
	case consts.HK + consts.SLP1:
		return HKToSlp(str)
	case consts.IAST + consts.SLP1:
		return IASTToSlp(str)
	default:
		return str
	}
}

func SlpToDeva(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	isPrevConsonant := false
	for _, iRune := range runes {
		if devaRune, ok := consts.SlpToDeva[string(iRune)]; ok {
			isPrevConsonant = SlpToDevaForString(&replacement, devaRune, string(iRune), isPrevConsonant)
		} else {
			replacement.WriteString(string(iRune))
			isPrevConsonant = false
		}
	}

	if isPrevConsonant && string(runes[len(runes)-1]) != consts.DEFAULT_VOWEL {
		replacement.WriteString(string(consts.VIRAMA))
	}
	return replacement.String()
}

func SlpToDevaForString(replacement *strings.Builder, devaRune rune, ch string, isPrevConsonant bool) bool {
	tmp := contains(consts.DevaConsonants, devaRune)
	if isPrevConsonant {
		if tmp {
			replacement.WriteString(string(consts.VIRAMA))
			replacement.WriteString(string(devaRune))
			return true
		} else if containsStrings(consts.SlpVowels, ch) {
			if ch != consts.DEFAULT_VOWEL {
				replacement.WriteString(string(consts.SlpVowelMarks[ch]))
			}
			return false
		} else {
			if containsStrings(consts.PUNCTUATIONS, ch) {
				replacement.WriteString(string(consts.VIRAMA))
			}
			if contains(consts.DevaDandas, devaRune) {
				replacement.WriteString(string(consts.VIRAMA))
			}
			replacement.WriteString(string(devaRune))
			return false
		}
	} else {
		replacement.WriteString(string(devaRune))
		return tmp
	}
}

func HKToDeva(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	isPrevConsonant := false
	doubleString := ""
	for _, iRune := range runes {
		if doubleString == "" {
			if isDoubleString(string(iRune)) {
				doubleString = string(iRune)
				continue
			}
		} else {
			if string(iRune) == consts.H {
				if containsStrings(consts.WithHStrings, doubleString) {
					if isPrevConsonant {
						replacement.WriteString(string(consts.VIRAMA))
					}
					replacement.WriteString(string(consts.HKToDeva[string(doubleString)+consts.H]))
					isPrevConsonant = true
					doubleString = ""
					continue
				}
			} else if containsStrings(consts.WithAStrings, string(iRune)) {
				if consts.DEFAULT_VOWEL == doubleString || consts.VOWEL_RR == doubleString {
					if isPrevConsonant {
						replacement.WriteString(string(consts.HKVowelMarks[doubleString+string(iRune)]))
					} else {
						replacement.WriteString(string(consts.HKToDeva[doubleString+string(iRune)]))
					}
					isPrevConsonant = false
					doubleString = ""
					continue
				}
			}
			//
			if devaRune, ok := consts.HKToDeva[doubleString]; ok {
				isPrevConsonant = HKToDevaForString(&replacement, devaRune, doubleString, isPrevConsonant)
			} else {
				replacement.WriteString(doubleString)
				isPrevConsonant = false
			}
			if isDoubleString(string(iRune)) {
				doubleString = string(iRune)
				continue
			}
			doubleString = ""
		}
		if devaRune, ok := consts.HKToDeva[string(iRune)]; ok {
			isPrevConsonant = HKToDevaForString(&replacement, devaRune, string(iRune), isPrevConsonant)
		} else {
			replacement.WriteString(string(iRune))
			isPrevConsonant = false
		}
	}

	if isPrevConsonant && string(doubleString) != consts.DEFAULT_VOWEL {
		replacement.WriteString(string(consts.VIRAMA))
	}
	return replacement.String()
}

func isDoubleString(str string) bool {
	if containsStrings(consts.WithHStrings, str) || consts.DEFAULT_VOWEL == str || consts.VOWEL_RR == str {
		return true
	}
	return false
}

func HKToDevaForString(replacement *strings.Builder, devaRune rune, ch string, isPrevConsonant bool) bool {
	tmp := contains(consts.DevaConsonants, devaRune)
	if isPrevConsonant {
		if tmp {
			replacement.WriteString(string(consts.VIRAMA))
			replacement.WriteString(string(devaRune))
			return true
		} else if containsStrings(consts.HKVowels, ch) {
			if ch != consts.DEFAULT_VOWEL {
				replacement.WriteString(string(consts.HKVowelMarks[ch]))
			}
			return false
		} else {
			if containsStrings(consts.PUNCTUATIONS, ch) {
				replacement.WriteString(string(consts.VIRAMA))
			}
			if contains(consts.DevaDandas, devaRune) {
				replacement.WriteString(string(consts.VIRAMA))
			}
			replacement.WriteString(string(devaRune))
			return false
		}
	} else {
		replacement.WriteString(string(devaRune))
		return tmp
	}
}

func SlpToHK(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	for _, iRune := range runes {
		if iastString, ok := consts.SlpToHKString[string(iRune)]; ok {
			replacement.WriteString(iastString)
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func HKToSlp(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	skip := false
	for i, iRune := range runes {
		if skip {
			skip = false
			continue
		}
		if i+1 < len(runes) {
			if slpDoubleString, ok := consts.HKToSlpString[string(iRune)+string(runes[i+1])]; ok {
				replacement.WriteString(slpDoubleString)
				skip = true
			} else if slpString, ok := consts.HKToSlpString[string(iRune)]; ok {
				replacement.WriteString(slpString)
			} else {
				replacement.WriteString(string(iRune))
			}
		} else if slpString, ok := consts.HKToSlpString[string(iRune)]; ok {
			replacement.WriteString(slpString)
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func IASTToSlp(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	skip := false
	for i, iRune := range runes {
		if skip {
			skip = false
			continue
		}
		if i+1 < len(runes) {
			if slpDoubleString, ok := consts.IASTToSlpString[string(iRune)+string(runes[i+1])]; ok {
				replacement.WriteString(slpDoubleString)
				skip = true
			} else if slpString, ok := consts.IASTToSlp[iRune]; ok {
				replacement.WriteString(slpString)
			} else if slpString, ok := consts.IASTToSlpString[string(iRune)]; ok {
				replacement.WriteString(slpString)
			} else {
				replacement.WriteString(string(iRune))
			}
		} else if slpString, ok := consts.IASTToSlp[iRune]; ok {
			replacement.WriteString(slpString)
		} else if slpString, ok := consts.IASTToSlpString[string(iRune)]; ok {
			replacement.WriteString(slpString)
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func SlpToIAST(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	for _, iRune := range runes {
		if iastRune, ok := consts.SlpToIAST[string(iRune)]; ok {
			replacement.WriteString(string(iastRune))
		} else if iastString, ok := consts.SlpToIASTString[string(iRune)]; ok {
			replacement.WriteString(iastString)
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func HKToIAST(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	isPrevL := false
	for _, iRune := range runes {
		if !isPrevL {
			if string(iRune) == "l" {
				isPrevL = true
				continue
			}
		} else {
			if string(iRune) == "R" {
				replacement.WriteString(string(consts.HKToIAST["lR"]))
				isPrevL = false
				continue
			}
			replacement.WriteString("l")
			isPrevL = false
		}
		if iastRune, ok := consts.HKToIAST[string(iRune)]; ok {
			replacement.WriteString(string(iastRune))
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func IASTToHK(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	for _, iRune := range runes {
		if hkRune, ok := consts.IASTToHK[iRune]; ok {
			replacement.WriteString(hkRune)
		} else {
			replacement.WriteString(string(iRune))
		}
	}
	return replacement.String()
}

func IASTToDeva(str string) string {
	var replacement strings.Builder
	runes := []rune(str)
	isPrevConsonant := false
	doubleRune := rune(0)
	for _, iRune := range runes {
		if doubleRune == rune(0) {
			if isDoubleRune(iRune) {
				doubleRune = iRune
				continue
			}
		} else {
			if string(iRune) == consts.H {
				if withH, ok := consts.WithH[doubleRune]; ok {
					if isPrevConsonant {
						replacement.WriteString(string(consts.VIRAMA))
					}
					replacement.WriteString(string(withH))
					isPrevConsonant = true
					doubleRune = rune(0)
					continue
				}
				if containsStrings(consts.WithHStrings, string(doubleRune)) {
					if isPrevConsonant {
						replacement.WriteString(string(consts.VIRAMA))
					}
					replacement.WriteString(string(consts.IASTToDevaString[string(doubleRune)+consts.H]))
					isPrevConsonant = true
					doubleRune = rune(0)
					continue
				}
			} else if containsStrings(consts.WithAStrings, string(iRune)) {
				if consts.DEFAULT_VOWEL_DEVA_RUNE == doubleRune {
					if isPrevConsonant {
						replacement.WriteString(string(consts.DevaVowelsMarksFromString[string(doubleRune)+string(iRune)]))
					} else {
						replacement.WriteString(string(consts.IASTToDevaString[string(doubleRune)+string(iRune)]))
					}
					isPrevConsonant = false
					doubleRune = rune(0)
					continue
				}
			}
			//
			if doubleRune < unicode.MaxASCII {
				devaRune := consts.IASTToDevaString[string(doubleRune)]
				isPrevConsonant = IASTToDevaForRune(&replacement, devaRune, doubleRune, isPrevConsonant)
			} else if devaRune, ok := consts.IASTToDeva[doubleRune]; ok {
				isPrevConsonant = IASTToDevaForRune(&replacement, devaRune, doubleRune, isPrevConsonant)
			}
			if isDoubleRune(iRune) {
				doubleRune = iRune
				continue
			}
			doubleRune = rune(0)
		}
		if iRune < unicode.MaxASCII {
			if devaRune, ok := consts.IASTToDevaString[string(iRune)]; ok {
				isPrevConsonant = IASTToDevaForRune(&replacement, devaRune, iRune, isPrevConsonant)
			} else {
				replacement.WriteString(string(iRune))
				isPrevConsonant = false
			}
		} else if devaRune, ok := consts.IASTToDeva[iRune]; ok {
			isPrevConsonant = IASTToDevaForRune(&replacement, devaRune, iRune, isPrevConsonant)
		}
	}

	if isPrevConsonant && string(doubleRune) != consts.DEFAULT_VOWEL {
		replacement.WriteString(string(consts.VIRAMA))
	}
	return replacement.String()
}

func isDoubleRune(iRune rune) bool {
	if _, ok := consts.WithH[iRune]; ok {
		return true
	}
	if containsStrings(consts.WithHStrings, string(iRune)) || consts.DEFAULT_VOWEL == string(iRune) {
		return true
	}
	return false
}

func IASTToDevaForRune(replacement *strings.Builder, devaRune rune, iRune rune, isPrevConsonant bool) bool {
	tmp := contains(consts.DevaConsonants, devaRune)
	if isPrevConsonant {
		if tmp {
			replacement.WriteString(string(consts.VIRAMA))
			replacement.WriteString(string(devaRune))
			return true
		} else if containsStrings(consts.IASTVowels, string(iRune)) {
			if string(iRune) != consts.DEFAULT_VOWEL {
				replacement.WriteString(string(consts.DevaVowelsMarksFromString[string(iRune)]))
			}
			return false
		} else if contains(consts.IASTVowelsRunes, iRune) {
			replacement.WriteString(string(consts.DevaVowelsMarksFromRunes[devaRune]))
			return false
		} else {
			if _, ok := consts.PUNCTUATION[iRune]; ok {
				replacement.WriteString(string(consts.VIRAMA))
			}
			if contains(consts.DevaDandas, devaRune) {
				replacement.WriteString(string(consts.VIRAMA))
			}
			replacement.WriteString(string(devaRune))
			return false
		}
	} else {
		replacement.WriteString(string(devaRune))
		return tmp
	}
}

func devaToIAST(str string) string {
	return fromDeva(str, consts.DevaToIAST)
}

func devaToHk(str string) string {
	return fromDeva(str, consts.DevaToHK)
}

func devaToSlp(str string) string {
	return fromDeva(str, consts.DevaToSlp)
}

func fromDeva(str string, arrMap map[rune]string) string {
	var replacement strings.Builder
	runes := []rune(str)
	isPrevConsonant := false
	for _, devaRune := range runes {
		if value, ok := arrMap[devaRune]; ok {
			tmp := contains(consts.DevaConsonants, devaRune)
			if isPrevConsonant && (tmp || contains(consts.DevaEndings, devaRune)) {
				replacement.WriteString(consts.DEFAULT_VOWEL)
			}
			replacement.WriteString(value)
			isPrevConsonant = tmp
			continue
		}
		if value, ok := consts.PUNCTUATION[devaRune]; ok {
			if isPrevConsonant {
				replacement.WriteString(consts.DEFAULT_VOWEL)
			}
			replacement.WriteString(value)
			isPrevConsonant = false
		} else if devaRune < unicode.MaxASCII {
			replacement.WriteString(string(devaRune))
			continue
		}
	}
	if isPrevConsonant {
		replacement.WriteString(consts.DEFAULT_VOWEL)
	}
	return replacement.String()
}

func contains(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func containsStrings(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
