package main
import "fmt"

func bubbleSort(arr []int) []int{

	for i:=0; i<len(arr);i++{

		for j:=i+1;j<len(arr);j++{


			if arr[j]<arr[i]{

				arr[j],arr[i]=arr[i],arr[j]
			}
		}

	}
	return arr
}
func main(){

	a:=[]int{12,3,4,63,67,2,43,67,53}
	b:=bubbleSort(a)
	fmt.Print(b)
}