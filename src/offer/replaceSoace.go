package offer

import "strings"

func replaceSpaceStr(s string) string {
	var split []string = strings.Split(s, " ")
	return strings.Join(split, "%20")
}

func replaceSpaceStr2(s string) string {

	var str string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += s[i : i+1]
		} else {
			str += "%20"
		}
	}
	return str

}
