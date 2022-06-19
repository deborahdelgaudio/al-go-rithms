package main

import (
	"fmt"
	"github.com/deborahdelgaudio/al-go-rithms/arrays"
)

func main(){
	arr := []int{10, 5, 7, 9, 6, 8, 4, 1, 3, 2, 11}

	fmt.Println("Bubble Sorting Algorithm in Go")
	fmt.Printf("Not Sorted: %v", arr);
	fmt.Printf("\nSorted: %v", arrays.BubbleSort(arr));

	fmt.Println("\nMerge Sorting Algorithm in Go");
	fmt.Printf("Not Sorted: %v", arr);
	fmt.Printf("\nSorted: %v", arrays.MergeSort(arr));

	fmt.Println("\nInsert Sorting Algorithm in Go")
	fmt.Printf("Not Sorted: %v", arr);
	fmt.Printf("\nSorted: %v", arrays.InsertSort(arr));

}