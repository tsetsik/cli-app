package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

// Temperature that returns closes temperature to given one
func (a *App) Temperature() int {
	needle, err := strconv.Atoi(a.args[1])
	if err != nil {
		a.Log("Sorry you need to provide valid int as second argument for needle")
		return 1
	}

	haystack := []int{}
	if err := json.Unmarshal([]byte(a.args[2]), &haystack); err != nil {
		a.Log("Sorry you need to provide valid slice of integers for haystack")
		return 1
	}

	mappings := map[int][]int{}
	lists := []int{}
	for _, v := range haystack {
		var num int
		if v > needle {
			num = (v - needle)
		} else {
			num = (needle - v)
		}

		mappings[num] = append(mappings[num], v)
		lists = append(lists, num)
	}

	sort.Ints(lists)
	searched := mappings[lists[0]]

	if len(searched) > 1 {
		sort.Ints(searched)
	}

	a.Log(fmt.Sprintf("The answer is: %d", searched[0]))

	return 0
}
