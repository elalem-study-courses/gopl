// package description
package main

import (
	"fmt"
	"sort"
)

var namesStringSlice sort.StringSlice
var namesSlice []string

func init() {
	namesStringSlice = sort.StringSlice{"mohamed", "ahmed", "mahmoud", "abd elslam", "elalem"}
	namesSlice = []string{"mohamed", "ahmed", "mahmoud", "abd elslam", "elalem"}
}

func main() {
	sort.Strings(namesSlice)
	namesStringSlice.Sort()

	fmt.Println(namesSlice)
	fmt.Println(namesStringSlice)
}
