package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func closestToZero(s []int) int {
	sort.Ints(s)
	num := s[0]

	// In the cases where the distance is equal between two ints
	if len(s) == 2 {
		numB := s[1]
		if num < 0 && numB <= 0 {
			num = numB
		}
	}

	return num
}

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

	a.Log(fmt.Sprintf("The answer is: %d", closestToZero(searched)))

	return 0
}
