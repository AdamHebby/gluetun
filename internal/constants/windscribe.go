package constants

import (
	"net"

	"github.com/qdm12/gluetun/internal/models"
)

//nolint:lll
const (
	WindscribeCertificate        = "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"
	WindscribeOpenvpnStaticKeyV1 = "5801926a57ac2ce27e3dfd1dd6ef82042d82bd4f3f0021296f57734f6f1ea714a6623845541c4b0c3dea0a050fe6746cb66dfab14cda27e5ae09d7c155aa554f399fa4a863f0e8c1af787e5c602a801d3a2ec41e395a978d56729457fe6102d7d9e9119aa83643210b33c678f9d4109e3154ac9c759e490cb309b319cf708cae83ddadc3060a7a26564d1a24411cd552fe6620ea16b755697a4fc5e6e9d0cfc0c5c4a1874685429046a424c026db672e4c2c492898052ba59128d46200b40f880027a8b6610a4d559bdc9346d33a0a6b08e75c7fd43192b162bfd0aef0c716b31584827693f676f9a5047123466f0654eade34972586b31c6ce7e395f4b478cb"
)

func WindscribeRegionChoices() (choices []string) {
	servers := WindscribeServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return choices
}

func WindscribeCityChoices() (choices []string) {
	servers := WindscribeServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return choices
}

func WindscribeHostnameChoices() (choices []string) {
	servers := WindscribeServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return choices
}

//nolint:lll
func WindscribeServers() []models.WindscribeServer {
	return []models.WindscribeServer{
		{Region: "Albania", City: "Tirana", Hostname: "al-002.whiskergalaxy.com", IP: net.IP{31, 171, 152, 179}},
		{Region: "Argentina", City: "Buenos Aires", Hostname: "ar-001.whiskergalaxy.com", IP: net.IP{200, 85, 152, 110}},
		{Region: "Argentina", City: "Buenos Aires", Hostname: "ar-003.whiskergalaxy.com", IP: net.IP{167, 250, 6, 121}},
		{Region: "Argentina", City: "Buenos Aires", Hostname: "ar-004.whiskergalaxy.com", IP: net.IP{190, 105, 236, 50}},
		{Region: "Argentina", City: "Buenos Aires", Hostname: "ar-005.whiskergalaxy.com", IP: net.IP{190, 105, 236, 32}},
		{Region: "Argentina", City: "Buenos Aires", Hostname: "ar-006.whiskergalaxy.com", IP: net.IP{190, 105, 236, 19}},
		{Region: "Australia", City: "Adelaide ", Hostname: "au-011.whiskergalaxy.com", IP: net.IP{103, 108, 92, 83}},
		{Region: "Australia", City: "Adelaide", Hostname: "au-008.whiskergalaxy.com", IP: net.IP{116, 90, 72, 243}},
		{Region: "Australia", City: "Brisbane", Hostname: "au-007.whiskergalaxy.com", IP: net.IP{103, 62, 50, 208}},
		{Region: "Australia", City: "Brisbane", Hostname: "au-014.whiskergalaxy.com", IP: net.IP{43, 245, 160, 35}},
		{Region: "Australia", City: "Canberra", Hostname: "au-010.whiskergalaxy.com", IP: net.IP{116, 206, 229, 131}},
		{Region: "Australia", City: "Melbourne ", Hostname: "au-005.whiskergalaxy.com", IP: net.IP{45, 121, 209, 160}},
		{Region: "Australia", City: "Melbourne ", Hostname: "au-013.whiskergalaxy.com", IP: net.IP{116, 206, 228, 67}},
		{Region: "Australia", City: "Perth", Hostname: "au-004.whiskergalaxy.com", IP: net.IP{45, 121, 208, 160}},
		{Region: "Australia", City: "Perth", Hostname: "au-012.whiskergalaxy.com", IP: net.IP{103, 77, 234, 211}},
		{Region: "Australia", City: "Sydney", Hostname: "au-009.whiskergalaxy.com", IP: net.IP{103, 77, 233, 67}},
		{Region: "Australia", City: "Sydney", Hostname: "au-015.whiskergalaxy.com", IP: net.IP{103, 1, 213, 211}},
		{Region: "Austria", City: "Vienna", Hostname: "at-001.whiskergalaxy.com", IP: net.IP{217, 64, 127, 11}},
		{Region: "Austria", City: "Vienna", Hostname: "at-002.whiskergalaxy.com", IP: net.IP{89, 187, 168, 66}},
		{Region: "Belgium", City: "Brussels", Hostname: "be-001.whiskergalaxy.com", IP: net.IP{194, 187, 251, 147}},
		{Region: "Belgium", City: "Brussels", Hostname: "be-002.whiskergalaxy.com", IP: net.IP{185, 232, 21, 131}},
		{Region: "Brazil", City: "Sao Paulo", Hostname: "br-004.whiskergalaxy.com", IP: net.IP{177, 67, 80, 59}},
		{Region: "Brazil", City: "Sao Paulo", Hostname: "br-005.whiskergalaxy.com", IP: net.IP{177, 54, 157, 178}},
		{Region: "Brazil", City: "Sao Paulo", Hostname: "br-006.whiskergalaxy.com", IP: net.IP{177, 54, 148, 247}},
		{Region: "Bulgaria", City: "Sofia", Hostname: "bg-001.whiskergalaxy.com", IP: net.IP{185, 94, 192, 35}},
		{Region: "Canada East", City: "Halifax", Hostname: "ca-029.whiskergalaxy.com", IP: net.IP{199, 204, 208, 158}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-004.whiskergalaxy.com", IP: net.IP{66, 70, 148, 80}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-027.whiskergalaxy.com", IP: net.IP{144, 168, 163, 160}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-028.whiskergalaxy.com", IP: net.IP{144, 168, 163, 193}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-032.whiskergalaxy.com", IP: net.IP{104, 227, 235, 129}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-033.whiskergalaxy.com", IP: net.IP{198, 8, 85, 195}},
		{Region: "Canada East", City: "Montreal", Hostname: "ca-034.whiskergalaxy.com", IP: net.IP{198, 8, 85, 210}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-002.whiskergalaxy.com", IP: net.IP{104, 254, 92, 11}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-009.whiskergalaxy.com", IP: net.IP{104, 254, 92, 91}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-017.whiskergalaxy.com", IP: net.IP{184, 75, 212, 91}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-025.whiskergalaxy.com", IP: net.IP{192, 190, 19, 65}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-026.whiskergalaxy.com", IP: net.IP{192, 190, 19, 97}},
		{Region: "Canada East", City: "Toronto", Hostname: "ca-030.whiskergalaxy.com", IP: net.IP{23, 154, 160, 177}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-005.whiskergalaxy.com", IP: net.IP{162, 221, 207, 95}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-011.whiskergalaxy.com", IP: net.IP{104, 218, 61, 1}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-012.whiskergalaxy.com", IP: net.IP{104, 218, 61, 33}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-016.whiskergalaxy.com", IP: net.IP{208, 78, 41, 1}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-017.whiskergalaxy.com", IP: net.IP{208, 78, 41, 131}},
		{Region: "Canada West", City: "Vancouver", Hostname: "ca-west-019.whiskergalaxy.com", IP: net.IP{208, 78, 41, 163}},
		{Region: "Colombia", City: "Bogota", Hostname: "co-001.whiskergalaxy.com", IP: net.IP{138, 121, 203, 203}},
		{Region: "Colombia", City: "Bogota", Hostname: "co-002.whiskergalaxy.com", IP: net.IP{138, 186, 141, 155}},
		{Region: "Croatia", City: "Zagreb", Hostname: "hr-002.whiskergalaxy.com", IP: net.IP{85, 10, 56, 129}},
		{Region: "Cyprus", City: "Nicosia", Hostname: "cy-001.whiskergalaxy.com", IP: net.IP{157, 97, 132, 43}},
		{Region: "Czech Republic", City: "Prague ", Hostname: "cz-002.whiskergalaxy.com", IP: net.IP{185, 246, 210, 2}},
		{Region: "Czech Republic", City: "Prague", Hostname: "cz-001.whiskergalaxy.com", IP: net.IP{185, 156, 174, 11}},
		{Region: "Denmark", City: "Copenhagen", Hostname: "dk-001.whiskergalaxy.com", IP: net.IP{185, 206, 224, 195}},
		{Region: "Denmark", City: "Copenhagen", Hostname: "dk-002.whiskergalaxy.com", IP: net.IP{134, 90, 149, 147}},
		{Region: "Estonia", City: "Tallinn", Hostname: "ee-001.whiskergalaxy.com", IP: net.IP{46, 22, 211, 251}},
		{Region: "Estonia", City: "Tallinn", Hostname: "ee-002.whiskergalaxy.com", IP: net.IP{196, 196, 216, 131}},
		{Region: "Fake Antarctica", City: "Troll", Hostname: "aq-001.whiskergalaxy.com", IP: net.IP{23, 154, 160, 212}},
		{Region: "Fake Antarctica", City: "Troll", Hostname: "aq-002.whiskergalaxy.com", IP: net.IP{23, 154, 160, 222}},
		{Region: "Finland", City: "Helsinki", Hostname: "fi-002.whiskergalaxy.com", IP: net.IP{185, 112, 82, 227}},
		{Region: "Finland", City: "Helsinki", Hostname: "fi-003.whiskergalaxy.com", IP: net.IP{194, 34, 133, 82}},
		{Region: "Finland", City: "Helsinki", Hostname: "fi-004.whiskergalaxy.com", IP: net.IP{196, 244, 192, 51}},
		{Region: "France", City: "Paris", Hostname: "fr-004.whiskergalaxy.com", IP: net.IP{185, 156, 173, 187}},
		{Region: "France", City: "Paris", Hostname: "fr-005.whiskergalaxy.com", IP: net.IP{82, 102, 18, 35}},
		{Region: "France", City: "Paris", Hostname: "fr-008.whiskergalaxy.com", IP: net.IP{84, 17, 42, 34}},
		{Region: "France", City: "Paris", Hostname: "fr-009.whiskergalaxy.com", IP: net.IP{84, 17, 42, 2}},
		{Region: "France", City: "Paris", Hostname: "fr-011.whiskergalaxy.com", IP: net.IP{45, 89, 174, 35}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-003.whiskergalaxy.com", IP: net.IP{89, 249, 65, 19}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-006.whiskergalaxy.com", IP: net.IP{185, 130, 184, 195}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-009.whiskergalaxy.com", IP: net.IP{195, 181, 170, 66}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-010.whiskergalaxy.com", IP: net.IP{195, 181, 175, 98}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-011.whiskergalaxy.com", IP: net.IP{217, 138, 194, 115}},
		{Region: "Germany", City: "Frankfurt", Hostname: "de-012.whiskergalaxy.com", IP: net.IP{45, 87, 212, 51}},
		{Region: "Greece", City: "Athens", Hostname: "gr-002.whiskergalaxy.com", IP: net.IP{78, 108, 38, 155}},
		{Region: "Greece", City: "Athens", Hostname: "gr-004.whiskergalaxy.com", IP: net.IP{185, 226, 64, 111}},
		{Region: "Greece", City: "Athens", Hostname: "gr-005.whiskergalaxy.com", IP: net.IP{188, 123, 126, 146}},
		{Region: "Hong Kong", City: "Hong Kong", Hostname: "hk-005.whiskergalaxy.com", IP: net.IP{103, 10, 197, 99}},
		{Region: "Hong Kong", City: "Hong Kong", Hostname: "hk-006.whiskergalaxy.com", IP: net.IP{84, 17, 57, 114}},
		{Region: "Hungary", City: "Budapest", Hostname: "hu-001.whiskergalaxy.com", IP: net.IP{185, 104, 187, 43}},
		{Region: "Iceland", City: "Reykjavik", Hostname: "is-001.whiskergalaxy.com", IP: net.IP{82, 221, 139, 38}},
		{Region: "Iceland", City: "Reykjavik", Hostname: "is-002.whiskergalaxy.com", IP: net.IP{185, 165, 170, 2}},
		{Region: "India", City: "Chennai", Hostname: "in-005.whiskergalaxy.com", IP: net.IP{169, 38, 68, 188}},
		{Region: "India", City: "Chennai", Hostname: "in-006.whiskergalaxy.com", IP: net.IP{169, 38, 72, 14}},
		{Region: "India", City: "Chennai", Hostname: "in-007.whiskergalaxy.com", IP: net.IP{169, 38, 72, 12}},
		{Region: "India", City: "Mumbai", Hostname: "in-009.whiskergalaxy.com", IP: net.IP{165, 231, 253, 211}},
		{Region: "India", City: "New Delhi", Hostname: "in-008.whiskergalaxy.com", IP: net.IP{103, 205, 140, 227}},
		{Region: "Indonesia", City: "Jakarta", Hostname: "id-002.whiskergalaxy.com", IP: net.IP{45, 127, 134, 91}},
		{Region: "Ireland", City: "Dublin", Hostname: "ie-001.whiskergalaxy.com", IP: net.IP{185, 24, 232, 146}},
		{Region: "Ireland", City: "Dublin", Hostname: "ie-002.whiskergalaxy.com", IP: net.IP{185, 104, 219, 2}},
		{Region: "Ireland", City: "Dublin", Hostname: "ie-003.whiskergalaxy.com", IP: net.IP{23, 92, 127, 35}},
		{Region: "Israel", City: "Ashdod", Hostname: "il-002.whiskergalaxy.com", IP: net.IP{185, 191, 205, 139}},
		{Region: "Israel", City: "Jerusalem", Hostname: "il-001.whiskergalaxy.com", IP: net.IP{160, 116, 0, 27}},
		{Region: "Italy", City: "Milan", Hostname: "it-001.whiskergalaxy.com", IP: net.IP{37, 120, 135, 83}},
		{Region: "Italy", City: "Milan", Hostname: "it-004.whiskergalaxy.com", IP: net.IP{84, 17, 59, 66}},
		{Region: "Italy", City: "Milan", Hostname: "it-005.whiskergalaxy.com", IP: net.IP{89, 40, 182, 3}},
		{Region: "Italy", City: "Rome", Hostname: "it-003.whiskergalaxy.com", IP: net.IP{87, 101, 94, 195}},
		{Region: "Italy", City: "Rome", Hostname: "it-006.whiskergalaxy.com", IP: net.IP{37, 120, 207, 19}},
		{Region: "Japan", City: "Tokyo", Hostname: "jp-004.whiskergalaxy.com", IP: net.IP{193, 148, 16, 243}},
		{Region: "Japan", City: "Tokyo", Hostname: "jp-005.whiskergalaxy.com", IP: net.IP{89, 187, 161, 114}},
		{Region: "Latvia", City: "Riga", Hostname: "lv-003.whiskergalaxy.com", IP: net.IP{85, 254, 72, 23}},
		{Region: "Latvia", City: "Riga", Hostname: "lv-004.whiskergalaxy.com", IP: net.IP{89, 111, 33, 220}},
		{Region: "Lithuania", City: "Siauliai", Hostname: "lt-003.whiskergalaxy.com", IP: net.IP{85, 206, 163, 225}},
		{Region: "Malaysia", City: "Kuala Lumpur", Hostname: "my-001.whiskergalaxy.com", IP: net.IP{103, 106, 250, 31}},
		{Region: "Malaysia", City: "Kuala Lumpur", Hostname: "my-003.whiskergalaxy.com", IP: net.IP{103, 212, 69, 232}},
		{Region: "Mexico", City: "Guadalajara", Hostname: "mx-007.whiskergalaxy.com", IP: net.IP{201, 131, 125, 107}},
		{Region: "Mexico", City: "Guadalajara", Hostname: "mx-008.whiskergalaxy.com", IP: net.IP{143, 255, 57, 67}},
		{Region: "Mexico", City: "Mexico City", Hostname: "mx-009.whiskergalaxy.com", IP: net.IP{190, 103, 179, 211}},
		{Region: "Mexico", City: "Mexico City", Hostname: "mx-010.whiskergalaxy.com", IP: net.IP{190, 103, 179, 217}},
		{Region: "Moldova", City: "Chisinau", Hostname: "md-002.whiskergalaxy.com", IP: net.IP{178, 175, 144, 123}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-001.whiskergalaxy.com", IP: net.IP{46, 166, 143, 98}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-005.whiskergalaxy.com", IP: net.IP{185, 212, 171, 131}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-008.whiskergalaxy.com", IP: net.IP{185, 253, 96, 3}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-011.whiskergalaxy.com", IP: net.IP{84, 17, 46, 2}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-012.whiskergalaxy.com", IP: net.IP{37, 120, 192, 19}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-013.whiskergalaxy.com", IP: net.IP{72, 11, 157, 67}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-014.whiskergalaxy.com", IP: net.IP{72, 11, 157, 35}},
		{Region: "Netherlands", City: "Amsterdam", Hostname: "nl-015.whiskergalaxy.com", IP: net.IP{109, 201, 130, 2}},
		{Region: "New Zealand", City: "Auckland ", Hostname: "nz-003.whiskergalaxy.com", IP: net.IP{103, 108, 94, 163}},
		{Region: "New Zealand", City: "Auckland", Hostname: "nz-002.whiskergalaxy.com", IP: net.IP{103, 62, 49, 113}},
		{Region: "North Macedonia", City: "Skopje", Hostname: "mk-001.whiskergalaxy.com", IP: net.IP{185, 225, 28, 51}},
		{Region: "Norway", City: "Oslo", Hostname: "no-003.whiskergalaxy.com", IP: net.IP{185, 206, 225, 131}},
		{Region: "Norway", City: "Oslo", Hostname: "no-006.whiskergalaxy.com", IP: net.IP{37, 120, 203, 67}},
		{Region: "Panama", City: "Panama City", Hostname: "pa-001.whiskergalaxy.com", IP: net.IP{138, 186, 142, 203}},
		{Region: "Peru", City: "Lima", Hostname: "pe-002.whiskergalaxy.com", IP: net.IP{190, 120, 229, 139}},
		{Region: "Philippines", City: "Manila", Hostname: "ph-003.whiskergalaxy.com", IP: net.IP{141, 98, 215, 211}},
		{Region: "Philippines", City: "San Antonio", Hostname: "ph-002.whiskergalaxy.com", IP: net.IP{103, 103, 0, 118}},
		{Region: "Poland", City: "Warsaw", Hostname: "pl-002.whiskergalaxy.com", IP: net.IP{185, 244, 214, 35}},
		{Region: "Poland", City: "Warsaw", Hostname: "pl-004.whiskergalaxy.com", IP: net.IP{84, 17, 55, 98}},
		{Region: "Poland", City: "Warsaw", Hostname: "pl-005.whiskergalaxy.com", IP: net.IP{5, 133, 11, 116}},
		{Region: "Portugal", City: "Lisbon", Hostname: "pt-002.whiskergalaxy.com", IP: net.IP{94, 46, 13, 215}},
		{Region: "Portugal", City: "Lisbon", Hostname: "pt-003.whiskergalaxy.com", IP: net.IP{185, 15, 21, 66}},
		{Region: "Romania", City: "Bucharest", Hostname: "ro-006.whiskergalaxy.com", IP: net.IP{89, 46, 103, 147}},
		{Region: "Romania", City: "Bucharest", Hostname: "ro-008.whiskergalaxy.com", IP: net.IP{91, 207, 102, 147}},
		{Region: "Russia", City: "Moscow", Hostname: "ru-010.whiskergalaxy.com", IP: net.IP{95, 213, 193, 227}},
		{Region: "Russia", City: "Moscow", Hostname: "ru-011.whiskergalaxy.com", IP: net.IP{95, 213, 193, 195}},
		{Region: "Russia", City: "Saint Petersburg", Hostname: "ru-008.whiskergalaxy.com", IP: net.IP{94, 242, 62, 19}},
		{Region: "Russia", City: "Saint Petersburg", Hostname: "ru-009.whiskergalaxy.com", IP: net.IP{94, 242, 62, 67}},
		{Region: "Russia", City: "Saint Petersburg", Hostname: "ru-012.whiskergalaxy.com", IP: net.IP{188, 124, 42, 115}},
		{Region: "Russia", City: "Saint Petersburg", Hostname: "ru-013.whiskergalaxy.com", IP: net.IP{188, 124, 42, 99}},
		{Region: "Serbia", City: "Belgrade", Hostname: "rs-003.whiskergalaxy.com", IP: net.IP{141, 98, 103, 19}},
		{Region: "Singapore", City: "Singapore", Hostname: "sg-003.whiskergalaxy.com", IP: net.IP{185, 200, 117, 163}},
		{Region: "Singapore", City: "Singapore", Hostname: "sg-004.whiskergalaxy.com", IP: net.IP{82, 102, 25, 131}},
		{Region: "Singapore", City: "Singapore", Hostname: "sg-005.whiskergalaxy.com", IP: net.IP{103, 62, 48, 224}},
		{Region: "Singapore", City: "Singapore", Hostname: "sg-006.whiskergalaxy.com", IP: net.IP{156, 146, 56, 98}},
		{Region: "Singapore", City: "Singapore", Hostname: "sg-007.whiskergalaxy.com", IP: net.IP{156, 146, 56, 111}},
		{Region: "Slovakia", City: "Bratislava", Hostname: "sk-001.whiskergalaxy.com", IP: net.IP{185, 245, 85, 3}},
		{Region: "South Africa", City: "Johannesburg", Hostname: "za-001.whiskergalaxy.com", IP: net.IP{197, 242, 157, 235}},
		{Region: "South Africa", City: "Johannesburg", Hostname: "za-002.whiskergalaxy.com", IP: net.IP{129, 232, 167, 211}},
		{Region: "South Africa", City: "Johannesburg", Hostname: "za-003.whiskergalaxy.com", IP: net.IP{197, 242, 156, 53}},
		{Region: "South Africa", City: "Johannesburg", Hostname: "za-004.whiskergalaxy.com", IP: net.IP{165, 73, 248, 91}},
		{Region: "South Korea", City: "Seoul", Hostname: "kr-001.whiskergalaxy.com", IP: net.IP{103, 212, 223, 3}},
		{Region: "South Korea", City: "Seoul", Hostname: "kr-002.whiskergalaxy.com", IP: net.IP{218, 232, 76, 179}},
		{Region: "South Korea", City: "Seoul", Hostname: "kr-005.whiskergalaxy.com", IP: net.IP{45, 133, 194, 235}},
		{Region: "Spain", City: "Barcelona", Hostname: "es-004.whiskergalaxy.com", IP: net.IP{37, 120, 142, 227}},
		{Region: "Spain", City: "Madrid", Hostname: "es-002.whiskergalaxy.com", IP: net.IP{89, 238, 178, 43}},
		{Region: "Spain", City: "Madrid", Hostname: "es-003.whiskergalaxy.com", IP: net.IP{217, 138, 218, 99}},
		{Region: "Sweden", City: "Stockholm", Hostname: "se-001.whiskergalaxy.com", IP: net.IP{31, 13, 191, 67}},
		{Region: "Sweden", City: "Stockholm", Hostname: "se-002.whiskergalaxy.com", IP: net.IP{79, 142, 76, 198}},
		{Region: "Sweden", City: "Stockholm", Hostname: "se-003.whiskergalaxy.com", IP: net.IP{195, 181, 166, 129}},
		{Region: "Switzerland", City: "Zurich", Hostname: "ch-001.whiskergalaxy.com", IP: net.IP{31, 7, 57, 242}},
		{Region: "Switzerland", City: "Zurich", Hostname: "ch-003.whiskergalaxy.com", IP: net.IP{185, 156, 175, 179}},
		{Region: "Switzerland", City: "Zurich", Hostname: "ch-005.whiskergalaxy.com", IP: net.IP{89, 187, 165, 98}},
		{Region: "Switzerland", City: "Zurich", Hostname: "ch-006.whiskergalaxy.com", IP: net.IP{84, 17, 53, 2}},
		{Region: "Switzerland", City: "Zurich", Hostname: "ch-008.whiskergalaxy.com", IP: net.IP{37, 120, 213, 163}},
		{Region: "Taiwan", City: "Taipei", Hostname: "tw-008.whiskergalaxy.com", IP: net.IP{103, 4, 29, 77}},
		{Region: "Taiwan", City: "Taipei", Hostname: "tw-009.whiskergalaxy.com", IP: net.IP{185, 189, 160, 12}},
		{Region: "Taiwan", City: "Taipei", Hostname: "tw-010.whiskergalaxy.com", IP: net.IP{185, 189, 160, 27}},
		{Region: "Taiwan", City: "Taipei", Hostname: "tw-011.whiskergalaxy.com", IP: net.IP{185, 189, 160, 32}},
		{Region: "Thailand", City: "Bangkok", Hostname: "th-003.whiskergalaxy.com", IP: net.IP{27, 254, 130, 221}},
		{Region: "Thailand", City: "Bangkok", Hostname: "th-005.whiskergalaxy.com", IP: net.IP{202, 129, 16, 147}},
		{Region: "Thailand", City: "Bangkok", Hostname: "th-006.whiskergalaxy.com", IP: net.IP{202, 129, 16, 155}},
		{Region: "Tunisia", City: "Tunis", Hostname: "tn-001.whiskergalaxy.com", IP: net.IP{41, 231, 5, 23}},
		{Region: "Turkey", City: "Bursa", Hostname: "tr-001.whiskergalaxy.com", IP: net.IP{45, 123, 118, 156}},
		{Region: "Turkey", City: "Istanbul", Hostname: "tr-004.whiskergalaxy.com", IP: net.IP{45, 123, 119, 11}},
		{Region: "Turkey", City: "Istanbul", Hostname: "tr-006.whiskergalaxy.com", IP: net.IP{185, 125, 33, 227}},
		{Region: "Turkey", City: "Istanbul", Hostname: "tr-009.whiskergalaxy.com", IP: net.IP{79, 98, 131, 43}},
		{Region: "Turkey", City: "Istanbul", Hostname: "tr-011.whiskergalaxy.com", IP: net.IP{176, 53, 113, 163}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-016.whiskergalaxy.com", IP: net.IP{104, 129, 18, 3}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-020.whiskergalaxy.com", IP: net.IP{104, 129, 18, 131}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-034.whiskergalaxy.com", IP: net.IP{161, 129, 70, 195}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-046.whiskergalaxy.com", IP: net.IP{198, 12, 76, 211}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-049.whiskergalaxy.com", IP: net.IP{107, 150, 31, 3}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-050.whiskergalaxy.com", IP: net.IP{107, 150, 31, 67}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-051.whiskergalaxy.com", IP: net.IP{162, 222, 198, 67}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-054.whiskergalaxy.com", IP: net.IP{104, 223, 92, 163}},
		{Region: "US Central", City: "Atlanta", Hostname: "us-central-056.whiskergalaxy.com", IP: net.IP{206, 217, 143, 131}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-014.whiskergalaxy.com", IP: net.IP{69, 12, 94, 67}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-029.whiskergalaxy.com", IP: net.IP{198, 55, 125, 195}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-036.whiskergalaxy.com", IP: net.IP{204, 44, 112, 67}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-037.whiskergalaxy.com", IP: net.IP{204, 44, 112, 131}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-044.whiskergalaxy.com", IP: net.IP{206, 217, 139, 195}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-045.whiskergalaxy.com", IP: net.IP{172, 241, 131, 129}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-055.whiskergalaxy.com", IP: net.IP{206, 217, 139, 19}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-057.whiskergalaxy.com", IP: net.IP{172, 241, 26, 78}},
		{Region: "US Central", City: "Dallas", Hostname: "us-central-060.whiskergalaxy.com", IP: net.IP{198, 55, 126, 131}},
		{Region: "US Central", City: "Denver", Hostname: "us-central-043.whiskergalaxy.com", IP: net.IP{199, 115, 96, 83}},
		{Region: "US Central", City: "Denver", Hostname: "us-central-058.whiskergalaxy.com", IP: net.IP{198, 54, 128, 116}},
		{Region: "US Central", City: "Denver", Hostname: "us-central-062.whiskergalaxy.com", IP: net.IP{174, 128, 251, 147}},
		{Region: "US Central", City: "Kansas City", Hostname: "us-central-063.whiskergalaxy.com", IP: net.IP{38, 146, 5, 51}},
		{Region: "US Central", City: "Salt Lake City", Hostname: "us-central-047.whiskergalaxy.com", IP: net.IP{107, 182, 234, 240}},
		{Region: "US Central", City: "Salt Lake City", Hostname: "us-central-052.whiskergalaxy.com", IP: net.IP{67, 212, 238, 196}},
		{Region: "US East", City: "Boston", Hostname: "us-east-039.whiskergalaxy.com", IP: net.IP{199, 217, 104, 227}},
		{Region: "US East", City: "Boston", Hostname: "us-east-051.whiskergalaxy.com", IP: net.IP{199, 217, 105, 227}},
		{Region: "US East", City: "Buffalo", Hostname: "us-east-045.whiskergalaxy.com", IP: net.IP{104, 168, 34, 147}},
		{Region: "US East", City: "Buffalo", Hostname: "us-east-065.whiskergalaxy.com", IP: net.IP{198, 12, 64, 35}},
		{Region: "US East", City: "Charlotte", Hostname: "us-east-040.whiskergalaxy.com", IP: net.IP{67, 21, 32, 145}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-015.whiskergalaxy.com", IP: net.IP{68, 235, 50, 227}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-019.whiskergalaxy.com", IP: net.IP{23, 226, 141, 195}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-022.whiskergalaxy.com", IP: net.IP{167, 160, 172, 3}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-047.whiskergalaxy.com", IP: net.IP{23, 83, 91, 170}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-053.whiskergalaxy.com", IP: net.IP{107, 150, 29, 131}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-069.whiskergalaxy.com", IP: net.IP{68, 235, 35, 172}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-071.whiskergalaxy.com", IP: net.IP{68, 235, 35, 12}},
		{Region: "US East", City: "Chicago", Hostname: "us-east-077.whiskergalaxy.com", IP: net.IP{68, 235, 43, 204}},
		{Region: "US East", City: "Cleveland", Hostname: "us-east-078.whiskergalaxy.com", IP: net.IP{38, 101, 74, 19}},
		{Region: "US East", City: "Columbus", Hostname: "us-east-059.whiskergalaxy.com", IP: net.IP{67, 219, 146, 67}},
		{Region: "US East", City: "Detroit", Hostname: "us-east-079.whiskergalaxy.com", IP: net.IP{104, 244, 210, 51}},
		{Region: "US East", City: "Miami", Hostname: "us-east-006.whiskergalaxy.com", IP: net.IP{173, 44, 36, 67}},
		{Region: "US East", City: "Miami", Hostname: "us-east-012.whiskergalaxy.com", IP: net.IP{45, 87, 214, 35}},
		{Region: "US East", City: "Miami", Hostname: "us-east-028.whiskergalaxy.com", IP: net.IP{104, 223, 127, 195}},
		{Region: "US East", City: "Miami", Hostname: "us-east-049.whiskergalaxy.com", IP: net.IP{23, 82, 136, 93}},
		{Region: "US East", City: "Miami", Hostname: "us-east-067.whiskergalaxy.com", IP: net.IP{86, 106, 87, 83}},
		{Region: "US East", City: "New Jersey", Hostname: "us-east-020.whiskergalaxy.com", IP: net.IP{162, 222, 195, 67}},
		{Region: "US East", City: "New Jersey", Hostname: "us-east-054.whiskergalaxy.com", IP: net.IP{167, 160, 167, 195}},
		{Region: "US East", City: "New York", Hostname: "us-east-013.whiskergalaxy.com", IP: net.IP{185, 232, 22, 195}},
		{Region: "US East", City: "New York", Hostname: "us-east-046.whiskergalaxy.com", IP: net.IP{206, 217, 129, 227}},
		{Region: "US East", City: "New York", Hostname: "us-east-050.whiskergalaxy.com", IP: net.IP{173, 208, 45, 33}},
		{Region: "US East", City: "New York", Hostname: "us-east-064.whiskergalaxy.com", IP: net.IP{206, 217, 128, 3}},
		{Region: "US East", City: "New York", Hostname: "us-east-068.whiskergalaxy.com", IP: net.IP{142, 234, 200, 176}},
		{Region: "US East", City: "New York", Hostname: "us-east-073.whiskergalaxy.com", IP: net.IP{217, 138, 255, 163}},
		{Region: "US East", City: "New York", Hostname: "us-east-074.whiskergalaxy.com", IP: net.IP{217, 138, 255, 179}},
		{Region: "US East", City: "Orlando", Hostname: "us-east-052.whiskergalaxy.com", IP: net.IP{198, 147, 22, 225}},
		{Region: "US East", City: "Philadelphia", Hostname: "us-east-060.whiskergalaxy.com", IP: net.IP{76, 72, 175, 99}},
		{Region: "US East", City: "Philadelphia", Hostname: "us-east-061.whiskergalaxy.com", IP: net.IP{156, 96, 59, 102}},
		{Region: "US East", City: "Washington DC", Hostname: "us-east-048.whiskergalaxy.com", IP: net.IP{23, 82, 8, 143}},
		{Region: "US East", City: "Washington DC", Hostname: "us-east-055.whiskergalaxy.com", IP: net.IP{23, 105, 170, 139}},
		{Region: "US East", City: "Washington DC", Hostname: "us-east-057.whiskergalaxy.com", IP: net.IP{23, 105, 170, 130}},
		{Region: "US East", City: "Washington DC", Hostname: "us-east-058.whiskergalaxy.com", IP: net.IP{23, 105, 170, 151}},
		{Region: "US West", City: "Bend", Hostname: "us-west-038.whiskergalaxy.com", IP: net.IP{104, 152, 222, 33}},
		{Region: "US West", City: "Las Vegas", Hostname: "us-west-018.whiskergalaxy.com", IP: net.IP{82, 102, 30, 67}},
		{Region: "US West", City: "Las Vegas", Hostname: "us-west-030.whiskergalaxy.com", IP: net.IP{37, 120, 147, 163}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-004.whiskergalaxy.com", IP: net.IP{185, 236, 200, 35}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-015.whiskergalaxy.com", IP: net.IP{216, 45, 53, 131}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-027.whiskergalaxy.com", IP: net.IP{212, 103, 49, 67}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-040.whiskergalaxy.com", IP: net.IP{89, 187, 185, 34}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-044.whiskergalaxy.com", IP: net.IP{192, 3, 20, 51}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-047.whiskergalaxy.com", IP: net.IP{172, 241, 214, 202}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-055.whiskergalaxy.com", IP: net.IP{104, 129, 3, 67}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-059.whiskergalaxy.com", IP: net.IP{104, 129, 3, 163}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-060.whiskergalaxy.com", IP: net.IP{217, 138, 217, 51}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-063.whiskergalaxy.com", IP: net.IP{198, 23, 242, 147}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-065.whiskergalaxy.com", IP: net.IP{217, 138, 217, 211}},
		{Region: "US West", City: "Los Angeles", Hostname: "us-west-066.whiskergalaxy.com", IP: net.IP{89, 187, 187, 98}},
		{Region: "US West", City: "Phoenix", Hostname: "us-west-046.whiskergalaxy.com", IP: net.IP{23, 83, 130, 166}},
		{Region: "US West", City: "Phoenix", Hostname: "us-west-061.whiskergalaxy.com", IP: net.IP{23, 83, 131, 187}},
		{Region: "US West", City: "San Francisco", Hostname: "us-west-048.whiskergalaxy.com", IP: net.IP{172, 241, 250, 131}},
		{Region: "US West", City: "San Francisco", Hostname: "us-west-053.whiskergalaxy.com", IP: net.IP{209, 58, 129, 121}},
		{Region: "US West", City: "San Francisco", Hostname: "us-west-054.whiskergalaxy.com", IP: net.IP{172, 255, 125, 141}},
		{Region: "US West", City: "San Jose", Hostname: "us-west-052.whiskergalaxy.com", IP: net.IP{66, 115, 176, 3}},
		{Region: "US West", City: "Santa Clara", Hostname: "us-west-050.whiskergalaxy.com", IP: net.IP{167, 88, 60, 227}},
		{Region: "US West", City: "Santa Clara", Hostname: "us-west-051.whiskergalaxy.com", IP: net.IP{167, 88, 60, 243}},
		{Region: "US West", City: "Seattle", Hostname: "us-west-043.whiskergalaxy.com", IP: net.IP{23, 94, 74, 99}},
		{Region: "US West", City: "Seattle", Hostname: "us-west-045.whiskergalaxy.com", IP: net.IP{64, 120, 2, 174}},
		{Region: "US West", City: "Seattle", Hostname: "us-west-056.whiskergalaxy.com", IP: net.IP{104, 129, 56, 67}},
		{Region: "US West", City: "Seattle", Hostname: "us-west-057.whiskergalaxy.com", IP: net.IP{104, 129, 56, 131}},
		{Region: "US West", City: "Seattle", Hostname: "us-west-062.whiskergalaxy.com", IP: net.IP{198, 12, 116, 195}},
		{Region: "Ukraine", City: "Kyiv", Hostname: "ua-006.whiskergalaxy.com", IP: net.IP{45, 141, 156, 11}},
		{Region: "Ukraine", City: "Kyiv", Hostname: "ua-007.whiskergalaxy.com", IP: net.IP{45, 141, 156, 50}},
		{Region: "United Arab Emirates", City: "Dubai", Hostname: "ae-001.whiskergalaxy.com", IP: net.IP{45, 9, 249, 43}},
		{Region: "United Kingdom", City: "Edinburgh", Hostname: "uk-026.whiskergalaxy.com", IP: net.IP{193, 36, 118, 243}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-007.whiskergalaxy.com", IP: net.IP{185, 212, 168, 133}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-013.whiskergalaxy.com", IP: net.IP{89, 238, 150, 229}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-014.whiskergalaxy.com", IP: net.IP{2, 58, 29, 145}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-015.whiskergalaxy.com", IP: net.IP{2, 58, 29, 17}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-017.whiskergalaxy.com", IP: net.IP{84, 17, 50, 130}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-021.whiskergalaxy.com", IP: net.IP{212, 102, 63, 32}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-022.whiskergalaxy.com", IP: net.IP{212, 102, 63, 62}},
		{Region: "United Kingdom", City: "London", Hostname: "uk-024.whiskergalaxy.com", IP: net.IP{217, 138, 254, 51}},
		{Region: "United Kingdom", City: "Manchester", Hostname: "uk-008.whiskergalaxy.com", IP: net.IP{81, 92, 207, 69}},
		{Region: "United Kingdom", City: "Manchester", Hostname: "uk-010.whiskergalaxy.com", IP: net.IP{89, 238, 135, 133}},
		{Region: "United Kingdom", City: "Manchester", Hostname: "uk-025.whiskergalaxy.com", IP: net.IP{89, 44, 201, 99}},
		{Region: "Vietnam", City: "Hanoi", Hostname: "vn-001.whiskergalaxy.com", IP: net.IP{103, 9, 76, 197}},
		{Region: "Vietnam", City: "Hanoi", Hostname: "vn-002.whiskergalaxy.com", IP: net.IP{103, 9, 79, 186}},
		{Region: "Vietnam", City: "Hanoi", Hostname: "vn-003.whiskergalaxy.com", IP: net.IP{103, 9, 79, 219}},
		{Region: "WINDFLIX CA", City: "Toronto", Hostname: "wf-ca-003.whiskergalaxy.com", IP: net.IP{104, 218, 60, 111}},
		{Region: "WINDFLIX CA", City: "Toronto", Hostname: "wf-ca-004.whiskergalaxy.com", IP: net.IP{104, 254, 92, 99}},
		{Region: "WINDFLIX JP", City: "Tokyo", Hostname: "wf-jp-002.whiskergalaxy.com", IP: net.IP{5, 181, 235, 67}},
		{Region: "WINDFLIX UK", City: "London", Hostname: "wf-uk-001.whiskergalaxy.com", IP: net.IP{45, 9, 248, 3}},
		{Region: "WINDFLIX UK", City: "London", Hostname: "wf-uk-006.whiskergalaxy.com", IP: net.IP{81, 92, 200, 85}},
		{Region: "WINDFLIX UK", City: "London", Hostname: "wf-uk-007.whiskergalaxy.com", IP: net.IP{89, 47, 62, 83}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-010.whiskergalaxy.com", IP: net.IP{38, 132, 122, 195}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-011.whiskergalaxy.com", IP: net.IP{38, 132, 122, 131}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-012.whiskergalaxy.com", IP: net.IP{185, 232, 22, 131}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-013.whiskergalaxy.com", IP: net.IP{217, 138, 206, 211}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-014.whiskergalaxy.com", IP: net.IP{77, 81, 136, 99}},
		{Region: "WINDFLIX US", City: "New York", Hostname: "wf-us-015.whiskergalaxy.com", IP: net.IP{38, 132, 101, 211}},
	}
}
