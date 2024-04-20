package main

import "sort"

func main() {
}

func sortPeople(names []string, heights []int) []string {
	nameHeightArr := make(nameHeights, len(names))
	for i := 0; i < len(names); i++ {
		nameHeightArr[i] = nameHeight{
			Name:   names[i],
			Height: heights[i],
		}
	}
	sort.Sort(nameHeightArr)
	result := make([]string, len(names))
	for i := 0; i < len(names); i++ {
		result[len(names)-i-1] = nameHeightArr[i].Name
	}
	return result
}

type nameHeight struct {
	Name   string
	Height int
}

type nameHeights []nameHeight

func (nhs nameHeights) Len() int {
	return len(nhs)
}

func (nhs nameHeights) Less(i, j int) bool { return nhs[i].Height < nhs[j].Height }

func (nhs nameHeights) Swap(i, j int) { nhs[i], nhs[j] = nhs[j], nhs[i] }
