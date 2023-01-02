package main

import (
	mapset "../pkg/mod/github.com/deckarep/golang-set/v2"
	"fmt"
)

func main() {
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")

	fmt.Print(required)
}
