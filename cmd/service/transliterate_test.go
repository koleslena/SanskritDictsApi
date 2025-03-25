package service

import (
	"SanskritDictsApi/cmd/consts"
	"fmt"
	"testing"
)

type Params struct {
	from, to string
}

type validationTestData struct {
	params Params
	str    string
	result string
}

var (
	validationTest = []validationTestData{
		{ //SLP1 to DEVANAGARY
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "yogaH",
			result: "योगः",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "san",
			result: "सन्",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "acalaSrezWa",
			result: "अचलश्रेष्ठ",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san |",
			result: "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ||4.7||",
			result: "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।४.७।।",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ||4.6||",
			result: "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
		}, {
			params: Params{from: consts.SLP1, to: consts.DEVANAGARY},
			str:    "1234567890 ||4.6||",
			result: "१२३४५६७८९० ।।४.६।।",
		}, { // DEVANAGARY to SLP1
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "अचलश्रेष्ठ",
			result: "acalaSrezWa",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
			result: "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।4.7।।",
			result: "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ..4.7..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
			result: "bahUni me vyatItAni janmAni tava cArjuna.",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
			result: "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ..4.6..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.SLP1},
			str:    "१२३४५६७८९० ।।४.६।।",
			result: "1234567890 ..4.6..",
			// DEVANAGARY to IAST
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "अचलश्रेष्ठ",
			result: "acalaśreṣṭha",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
			result: "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।4.7।।",
			result: "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ..4.7..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
			result: "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
			result: "bahūni me vyatītāni janmāni tava cārjuna.",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
			result: "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ..4.6..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.IAST},
			str:    "१२३४५६७८९० ।।४.६।।",
			result: "1234567890 ..4.6..",
			// 	DEVANAGARY to HK
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "अचलश्रेष्ठ",
			result: "acalazreSTha",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
			result: "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।4.7।।",
			result: "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ..4.7..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH .",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
			result: "bahUni me vyatItAni janmAni tava cArjuna.",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
			result: "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ..4.6..",
		}, {
			params: Params{from: consts.DEVANAGARY, to: consts.HK},
			str:    "१२३४५६७८९० ।।४.६।।",
			result: "1234567890 ..4.6..",
			// IAST to DEVANAGARY
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "yogaḥ",
			result: "योगः",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "san",
			result: "सन्",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "acalaśreṣṭha",
			result: "अचलश्रेष्ठ",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san |",
			result: "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ||4.7||",
			result: "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।४.७।।",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ |",
			result: "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "bahūni me vyatītāni janmāni tava cārjuna|",
			result: "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ||4.6||",
			result: "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
		}, {
			params: Params{from: consts.IAST, to: consts.DEVANAGARY},
			str:    "1234567890 ||4.6||",
			result: "१२३४५६७८९० ।।४.६।।",
		}, { // SLP to HK
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "acalaSrezWa",
			result: "acalazreSTha",
		}, {
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san ।",
			result: "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san ।",
		}, {
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ||4.7||",
			result: "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ||4.7||",
		}, {
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
		}, {
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "bahUni me vyatItAni janmAni tava cArjuna|",
		}, {
			params: Params{from: consts.SLP1, to: consts.HK},
			str:    "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ||4.6||",
			result: "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ||4.6||",
		}, { // HK to Slp
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "acalazreSTha",
			result: "acalaSrezWa",
		}, {
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san ।",
			result: "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san ।",
		}, {
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ||4.7||",
			result: "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ||4.7||",
		}, {
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
		}, {
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "bahUni me vyatItAni janmAni tava cArjuna|",
		}, {
			params: Params{from: consts.HK, to: consts.SLP1},
			str:    "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ||4.6||",
			result: "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ||4.6||",
		}, { // IAST to HK
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "acalaśreṣṭha",
			result: "acalazreSTha",
		}, {
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san ।",
			result: "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san ।",
		}, {
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ||4.7||",
			result: "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ||4.7||",
		}, {
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ |",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
		}, {
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "bahūni me vyatītāni janmāni tava cārjuna|",
			result: "bahUni me vyatItAni janmAni tava cArjuna|",
		}, {
			params: Params{from: consts.IAST, to: consts.HK},
			str:    "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ||4.6||",
			result: "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ||4.6||",
			// HK to IAST
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "acalazreSTha",
			result: "acalaśreṣṭha",
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san ।",
			result: "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san ।",
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ||4.7||",
			result: "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ||4.7||",
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ |",
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "bahūni me vyatītāni janmāni tava cārjuna|",
		}, {
			params: Params{from: consts.HK, to: consts.IAST},
			str:    "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ||4.6||",
			result: "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ||4.6||",
		}, { // SLP to IAST
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "acalaSrezWa",
			result: "acalaśreṣṭha",
		}, {
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san |",
			result: "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san |",
		}, {
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ||4.7||",
			result: "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ||4.7||",
		}, {
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ |",
		}, {
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "bahūni me vyatītāni janmāni tava cārjuna|",
		}, {
			params: Params{from: consts.SLP1, to: consts.IAST},
			str:    "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ||4.6||",
			result: "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ||4.6||",
		}, { // IAST to SLP
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "acalaśreṣṭha",
			result: "acalaSrezWa",
		}, {
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "ajo'pi sannavyayātmā bhūtānāmīśvaro'pi san |",
			result: "ajo'pi sannavyayAtmA BUtAnAmISvaro'pi san |",
		}, {
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "abhyutthānamadharmasya tadā''tmānaṃ sṛjāmyaham ||4.7||",
			result: "aByutTAnamaDarmasya tadA''tmAnaM sfjAmyaham ||4.7||",
		}, {
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "sa evāyaṃ mayā te'dya yogaḥ proktaḥ purātanaḥ |",
			result: "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
		}, {
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "bahūni me vyatītāni janmāni tava cārjuna|",
			result: "bahUni me vyatItAni janmAni tava cArjuna|",
		}, {
			params: Params{from: consts.IAST, to: consts.SLP1},
			str:    "prakṛtiṃ svāmadhiṣṭhāya saṃbhavāmyātmamāyayā ||4.6||",
			result: "prakftiM svAmaDizWAya saMBavAmyAtmamAyayA ||4.6||",
		}, { // HK to DEVANAGARY
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "yogaH",
			result: "योगः",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "san",
			result: "सन्",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "acalazreSTha",
			result: "अचलश्रेष्ठ",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "ajo'pi sannavyayAtmA bhUtAnAmIzvaro'pi san ।",
			result: "अजोऽपि सन्नव्ययात्मा भूतानामीश्वरोऽपि सन् ।",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "abhyutthAnamadharmasya tadA''tmAnaM sRjAmyaham ||4.7||",
			result: "अभ्युत्थानमधर्मस्य तदाऽऽत्मानं सृजाम्यहम् ।।४.७।।",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "sa evAyaM mayA te'dya yogaH proktaH purAtanaH |",
			result: "स एवायं मया तेऽद्य योगः प्रोक्तः पुरातनः ।",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "bahUni me vyatItAni janmAni tava cArjuna|",
			result: "बहूनि मे व्यतीतानि जन्मानि तव चार्जुन।",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "prakRtiM svAmadhiSThAya saMbhavAmyAtmamAyayA ||4.6||",
			result: "प्रकृतिं स्वामधिष्ठाय संभवाम्यात्ममायया ।।४.६।।",
		}, {
			params: Params{from: consts.HK, to: consts.DEVANAGARY},
			str:    "1234567890 ||4.6||",
			result: "१२३४५६७८९० ।।४.६।।",
		},
	}
)

func TestTransliterate(t *testing.T) {
	for caseNum, data := range validationTest {
		transliteration := NewTransliteration(data.params.from, data.params.to)
		cleanData := transliteration.Transliterate(data.str)
		fmt.Println(cleanData)
		if data.result != cleanData {
			t.Errorf("validation case #%d failed, expected %s but got %s", caseNum+1, data.result, cleanData)
		}
	}
}
