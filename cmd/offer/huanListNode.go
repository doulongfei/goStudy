package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")
	fmt.Print(required)
}
