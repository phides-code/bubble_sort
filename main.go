/*Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence 
of up to 10 integers. The program should print the integers out on one line, in sorted order, 
from least to greatest. Use your favorite search tool to find a description of how the bubble 
sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice 
of integers as an argument and returns  nothing. The BubbleSort() function should modify the 
slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the 
position of two adjacent elements in the slice. You should write a Swap() function which 
performs this operation. Your Swap() function should take two arguments, a slice of integers 
and an index value i which indicates a position in the slice. The Swap() function should 
return nothing, but it should swap the contents of the slice in position i with the contents 
in position i+1.*/

package main

import (
	"fmt"
)

func main() {
	var intSequence = make([]int, 10)

	// prompt the user to enter a sequence of 10 ints
	fmt.Println("Please enter a sequence of 10 integers. ")

	for i := 0; i < 10; i++ {
		fmt.Printf("Int #%d: ", i+1)
		fmt.Scan(&intSequence[i])
	}

	// sort the sequence
	BubbleSort(intSequence)

	// print the sorted sequence out on one line
	fmt.Println(intSequence)
}

func BubbleSort(intSequence []int) {
	for j := 0; j < 9; j++ { // taking one slot at a time
		for i := 0; i < 9; i++ { // comparing it to each other slot
			// 9 because we don't need to check the last slot 
			if intSequence[i] > intSequence[i+1] {
				Swap(intSequence, i)
			}
		}
	}
}

func Swap(intSequence []int, i int) {
	hold := intSequence[i]
	intSequence[i] = intSequence[i+1]
	intSequence[i+1] = hold
}
