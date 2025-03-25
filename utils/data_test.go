package utils

import (
	"testing"
)

type validationTestData struct {
	str    string
	result string
}

var (
	validationTest = []validationTestData{
		{
			str:    "H3,79144,79144",
			result: "79144",
		}, {
			str:    "H1,79137,79142",
			result: "79137,79138,79139,79140,79141,79142",
		}, {
			str:    "H2,79098,79117",
			result: "79098,79099,79100,79101,79102,79103,79104,79105,79106,79107,79108,79109,79110,79111,79112,79113,79114,79115,79116,79117",
		}, {
			str:    "H1,79045,79045;H1,79046,79064",
			result: "79045,79046,79047,79048,79049,79050,79051,79052,79053,79054,79055,79056,79057,79058,79059,79060,79061,79062,79063,79064",
		}, {
			str:    "H1,77761,77769;H2,78261,78261",
			result: "77761,77762,77763,77764,77765,77766,77767,77768,77769,78261",
		}, {
			str:    "H2,117.11,117.11",
			result: "117.11",
		},
	}
)

func TestParseData(t *testing.T) {
	for caseNum, data := range validationTest {
		cleanData := CleanData(data.str)
		if data.result != cleanData {
			t.Errorf("validation case #%d failed, expected %s but got %s", caseNum+1, data.result, cleanData)
		}
	}
}
