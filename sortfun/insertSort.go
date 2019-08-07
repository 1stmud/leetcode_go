package main

import "fmt"

func insertSort (arr []int)[]int{

	for k,v:=range arr{
		current:=v
		preIndex:=k-1
		for preIndex>=0 && current<arr[preIndex]{
			arr[preIndex+1]=arr[preIndex]
			preIndex--
		}
		arr[preIndex+1]=current
	}


	return arr
}

func main(){

	a:=[]int{12,3,4,63,67,2,43,67,53}
	b:=insertSort(a)
	fmt.Print(b)

}