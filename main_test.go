package main

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	log.Print(search(nums, target))
}

func search(nums []int, target int) int {
	sortNum := sortNumber(nums)
	var (
		index       int
		newEndIndex int = len(sortNum)
		halfIndex   int = len(sortNum) / 2
	)

Loop:
	for {
		switch {
		case sortNum[halfIndex] == target:
			index = halfIndex
			break Loop
		case halfIndex <= 0:
			index = -1
			break Loop
		case target > sortNum[halfIndex]:
			mid := (len(sortNum) - halfIndex) / 2
			halfIndex += mid
			if halfIndex > newEndIndex {
				index = -1
				break Loop
			}
		case target < sortNum[halfIndex]:
			mid := halfIndex / 2
			newEndIndex = halfIndex
			halfIndex -= mid
		}
	}

	return index

}

func sortNumber(num []int) []int {
	for i, v := range num {
		if i == len(num)-1 {
			break
		}

		if v > num[i+1] {
			num[i], num[i+1] = num[i+1], v
		}
	}
	return num
}
