package utils

import "sort"

// BubbleSort []int {9,8,7,6,5}
// BubbleSort []int {5,6,7,8,9}
func BubbleSort(elements []int)  {
	keepRunning := true
	for keepRunning{
		keepRunning = false
		for i :=0; i<len(elements)-1; i++ {
			if elements[i] > elements[i+1]{
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}

	return
	
}

func Sort(element []int) {

	if len(element) <1000{
		BubbleSort(element)
		return
	}

	sort.Ints(element)

}