package main

import "fmt"

//数组下标从0开始
func merge(A *[]int,start,mid,end int){
	n1:=mid-start+1
	n2:=end-mid
	l:=make([]int,n1)
	r:=make([]int,n2)
	//这里数组下标从0开始，所以不用减1
	for i:=0;i<n1;i++{
		l[i]=(*A)[start+i]
	}
	for i:=0;i<n2;i++{
		r[i]=(*A)[mid+1+i]
	}
	i,j:=0,0
	for k:=start;k<=end;k++{
		//如果i或者j其中任意一个到达自己的数组的底部，但是k没有到达数组底部，说明l或r其中一个数组仍有剩余元素
		if i==n1{
			for index:=0;j<n2;j++ {
				(*A)[k+index]=r[j]
				index++
				return
			}
		}else if j==n2{
			for index:=0;i<n1;i++ {
				(*A)[k+index]=l[i]
				index++
				return
			}
		}

		if l[i]<r[j] {
			(*A)[k]=l[i]
			i++
		}else {
			(*A)[k]=r[j]
			j++
		}
	}
}

func MergeSort(A []int,start,end int) []int {
	if start<end {
		mid:=(start+end)/2
		MergeSort(A,start,mid)
		MergeSort(A,mid+1,end)
		merge(&A,start,mid,end)
	}
	return A
}



func main() {
	A:=[]int{3,41,52,26,38,57,9,49}
	fmt.Println(MergeSort(A,0,len(A)-1))
}
