package utils

import (
	"testing"
)

type httpTestData struct {
	str    string
	result string
}

var (
	jsonTest = []httpTestData{
		{
			str:    "<H1><h><key1>jal</key1><key2>jal</key2></h><body><s>jal</s>   <ab>cl.</ab> 1. <s>°lati</s> (<ab>pf.</ab> <s>jajAla</s>, <ls>Pāṇ. viii, 4, 54</ls>, <ab>Sch.</ab>), ‘to be rich’ or ‘to cover’ (derived <ab>fr.</ab> <s>jAla</s>?), <ls>Dhātup. xx, 3</ls>; <div n=\"to\"/>to be sharp, <ls>ib.</ls>; <div n=\"to\"/>to be stiff or dull (for <s>jaq</s>, derived <ab>fr.</ab> <s>jaqa</s>), <ls>ib.</ls> : <ab>cl.</ab> 10. <s>jAlayati</s>, to cover, <ls n=\"Dhātup.\">xxxii, 10</ls>.<info westergaard=\"jala,20.3,01.0561\"/><info verb=\"root\" cp=\"1,10\"/></body><tail><L>77760</L><pc>414,2</pc></tail></H1>",
			result: "79144",
		},
	}
)

func TestParseXMLToJSON(t *testing.T) {
	for caseNum, data := range jsonTest {
		jsonData := XmlToJson(data.str)
		if data.result != jsonData {
			t.Errorf("validation case #%d failed, expected %s but got %s", caseNum+1, data.result, jsonData)
		}
	}
}
