package main

import "fmt"

func InsertionSort(a []int)[]int{
	//注意课本上的数组下标是从1开始，而实际实现中是从0开始
	for j:=1;j<len(a);j++ {
		key:=a[j]
		i:=j-1
		//此处同上，注意跳出循环的边界值
		for i>=0&&a[i]>key  {
			a[i+1]=a[i]
			i--
		}
		a[i+1]=key
	}
	return a
}

func main() {
	a := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(InsertionSort(a))
}