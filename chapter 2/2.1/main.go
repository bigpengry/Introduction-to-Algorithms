package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(A *[]int,start,end int){
	//注意课本上的数组下标是从1开始，而实际实现中是从0开始
	for j:=start+1;j<end;j++ {
		key:=(*A)[j]
		i:=j-1
		//此处同上，注意跳出循环的边界值
		for i>=start&&(*A)[i]>key  {
			(*A)[i+1]=(*A)[i]
			i--
		}
		(*A)[i+1]=key
	}
}

func init(){
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func main() {
	n:=1000
	A:=make([]int,n)
	B:=make([]int,n)
	for i:=0;i<n ;i++  {
		A[i]=rand.Intn(n)
		B[i]=rand.Intn(n)
	}
	m:=make(map[int]int)
	for _,v:=range A {
		m[v]=m[v]+1
	}
	fmt.Println(m)

	fmt.Println(A)
	InsertionSort(&A,0,len(A))
	fmt.Println(A)
}