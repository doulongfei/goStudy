package main

import "strings"

func replaceSpaceStr(s string) string {
	var split []string = strings.Split(s, " ")
	return strings.Join(split, "%20")
}
