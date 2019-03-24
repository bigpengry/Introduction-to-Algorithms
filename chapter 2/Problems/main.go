package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(A []int,start,end int){
	//注意课本上的数组下标是从1开始，而实际实现中是从0开始
	for j:=start+1;j<=end;j++ {
		key:=A[j]
		i:=j-1
		//此处同上，注意跳出循环的边界值
		for i>=start&&A[i]>key  {
			A[i+1]=A[i]
			i--
		}
		A[i+1]=key
	}
}

//数组下标从0开始
func merge(A []int,start,mid,end int){
	n1:=mid-start+1
	n2:=end-mid
	l:=make([]int,n1)
	r:=make([]int,n2)
	//这里数组下标从0开始，所以不用减1
	for i:=0;i<n1;i++{
		l[i]=A[start+i]
	}
	for i:=0;i<n2;i++{
		r[i]=A[mid+1+i]
	}
	i,j:=0,0
	for k:=start;k<=end;k++{
		//如果i或者j其中任意一个到达自己的数组的底部，但是k没有到达数组底部，说明l或r其中一个数组仍有剩余元素
		if i==n1{
			for index:=0;j<n2;j++ {
				A[k+index]=r[j]
				index++
				return
			}
		}else if j==n2{
			for index:=0;i<n1;i++ {
				A[k+index]=l[i]
				index++
				return
			}
		}

		if l[i]<r[j] {
			A[k]=l[i]
			i++
		}else {
			A[k]=r[j]
			j++
		}
	}
}

func MergeSort(A []int,start,end int) []int {
	if start<end {
		mid:=(start+end)/2
		MergeSort(A,start,mid)
		MergeSort(A,mid+1,end)
		merge(A,start,mid,end)
	}
	return A
}

func MixSort(A []int,start,end int)[]int{
	if end-start<100{
		InsertionSort(A,start,end)
	}else {
		mid:=(start+end)/2
		MixSort(A,start,mid)
		MixSort(A,mid+1,end)
		merge(A,start,mid,end)
	}
	return A
}

func init(){
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func main() {
	n:=1000000
	A:=make([]int,n)
	B:=make([]int,n)
	for i:=0;i<n ;i++  {
		A[i]=rand.Intn(n)
	}
	B=A

	t1 := time.Now()
	MixSort(A,0,len(A)-1)
	runtimeMixSort := time.Since(t1)
	t2 := time.Now()
	MergeSort(B,0,len(A)-1)
	runtimeMergeSort := time.Since(t2)
	for k,v:=range A{
		if v!=B[k] {
			fmt.Println("A!=B")
			break
		}
	}
	fmt.Println("MixSort cost time is:",runtimeMixSort)
	fmt.Println("MergeSort cost time is:",runtimeMergeSort)
}
