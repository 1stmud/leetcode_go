package main

import "fmt"

func selectSort( arr []int) []int{

	for i:=0; i<len(arr);i++{
		minIndex:=i
		for j:=i+1; j<len(arr);j++{

			if arr[j]< arr[minIndex]{
				minIndex=j
			}

		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]

	}
return arr

}

func main(){

	a:=[]int{12,3,4,63,67,2,43,67,53}
	b:=selectSort(a)
	fmt.Print(b)
}